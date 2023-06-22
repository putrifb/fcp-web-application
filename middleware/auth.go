package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		//Mengambil cookie dengan nama session_token dari request dengan key session_token. Cookie ini berisi JWT token yang digunakan untuk autentikasi.
		res, err := ctx.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				if ctx.GetHeader("Content-Type") == "application/json" {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				} else {
					ctx.Redirect(http.StatusSeeOther, "/user/login")
					ctx.Abort()
				}
				return
			}
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		//Parsing JWT token pada cookie tersebut untuk mendapatkan claims yang berisi informasi mengenai email
		c := &model.Claims{}

		token, err := jwt.ParseWithClaims(res, c, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid"})
			return
		}
		ctx.Set("email", c.Email)
		ctx.Next() // TODO: answer here
	})
}
