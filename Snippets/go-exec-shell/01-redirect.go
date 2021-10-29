package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//定义一个每秒1次输出的shell
	cmdStr := `
#!/bin/bash
for var in {1..10}
do
	sleep 1
     echo "Hello, Welcome ${var} times "
done`
	cmd := exec.Command("bash", "-c",
		cmdStr+" >> file.log") //重定向
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
}

