package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("ansible-playbook", "-i /root/install/go-ansible/example/golang-lab/inventorydir redis.yaml")
	//cmd  := exec.Command("rm", "/root/install/exec-go/test/ok")

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Println("Return with a nil")
	}


	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	//err = cmd.Wait()

	log.Printf("Command finished with error: %v", err)
}
