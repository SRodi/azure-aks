package test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// fetch the sensitive output from OpenTofu
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
		return nil, fmt.Errorf("failed to create Kubernetes client: %v", err)
	}
	return clientset, nil
}

// creates a new REST config using the provided options
func newRESTConfig(caCert, clientKey, clientCert, host string) *rest.Config {
	return &rest.Config{
		Host: host,
		TLSClientConfig: rest.TLSClientConfig{
			CAData:   []byte(caCert),
			CertData: []byte(clientCert),
			KeyData:  []byte(clientKey),
		},
	}
}
