package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {

	config := &rest.Config{
		Host: "http://127.0.0.1:8001",
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(clientSet.ServerVersion())

	list, err := clientSet.CoreV1().Pods("myweb").List(context.Background(), metav1.ListOptions{})
	for _, item := range list.Items {
		fmt.Println(item.Name)
	}
}
