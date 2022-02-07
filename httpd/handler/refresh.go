package handler

import (
	"fmt"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Refresh(c *gin.Context) {
	tokenString, err := c.Cookie("refresh")
	log.Println("tokenString: ", tokenString)
	if err == nil && tokenString != "" {
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("REFRESH_SECRET")), nil
		})
		if err != nil {
			ErrorHtml(c, "login.html", "Login", "Refresh failed 1 ", err.Error())
			log.Println("refresh token: ", token)
			return
		}
		if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
			ErrorHtml(c, "login.html", "Login", "Refresh failed 2 ", err.Error())
			return
		}

		//Since token is valid, get the uuid:
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			log.Println("claims: ", claims)

			refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
			if !ok {
				ErrorHtml(c, "login.html", "Login", "Refresh failed 3 ", err.Error())
				return
			}
			if deleted, delErr := h.backend.DeleteAuthor(`jwtMetadata:RefreshUuid:` + refreshUuid); delErr != nil || deleted == 0 {
				ErrorHtml(c, "login.html", "Login", "Refresh failed 4 ", "Unauthorized")
				return
			}
			//Create new pairs of refresh and access tokens
			c.SetCookie("token", "", -1, "", "", false, true)
			c.SetCookie("refresh", "", -1, "", "", false, true)

			email := claims["email"].(string)
			tokens, createErr := CreateToken(email)
			if createErr != nil {
				ErrorHtml(c, "login.html", "Login", "Refresh failed", createErr.Error())
				return
			}

			if saveErr := h.backend.CreateAuthor(email, tokens); saveErr != nil {
				ErrorHtml(c, "login.html", "Login", "Refresh failed 5 ", saveErr.Error())
				return
			}
			log.Print("access_token: ", tokens.AccessToken)
			log.Print("refresh_token: ", tokens.RefreshToken)
			c.SetCookie("token", tokens.AccessToken, 900, "", "", false, true)
			c.SetCookie("refresh", tokens.RefreshToken, 86400, "", "", false, true)
			log.Println("Successfully refreshed again!!")
		} else {
			ErrorHtml(c, "login.html", "Login", "Refresh failed 6 ", "refresh has expired. Please log in.")
		}
		return
	} else {
		ErrorHtml(c, "login.html", "Login", "Refresh failed 7 ", err.Error())
	}

}
