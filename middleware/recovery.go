package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
)

// Recovery é um middleware que recupera de panics na aplicação
// e garante que a requisição termine graciosamente com um erro 500
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Obtém o timestamp atual
				currentTime := time.Now().Format("2006/01/02 - 15:04:05")

				// Obtém o stack trace
				stack := debug.Stack()

				// Log do erro com informações detalhadas
				errorLog := fmt.Sprintf("[Recovery] %s\nPanic: %v\nURL: %s\nMethod: %s\nStack Trace:\n%s",
					currentTime,
					err,
					c.Request.URL.Path,
					c.Request.Method,
					string(stack),
				)

				// Imprime o log em vermelho para destaque
				fmt.Printf("\033[31m%s\033[0m\n", errorLog)

				// Responde com status 500 e mensagem de erro
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"erro": "Ocorreu um erro interno no servidor",
				})
			}
		}()

		c.Next()
	}
}