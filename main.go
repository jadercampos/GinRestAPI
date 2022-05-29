package main

import (
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Jader Campos",
	})
}

func main() {
	r := gin.Default()
	r.GET("/alunos", ExibeTodosAlunos)
	r.Run(":5000")
}
