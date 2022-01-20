package handler

import (
	"GO-GIN_REST_API/entity"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

/*
func CreateToken(userid uint64) (*TokenMetadata, error)
func ExtractToken(r *http.Request) string
func VerifyToken(r *http.Request) (*jwt.Token, error)
func TokenValid(r *http.Request) error
func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error)
*/

func CreateToken(email string) (*entity.TokenMetadata, error) {
	tm := &entity.TokenMetadata{}
	tm.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	tm.AccessUuid = uuid.NewV4().String()
	/*
		Since the uuid is unique each time it is created, a user can create more than one token. This happens when a user is logged in on different devices. The user can also logout from any of the devices without them being logged out from all devices.
	*/
	tm.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
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

// //extract the token from the request header
// //exsist?
// func ExtractToken(r *http.Request) string {
// 	bearToken := r.Header.Get("Authorization")
// 	//normally Authorization the_token_xxx
// 	strArr := strings.Split(bearToken, " ")
// 	log.Println("ExtractToken.strArr: ", strArr)
// 	if len(strArr) == 2 {
// 		return strArr[1]
// 	}
// 	return ""
// }

// //verity?
// func VerifyToken(r *http.Request) (*jwt.Token, error) {
// 	tokenString := ExtractToken(r)
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		//Make sure that the token method conform to "SigningMethodHMAC"
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("ACCESS_SECRET")), nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	log.Println("VerifyToken.token: ", token)
// 	return token, nil
// }

//expired?valid?
// func TokenValid(r *http.Request) error {
// 	token, err := VerifyToken(r)
// 	if err != nil {
// 		return err
// 	}
// 	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
// 		return err
// 	}
// 	return nil
// }

//extract the token metadata that will lookup in Redis
// func ExtractTokenMetadata(r *http.Request) (*entity.AccessDetails, error) {
// 	token, err := VerifyToken(r)
// 	if err != nil {
// 		return nil, err
// 	}
// 	claims, ok := token.Claims.(jwt.MapClaims)

// 	if ok && token.Valid {
// 		accessUuid, ok := claims["access_uuid"].(string)
// 		if !ok {
// 			return nil, err
// 		}
// 		email := claims["email"].(string)

// 		return &entity.AccessDetails{
// 			AccessUuid: accessUuid,
// 			Email:      email,
// 		}, nil
// 	}
// 	return nil, err
// }
