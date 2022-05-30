package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jadercampos/GinRestAPI/database"
	"github.com/jadercampos/GinRestAPI/models"
	_ "github.com/swaggo/swag/example/celler/httputil"
	_ "github.com/swaggo/swag/example/celler/model"
)

// Saudacao godoc
// @Summary      Exibe uma saudção
// @Description  Só está te cumprimentando UAI!
// @Tags         oi_sumido
// @Accept       json
// @Produce      json
// @Param        nome   path      string  true  "Nome do caboclo"
// @Success      200  {string}   models.Aluno
// @Failure      404  {object}  httputil.HTTPError
// @Router       /{nome} [get]
func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}

// ExibeTodosAlunos godoc
// @Summary      Exibe todos os alunos
// @Description  Exibe todos os alunos cadastrados
// @Tags         alunos
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Aluno
// @Failure      404  {object}  httputil.HTTPError
// @Router       /alunos [get]
func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

// BuscaAlunoPorID godoc
// @Summary      Exibe um aluno
// @Description  Busca o aluno pelo ID
// @Tags         aluno
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID do Aluno"
// @Success      200  {object}   models.Aluno
// @Failure      404  {object}  httputil.HTTPError
// @Router       /alunos/{id}  [get]
func BuscaAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

// CriaNovoAluno godoc
// @Summary      Adiciona aluno
// @Description  Adiciona um novo aluno
// @Tags         aluno
// @Accept       json
// @Produce      json
// @Param        account  body      models.Aluno  true  "Novo Aluno"
// @Success      200      {object}  models.Aluno
// @Failure      404      {object}  httputil.HTTPError
// @Router       /alunos [put]
func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// DeletaAlunos godoc
// @Summary      Deleta aluno
// @Description  Deleta aluno pelo ID
// @Tags         aluno
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID do Aluno"
// @Success      200  {object}   models.Aluno
// @Failure      404  {object}  httputil.HTTPError
// @Router       /alunos/{id}  [delete]
func DeletaAlunos(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

// EditaAlunos godoc
// @Summary      Atualiza aluno
// @Description  Atualiza os dados de um aluno existente
// @Tags         aluno
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID do Aluno"
// @Param        account  body      models.Aluno  true  "Novo Aluno"
// @Success      200  {object}   models.Aluno
// @Failure      404  {object}  httputil.HTTPError
// @Router       /alunos/{id}  [patch]
func EditaAlunos(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorCPF godoc
// @Summary      Exibe um aluno
// @Description  Busca o aluno pelo CPF
// @Tags         aluno
// @Accept       json
// @Produce      json
// @Param        cpf   path      string  true  "CPF do Aluno"
// @Success      200  {object}   models.Aluno
// @Failure      404  {object}  httputil.HTTPError
// @Router       /alunos/cpf/{cpf}  [get]
func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Params.ByName("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
