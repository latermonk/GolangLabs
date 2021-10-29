package main

import (
	"fmt"
	"log"

	"github.com/mattn/go-pipeline"
)


func main() {
	out, err := pipeline.Output(
		[]string{"ansible-playbook", "-i", "/root/install/go-ansible/example/golang-lab/inventorydir", "redis.yaml"},
	)
	if err != nil {
		log.Fatal(err)
	}


	// Print the result
	fmt.Println(string(out))
}