package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger é um middleware que registra informações sobre as requisições HTTP
// incluindo método, path, status code, latência e outros detalhes relevantes
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Marca o tempo de início da requisição
		startTime := time.Now()

		// Processa a requisição
		c.Next()

		// Coleta informações após o processamento
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// Monta o log com as informações da requisição
		log := fmt.Sprintf("[%s] | %s | %s | %d | %v | %s",
			endTime.Format("2006-01-02 15:04:05"),  // Timestamp
			c.Request.Method,                        // Método HTTP
			c.Request.URL.Path,                      // Caminho da requisição
			c.Writer.Status(),                       // Status HTTP
			latency,                                 // Tempo de processamento
			c.ClientIP(),                            // IP do cliente
		)

		// Define a cor do log baseado no status HTTP
		switch {
		case c.Writer.Status() >= 500:
			// Vermelho para erros do servidor
			fmt.Printf("\033[31m%s\033[0m\n", log)
		case c.Writer.Status() >= 400:
			// Amarelo para erros do cliente
			fmt.Printf("\033[33m%s\033[0m\n", log)
		case c.Writer.Status() >= 300:
			// Ciano para redirecionamentos
			fmt.Printf("\033[36m%s\033[0m\n", log)
		default:
			// Verde para sucessos
			fmt.Printf("\033[32m%s\033[0m\n", log)
		}
	}
}