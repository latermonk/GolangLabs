package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("ansible-playbook", "-i /root/install/go-ansible/example/golang-lab/inventorydir", "redis.yaml")
	pipe, _ := cmd.StdoutPipe()


	//打印輸出的命令
	fmt.Println(cmd)


	//执行Start函数
	errr := cmd.Start()
	if errr != nil {
		log.Fatal(errr)
	} else {
		fmt.Println("Return with a nil")
	}


	//读取执行结果并打印结果
	reader := bufio.NewReader(pipe)
	line, err2 := reader.ReadString('\n')
	for err2 == nil {
		fmt.Println(line)
		line, err2 = reader.ReadString('\n')
	}

	//执行wait函数
	log.Printf("Waiting for command to finish...")
	err := cmd.Wait()
	log.Printf("Command finished with error: %v", err)


}


