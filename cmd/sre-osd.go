package main

import (
	"flag"
	"fmt"
	"github.com/openshiftoperations/sre-osd/pkg/srek8s"
	//"io/ioutil"
	//apiv1 "k8s.io/api/core/v1"
	//rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/client-go/kubernetes"
	//"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
	//"reflect"
)

var clusterFlag = flag.String("cluster", "", "Specify the name of the cluster")
var configFlag = flag.String("fileconfig", "", "Specify the k8s object config file")

func init() {
	flag.StringVar(clusterFlag, "c", "", "Specify the name of the cluster")
	flag.StringVar(configFlag, "f", "", "Specify the k8s object config file")
}

func loadKubeConfig() *kubernetes.Clientset {

	//Build the path to the kubeconfig
	home := homedir.HomeDir()
	var kubeconfig string = filepath.Join(home, ".kube", "config")

	//Use the config to create the client
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return clientset
}

func checkErr(err error) {

	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %s", err))
	}
}

func main() {

	flag.Parse()

	clientset := loadKubeConfig()

	srek8s.Create(clientset, *configFlag)
}
