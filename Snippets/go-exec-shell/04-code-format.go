package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

//通过管道同步获取日志的函数
func syncLog(logger *log.Logger, reader io.ReadCloser) {
	//因为logger的print方法会自动添加一个换行，所以我们需要一个cache暂存不满一行的log
	cache := ""
	buf := make([]byte, 1024, 1024)
	for {
		strNum, err := reader.Read(buf)
		if strNum > 0 {
			outputByte := buf[:strNum]
			//这里的切分是为了将整行的log提取出来，然后将不满整行和下次一同打印
			outputSlice := strings.Split(string(outputByte), "\n")
			logText := strings.Join(outputSlice[:len(outputSlice)-1], "\n")
			logger.Printf("%s%s", cache, logText)
			cache = outputSlice[len(outputSlice)-1]
		}
		if err != nil {
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
	//打开一个文件，用作log封装输出
	f, _ := os.OpenFile("file.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	//创建封装的log，第三个参数设置log输出的格式
	logger := log.New(f, "", log.LstdFlags)
	logger.Print("start print log:")
	oldFlags := logger.Flags()
	//为了保证shell的输出和标准的log格式不冲突，并且为了整齐，关闭logger自身的格式
	logger.SetFlags(0)
	go syncLog(logger, cmdStdoutPipe)
	go syncLog(logger, cmdStderrPipe)
	err = cmd.Wait()
	//执行完后再打开log输出的格式
	logger.SetFlags(oldFlags)
	logger.Print("log print done")
	if err != nil {
		fmt.Println(err)
	}
}
