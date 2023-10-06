package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type TODO struct {
	Contnet string `json:"content"`
	Done    bool   `json:"done"`
}

var todos []TODO

func main() {
	r := gin.Default()

	// 添加 TODO
	r.POST("/todo", func(c *gin.Context) {
		var todo TODO
		c.BindJSON(&todo)
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 删除 TODO
	r.DELETE("/todo/:index", func(c *gin.Context) {
		index, _ := strconv.Atoi(c.Param("index"))
		todos = append(todos[:index], todos[index+1:]...)
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 修改 TODO
	r.PUT("/todo/:index", func(c *gin.Context) {
		index, _ := strconv.Atoi(c.Param("index"))
		var todo TODO
		c.BindJSON(&todo)
		todos[index] = todo
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 获取 TODO
	r.GET("/todo", func(c *gin.Context) {
		c.JSON(200, todos)
	})

	// 查询TODO
	r.GET("/todo/:index", func(c *gin.Context) {
		index, _ := strconv.Atoi(c.Param("index"))
		c.JSON(200, todos[index])
	})
}
