package main

import (
	"fmt"
	"log"
	"os"

	"k8s.io/apimachinery/pkg/api/resource"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] != "convert" {
		fmt.Println("Usage: convert <resource>")
		os.Exit(1)
	}
	q, err := resource.ParseQuantity(os.Args[2])

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(q.AsDec())
}
