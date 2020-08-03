package main

import (
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	kubeconfig := flag.String("kubeconfig","~/.kube/config","kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("",*kubeconfig)

	//config,err := clientcmd.DefaultClientConfig.ClientConfig()
	if err!=nil {
		panic(err)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err!=nil {
		panic(err)
	}

	pod,_ :=  clientSet.CoreV1().Pods("default").Get("details-v1-6c9f8bcbcb-5znvg",metav1.GetOptions{})
	fmt.Println(pod)
}
