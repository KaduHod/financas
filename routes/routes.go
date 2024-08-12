package routes

import (
	"financas/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
func Router (router *gin.Engine) {
	router.Static("/public", "./public")
    router.LoadHTMLGlob("templates/*")
    router.GET("/", func (c *gin.Context) {
        c.JSON(200, gin.H{
            "mensagem": "Olá, algum erro aconteceu para voce estar aqui!",
        })
    })
	router.GET("/financas/grafico", func (c* gin.Context) {
		c.HTML(200 , "grafico.tmpl", gin.H{})
	})
    router.GET("/financas/juroscomposto/simular",func(c *gin.Context) {
		var aplicacao services.AplicacaoFinanceira
		var aporte services.AporteAplicacaoFinanceira
		var erros []string
		quantidadeDeMeses, err := strconv.Atoi(c.Query("quantidadeDeMeses"))
		if err != nil {
			erros = append(erros, "quantidadeDeMeses inválido")
		}

		valorInicial, err := strconv.ParseFloat(c.Query("valorInicial"), 64)
		if err != nil {
			erros = append(erros, "valorInicial inválido")
		}

		taxa, err := strconv.ParseFloat(c.Query("taxa"), 64)
		if err != nil {
			erros = append(erros, "taxa inválida")
		}

		valorAporte, err := strconv.ParseFloat(c.Query("valorAporte"), 64)
		if err != nil {
			erros = append(erros, "valorAporte inválido")
		}

		valorAumentoAporte, err := strconv.ParseFloat(c.Query("valorAumentoAporte"), 64)
		if err != nil {
			erros = append(erros, "valorAumentoAporte inválido")
		}

		if len(erros) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"erros": erros,
				"mensagem": "falha",
			})
			return
		}

		frequenciaAporte := services.PegaTipoFrequencia(c.Query("frequenciaAporte"))
		frequenciaAumentoAporte := services.PegaTipoFrequencia(c.Query("frequenciaAumentoAporte"))
		aplicacao.QuantidadeDeMeses = quantidadeDeMeses
		aplicacao.ValorInicial = valorInicial
		aplicacao.Taxa = taxa
		aporte.ValorAporte = valorAporte
		aporte.ValorAumentoAporte = valorAumentoAporte
		aporte.FrequenciaAporte = frequenciaAporte
		aporte.FrequenciaAumentoAporte = frequenciaAumentoAporte
		aplicacao.Aporte = aporte
		aplicacao.IniciarSimulacao()
		c.JSON(200, gin.H{
			"aplicacao": aplicacao,
			"mensagem": "sucesso",
		})
	})
}
