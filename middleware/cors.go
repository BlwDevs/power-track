package middleware

import (
	"github.com/gin-gonic/gin"
)

// Cors retorna um middleware Gin que habilita o Cross-Origin Resource Sharing (CORS)
// para permitir requisições de diferentes origens. Este middleware é essencial para
// APIs que precisam ser acessadas por aplicações frontend em diferentes domínios.
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Permite requisições de qualquer origem (*).
		// Em produção, considere restringir para origens específicas.
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// Permite o envio de credenciais (cookies, headers de autenticação)
		// nas requisições cross-origin
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Define os headers HTTP que podem ser utilizados pela aplicação cliente
		// Inclui headers comuns como Content-Type, Authorization, etc.
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		// Define os métodos HTTP permitidos nas requisições cross-origin
		// POST: Criar recursos
		// OPTIONS: Pré-flight requests
		// GET: Buscar recursos
		// PUT: Atualizar recursos
		// DELETE: Remover recursos
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		// Trata requisições OPTIONS (pre-flight)
		// Retorna 204 (No Content) para confirmar que a requisição real pode ser feita
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// Continua para o próximo middleware ou handler na cadeia
		c.Next()
	}
}