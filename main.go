package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os/exec"
)

func main() {
	// 创建一个Gin路由器
	r := gin.Default()

	// 定义路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World! ",
		})
	})
	// git command
	r.POST("/cmd", func(c *gin.Context) {
		shell := c.Query("shell")
		shellFile := "cmd/" + shell + ".sh"
		out, err := exec.Command("bash", shellFile).Output()
		if err != nil {
			fmt.Println("执行命令出错：", err)
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"message": out,
			})
		}
	})

	// 启动Gin应用程序
	err := r.Run()
	if err != nil {
		return
	}
}
