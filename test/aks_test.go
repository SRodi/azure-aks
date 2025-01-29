package test

import (
	"context"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func TestAKSCluster(t *testing.T) {
	t.Parallel()

	// Define the OpenTofu options
	terraformOptions := &terraform.Options{
		// The path to where your OpenTofu code is located
		TerraformDir: "../examples/aks",

		// Variables to pass to our OpenTofu code using -var options
		Vars: map[string]interface{}{
			"resource_group_name": "test-rg",
			"location":            "uksouth",
			"prefix":              "test",
			"labels": map[string]string{
				"env": "test",
			},
		},
	}

	// Clean up resources with "tofu destroy" at the end of the test
	defer terraform.Destroy(t, terraformOptions)
	// Run "tofu init" and "tofu apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Fetch the outputs
	caCert := fetchSensitiveOutput(t, terraformOptions, "ca_certificate")
	clientKey := fetchSensitiveOutput(t, terraformOptions, "client_key")
	clientCert := fetchSensitiveOutput(t, terraformOptions, "client_certificate")
	host := fetchSensitiveOutput(t, terraformOptions, "host")

	// Decode the base64 encoded strings
	caCertDecoded := decodeBase64(t, caCert)
	clientKeyDecoded := decodeBase64(t, clientKey)
	clientCertDecoded := decodeBase64(t, clientCert)

	// Create a new REST config using the outputs
	restConfig := newRESTConfig(caCertDecoded, clientKeyDecoded, clientCertDecoded, host)

	// Create a new Kubernetes client using the REST config
	k8sClient, err := newK8sClient(restConfig)
	if err != nil {
		t.Fatalf("Failed to create Kubernetes client: %v", err)
	}

	// Use the Kubernetes client to interact with the cluster
	_, err = k8sClient.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		t.Fatalf("Failed to list namespaces: %v", err)
	}
}

// fetche the sensitive output from OpenTofu
func fetchSensitiveOutput(t *testing.T, options *terraform.Options, name string) string {
	defer func() {
		options.Logger = nil
	}()
	options.Logger = logger.Discard
	return terraform.Output(t, options, name)
}

// decode the base64 encoded string
func decodeBase64(t *testing.T, encoded string) string {
	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		t.Fatalf("Failed to decode base64 string: %v", err)
	}
	return string(decodedBytes)
}

// newK8sClient creates a new Kubernetes client using REST config
func newK8sClient(restConfig *rest.Config) (*kubernetes.Clientset, error) {
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Kubernetes client: %v", err)
	}
	return clientset, nil
}

// creates a new REST config using the provided options
func newRESTConfig(caCert, clientKey, clientCert, host string) *rest.Config {
	// Implement the logic to create a Kubernetes client using the provided options
	// This is a placeholder implementation
	return &rest.Config{
		Host: host,
		TLSClientConfig: rest.TLSClientConfig{
			CAData:   []byte(caCert),
			CertData: []byte(clientCert),
			KeyData:  []byte(clientKey),
		},
	}
}
