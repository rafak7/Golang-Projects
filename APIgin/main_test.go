package main

import (
	"APIgin/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := setupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req := httptest.NewRequest("GET", "/rafa", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	if resposta.Code != http.StatusOK {
		t.Fatalf("Esperava status 200, mas obteve %d", resposta.Code)
	}
}
