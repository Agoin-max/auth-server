package middleware

import (
	"auth-server/conf"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": "Authorization Failed!", "data": nil})
			ctx.Abort()
		}

		token, err := jwt.Parse(auth, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			// secret key
			secret := conf.AppSetting.JwtSecret
			return []byte(secret), nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx.Set("user", claims)
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			ctx.Abort()
		}
	}
}
