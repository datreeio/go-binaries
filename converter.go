package main

import (
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/api/resource"
)

func main() {
	q, err := resource.ParseQuantity(os.Args[1])

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(q.AsDec())
}
