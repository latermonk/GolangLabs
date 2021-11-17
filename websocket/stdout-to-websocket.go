package main

import (

	"github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
)

var idb *gorm.DB


var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


func wsplay(c *gin.Context){

	//wshandler(c.Writer, c.Request)

	// do install
	fmt.Println("=====wsplay=====")

	//Beginning
	args := []string{"-i", "/inventorydir", "redis.yaml"}
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

	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	//defer conn.Close()

	//echo 返回收到的信息
	go func() {
		for {
			t, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}

			conn.WriteMessage(t, msg)
			//conn.WriteMessage(t, pipe)

		}
		//conn.WriteMessage()


	}()

	s := bufio.NewScanner(io.MultiReader(pipe))

	for s.Scan() {

		conn.WriteMessage(1, s.Bytes())

	}


	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)

	fmt.Println("End of Playme")

}


/*  主函数 */

func main() {
	//gin server初始化
	r := gin.Default()

	//gin响应服务
	r.GET("/ws", wsplay) 

	//gin监听端口
	r.Run(":9999")

}

