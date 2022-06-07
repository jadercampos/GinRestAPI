package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jadercampos/GinRestAPI/controllers"
	docs "github.com/jadercampos/GinRestAPI/docs"
	websocket "github.com/jadercampos/GinRestAPI/web-socket"
	cors "github.com/rs/cors/wrapper/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	docs.SwaggerInfo.BasePath = "/"
	rWebAPI := gin.Default()
	rWebAPI.Use(cors.Default())
	rWebAPI.LoadHTMLGlob("templates/*")
	rWebAPI.Static("/assets", "./assets")
	rWebAPI.GET("/:nome", controllers.Saudacao)
	rWebAPI.GET("/alunos", controllers.ExibeTodosAlunos)
	rWebAPI.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	rWebAPI.POST("/alunos", controllers.CriaNovoAluno)
	rWebAPI.DELETE("/alunos/:id", controllers.DeletaAlunos)
	rWebAPI.PATCH("/alunos/:id", controllers.EditaAlunos)
	rWebAPI.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	rWebAPI.GET("/index", controllers.ExibePaginaIndex)
	rWebAPI.NoRoute(controllers.RotaNaoEncontrada)
	rWebAPI.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	go rWebAPI.Run(":5000")

	rWebSocket := gin.Default()
	rWebSocket.LoadHTMLGlob("templates/*")
	rWebSocket.GET("/wspage", func(c *gin.Context) { c.HTML(200, "ws.html", nil) })
	rWebSocket.GET("/ws", func(c *gin.Context) { websocket.WsHandler(c.Writer, c.Request) })
	rWebSocket.Run(":5001")
}
