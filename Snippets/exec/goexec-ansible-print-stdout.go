package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func main() {

	//Beginning
	args := []string{"-i", "/root/install/go-ansible/example/golang-lab/inventorydir", "redis.yaml"}
	cmd := exec.Command("ansible-playbook", args...)

	//构建Pipe管道
	pipe, _ := cmd.StdoutPipe()

	//打印輸出的命令
	fmt.Println(cmd)

	//exec Start func
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Return with a nil")
	}

	//READ AND PRINT THE RESULT TOSCREEN
	reader := bufio.NewReader(pipe)
	line, err2 := reader.ReadString('\n')
	for err2 == nil {
		fmt.Println(line)
		line, err2 = reader.ReadString('\n')
	}

	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}
