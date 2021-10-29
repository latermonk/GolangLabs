package main

import (
	"fmt"
	"os"
	"os/exec"
)

/*
使用exec.Command创建一个Shell后，就具有了两个变量：

Stdout io.Writer
Stderr io.Writer

这两个变量是用来指定程序的标准输出和标准错误输出的位置，
所以我们就利用这两个变量，直接打开文件，
然后将打开的文件指针赋值给这两个变量即可将程序的输出直接输出到文件中，
也能达到相同的效果，
参考程序如下：

 */

func main() {
	fmt.Println("=============")
	//定义一个每秒1次输出的shell
	cmdStr := `
#!/bin/bash
for var in {1..10}
do
	sleep 1
     echo "Hello, Welcome ${var} times "
done`
	cmd := exec.Command("bash", "-c", cmdStr)
	//打开一个文件
	f, _ := os.OpenFile("file.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	//指定输出位置
	cmd.Stderr = f
	cmd.Stdout = f
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
}
