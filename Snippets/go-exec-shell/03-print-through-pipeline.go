/*
exec.Shell不仅提供了Stdout和Stdin这两个变量，还有下面两个方法

func (c *Cmd) StdoutPipe() (io.ReadCloser, error) {
	...
}

func (c *Cmd) StderrPipe() (io.ReadCloser, error) {
	...
}


 */
package main

import (
"fmt"
"io"
"os"
"os/exec"
"strings"
)

//通过管道同步获取日志的函数
func syncLog(reader io.ReadCloser) {
	f, _ := os.OpenFile("file.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	buf := make([]byte, 1024, 1024)
	for {
		strNum, err := reader.Read(buf)
		if strNum > 0 {
			outputByte := buf[:strNum]
			f.WriteString(string(outputByte))
		}
		if err != nil {
			//读到结尾
			if err == io.EOF || strings.Contains(err.Error(), "file already closed") {
				err = nil
			}
		}
	}
}

func main() {
	//定义一个每秒1次输出的shell
	cmdStr := `
#!/bin/bash
for var in {1..10}
do
	sleep 1
     echo "Hello, Welcome ${var} times "
done`
	cmd := exec.Command("bash", "-c", cmdStr)
	//这里得到标准输出和标准错误输出的两个管道，此处获取了错误处理
	cmdStdoutPipe, _ := cmd.StdoutPipe()
	cmdStderrPipe, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	go syncLog(cmdStdoutPipe)
	go syncLog(cmdStderrPipe)
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
}


