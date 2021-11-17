#   MiddleWare 


#  基本认证的中间件

```
	r := gin.Default()

	r.Use(gin.BasicAuth(gin.Accounts{
		"admin": "123456",
	}))
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "首页")
	})
	
	r.Run(":8080")
  
 ```
 
 #  针对特定URL的Basic Authorization
 
 ```
 func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "首页")
	})

	adminGroup := r.Group("/admin")
	adminGroup.Use(gin.BasicAuth(gin.Accounts{
		"admin": "123456",
	}))

	adminGroup.GET("/index", func(c *gin.Context) {
		c.JSON(200, "后台首页")
	})

	r.Run(":8080")
}
 
 ```
 
 
 #  自定义中间件
 
 func costTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求前获取当前时间
		nowTime := time.Now()

		//请求处理
		c.Next()

		//处理后获取消耗时间
		costTime := time.Since(nowTime)
		url := c.Request.URL.String()
		fmt.Printf("the request URL %s cost %v\n", url, costTime)
	}
}


func main() {
	r := gin.New()

	r.Use(costTime())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "首页")
	})

	r.Run(":8080")
}


