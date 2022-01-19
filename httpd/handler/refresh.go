package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		mapToken := map[string]string{}
		if err := c.ShouldBindJSON(&mapToken); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}

		refreshToken := mapToken["refresh_token"]

		//verify the token

		//os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf")
		token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("REFRESH_SECRET")), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Refresh token expired")
			return
		}

		//is token valid?
		if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		//Since token is valid, get the uuid:
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
			if !ok {
				c.JSON(http.StatusUnprocessableEntity, err)
				return
			}
			email := claims["email"].(string)
			log.Println("********  refreshUuid  *********")
			log.Println(refreshUuid)
			log.Println("********  email  *********")
			log.Println(email)
			//Delete the previous Refresh Token
			deleted, delErr := h.backend.DeleteAuthor(`jwtMetadata:RefreshUuid:` + refreshUuid)
			if delErr != nil || deleted == 0 {
				c.JSON(http.StatusUnauthorized, "unauthorized")
				return
			}
			//Create new pairs of refresh and access tokens

			tokens, createErr := CreateToken(email)
			if createErr != nil {
				c.JSON(http.StatusForbidden, createErr.Error())
				return
			}
			//save the tokens metadata to redis
			saveErr := h.backend.CreateAuthor(email, tokens)
			if saveErr != nil {
				c.JSON(http.StatusForbidden, saveErr.Error())
				return
			}

			c.JSON(http.StatusCreated, gin.H{
				"access_token":  tokens.AccessToken,
				"refresh_token": tokens.RefreshToken,
			})
		} else {
			c.JSON(http.StatusUnauthorized, "refresh expired")
		}
	}
}
