package middleware

import (
	"fmt"
	"livecode-wmb-rest-api/utils"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type AuthTokenMiddleware interface {
	RequireToken() gin.HandlerFunc
}

type authTokenMiddleware struct {
	acctToken utils.Token
}

func (a *authTokenMiddleware) RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":"Unauthorized",
			})
			c.Abort()
		}

		tokenString := strings.Replace(h.AuthorizationHeader,"Bearer ","",-1)
		fmt.Println("tokenstring: ",tokenString)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":"Unauthorized at Empty TokenString",
			})
			c.Abort()
			return
		}

		token, err := a.acctToken.VerifyAccessToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":"Unauthorized at Verify",
			})
			c.Abort()
			return
		}
		userId, err := a.acctToken.FetchAccessToken(token)
		log.Println(token)
		log.Println(userId)
		if err != nil || userId == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":"Unauthorized at Fetch",
			})
			c.Abort()
			return
		}
		//fmt.Println("token: ", token)		
		
		if token != nil {
			c.Set("user-id", userId)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":"Unauthorized at Nil Token",
			})
			c.Abort()
			return
		}
	}
}

func NewTokenValidator(acctToken utils.Token) AuthTokenMiddleware {
	return &authTokenMiddleware{acctToken: acctToken}
}