package midllewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeRole(secretKey string, allowedRoles ...string) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// 1️⃣ Récupérer Authorization: Bearer xxx
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 2️⃣ Décoder le token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		// 3️⃣ Lire les claims (email + role)
		claims := token.Claims.(jwt.MapClaims)

		// ⚠ ATTENTION: ils doivent exister
		role, ok := claims["role"].(string)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token payload"})
			ctx.Abort()
			return
		}

		// 4️⃣ Vérifier si le rôle est autorisé
		allowed := false
		for _, r := range allowedRoles {
			if r == role {
				allowed = true
				break
			}
		}

		if !allowed {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			ctx.Abort()
			return
		}

		// 5️⃣ Stocker info utilisateur dans le contexte (optional)
		ctx.Set("email", claims["email"])
		ctx.Set("role", role)

		ctx.Next()
	}
}
