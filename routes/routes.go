package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jadercampos/GinRestAPI/controllers"
	docs "github.com/jadercampos/GinRestAPI/docs"
	cors "github.com/rs/cors/wrapper/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	docs.SwaggerInfo.BasePath = "/"
	r := gin.Default()
	r.Use(cors.Default())
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAlunos)
	r.PATCH("/alunos/:id", controllers.EditaAlunos)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.RotaNaoEncontrada)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":5000")

	//r.LoadHTMLFiles("ws.html")
	// r.GET("/wspage", func(c *gin.Context) { c.HTML(200, "ws.html", nil) })
	// r.GET("/ws", func(c *gin.Context) { websocket.WsHandler(c.Writer, c.Request) })

	// r.Run(":5001")
}
