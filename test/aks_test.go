package test

import (
	"context"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
