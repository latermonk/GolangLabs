package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	fmt.Println("=================================")
	fmt.Println("=============111111111111111111111111====================")

	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	fmt.Println("=============2222222222222222222222222222====================")

	r.GET("/incr", func(c *gin.Context) {

		fmt.Println("================333333333333333333333333333=================")
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})

		fmt.Println("===============444444444444444444444444444==================")
	})
	r.Run(":8000")
}
