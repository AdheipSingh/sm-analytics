package utils

import (
	"log"
	"os"

	versionedclient "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/client-go/tools/clientcmd"
)

// GetIstioClientset shall return the istio client
func GetIstioClientset() *versionedclient.Clientset {
	conf, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		log.Printf("error in getting Kubeconfig: %v", err)
	}

	ic, err := versionedclient.NewForConfig(conf)
	if err != nil {
		log.Fatalf("Failed to create istio client: %s", err)
	}
	return ic
}
