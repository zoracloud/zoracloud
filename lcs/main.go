package main

import (
	"fmt"

	"github.com/zoracloud/zoracloud/lcs/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	fmt.Println(kubernetes.KubeConfig)
	
	
}
