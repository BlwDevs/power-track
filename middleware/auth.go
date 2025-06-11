package middleware

import (
	"net/http"
	"os"
	"power-track/service"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		//tokenStr := token[len("Bearer "):]

		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !parsedToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expirado"})
				c.Abort()
				return
			}
		}

		// Verificar no banco se o usuário esta autenticado
		userID, ok := claims["id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}
		// Checar token no banco de dados

		user, err := userService.GetUserByID(uint(userID))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não encontrado"})
			c.Abort()
			return
		}
		if user.Token != token {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não corresponde ao usuário"})
			c.Abort()
			return
		}

		c.Set("userID", user.ID)

		c.Next()
	}
}
