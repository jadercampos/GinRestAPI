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
	r := gin.Default()

	// Enable CORS
	r.Use(cors.Default())

	// Static Content
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.LoadHTMLFiles("ws.html")

	// Rest API
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAlunos)
	r.PATCH("/alunos/:id", controllers.EditaAlunos)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.RotaNaoEncontrada)

	// Swagger documentation
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Web Socket
	r.GET("/wspage", func(c *gin.Context) { c.HTML(200, "ws.html", nil) })
	r.GET("/ws", func(c *gin.Context) { websocket.WsHandler(c.Writer, c.Request) })

	r.Run(":5000")
}
