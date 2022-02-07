package handler

import (
	"GO-GIN_REST_API/entity"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

func CreateToken(email string) (*entity.TokenMetadata, error) {
	tm := &entity.TokenMetadata{}
	tm.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	tm.AccessUuid = uuid.NewV4().String()

	tm.RtExpires = time.Now().Add(time.Hour * 24).Unix()
	tm.RefreshUuid = uuid.NewV4().String()

	var err error

	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = tm.AccessUuid
	atClaims["email"] = email
	atClaims["exp"] = tm.AtExpires

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tm.AccessToken, err = accessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf")
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = tm.RefreshUuid
	rtClaims["email"] = email
	rtClaims["exp"] = tm.RtExpires

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	tm.RefreshToken, err = refreshToken.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}

	return tm, nil
}
