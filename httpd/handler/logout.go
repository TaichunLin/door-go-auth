package handler

import (
	"fmt"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Logout(c *gin.Context) {
	if h.LogoutDel(c, "token", "ACCESS_SECRET", "access_uuid", "AccessUuid") && h.LogoutDel(c, "refresh", "REFRESH_SECRET", "refresh_uuid", "RefreshUuid") {
		c.SetCookie("token", "", -1, "", "", false, true)
		c.SetCookie("refresh", "", -1, "", "", false, true)
		c.Redirect(302, "/")
	}

}

func (h *Handler) LogoutDel(c *gin.Context, cookie string, secret string, uuid string, pre string) bool {
	tokenString, err := c.Cookie(cookie)
	log.Println("tokenString: ", tokenString)
	if err == nil && tokenString != "" {
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv(secret)), nil
		})
		if err != nil {
			ErrorHtml(c, "text.html", "Logout", "Logout failed", cookie+" has expired")
			return false
		}
		log.Println(cookie, token)
		if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
			ErrorHtml(c, "text.html", "Logout", "Logout failed", err.Error())
			return false
		}

		//Since token is valid, get the uuid:
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			log.Println("claims: ", claims)

			key, ok := claims[uuid].(string) //convert the interface to string
			if !ok {
				ErrorHtml(c, "text.html", "Logout", "Logout failed", err.Error())
				return false
			}

			//Delete the previous Refresh Token
			deleted, delErr := h.backend.DeleteAuthor("jwtMetadata:" + pre + ":" + key)
			if delErr != nil || deleted == 0 {
				ErrorHtml(c, "text.html", "Logout", "Logout failed", "Unauthorized")
				return false
			}
			return true
		} else {
			ErrorHtml(c, "text.html", "Logout", "Logout failed", cookie+" has expired")
			return false
		}

	} else {
		ErrorHtml(c, "text.html", "Logout", "Logout failed", err.Error())
		return false
	}
}
