package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	evenv1 "k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/clientcmd"

	v1 "github.com/openshift/api/image/v1"
	imagev1 "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
)

func main() {
	err := start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
}

func start() error {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return err
	}

	imageV1Client, err := imagev1.NewForConfig(config)
	if err != nil {
		return err
	}

	namespace := "myproject"
	// get all builds
	watches, err := imageV1Client.ImageStreams(namespace).Watch(metav1.ListOptions{})
	if err != nil {
		return err
	}
	Reconcile(<-watches.ResultChan())

	return nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func Reconcile(event evenv1.Event) error {
	switch o := event.Object.(type) {
	case *v1.ImageStream:

		//TODO:something


	}
	return nil
}
