package srek8s

import (
	"fmt"
	"io/ioutil"
	apiv1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"log"
	"reflect"
)

func checkErr(err error) {

	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %s", err))
	}
}

func Create(clientset *kubernetes.Clientset, filename string) {

	//Decode yaml file
	yamlFile, _ := ioutil.ReadFile(filename)
	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, err := decode([]byte(yamlFile), nil, nil)
	checkErr(err)

	switch o := obj.(type) {
	case *apiv1.Namespace:
		fmt.Println("Creating namespace...")
		_, err = clientset.CoreV1().Namespaces().Create(o)
		checkErr(err)
	case *apiv1.ServiceAccount:
		fmt.Println("Creating serviceaccount...")
		fmt.Println(o.Namespace)
		_, err = clientset.CoreV1().ServiceAccounts(o.Namespace).Create(o)
		checkErr(err)
	case *rbacv1.ClusterRole:
		fmt.Println("Creating ClusterRole...")
		_, err = clientset.RbacV1().ClusterRoles().Create(o)
		checkErr(err)
	case *rbacv1.ClusterRoleBinding:
		fmt.Println("Creating ClusterRoleBinding...")
		_, err = clientset.RbacV1().ClusterRoleBindings().Create(o)
		checkErr(err)
	default:
		fmt.Println("K8s object not currently handled")
		fmt.Println(reflect.TypeOf(o))
		fmt.Println(o)
	}
}

func Update(clientset *kubernetes.Clientset, filename string) {

	//Decode yaml file
	yamlFile, _ := ioutil.ReadFile(filename)
	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, err := decode([]byte(yamlFile), nil, nil)
	checkErr(err)

	switch o := obj.(type) {
	case *apiv1.Namespace:
		fmt.Println("Updating namespace...")
		_, err = clientset.CoreV1().Namespaces().Update(o)
		checkErr(err)
	case *apiv1.ServiceAccount:
		fmt.Println("Updating serviceaccount...")
		fmt.Println(o.Namespace)
		_, err = clientset.CoreV1().ServiceAccounts(o.Namespace).Update(o)
		checkErr(err)
	case *rbacv1.ClusterRole:
		fmt.Println("Updating ClusterRole...")
		_, err = clientset.RbacV1().ClusterRoles().Update(o)
		checkErr(err)
	case *rbacv1.ClusterRoleBinding:
		fmt.Println("Updating ClusterRoleBinding...")
		_, err = clientset.RbacV1().ClusterRoleBindings().Update(o)
		checkErr(err)
	default:
		fmt.Println("K8s object not currently handled")
		fmt.Println(reflect.TypeOf(o))
		fmt.Println(o)
	}

}
