package handler

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tingpo/pupgobackend/internal/gorm"
	"google.golang.org/api/oauth2/v2"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type SinginPayload struct {
	Token   string `json:"token"`
	Email   string `json:"email"`
	Account string `json:"account"`
	Type    string `json:"type"`
}

var httpClient = &http.Client{}

var sqlCnter *gorm.SQLCnter
var payloadCreator *gorm.PayloadCreator

func init() {
	mysqlConnector, err := gorm.GetConnectorFactory("mySQL")
	if err != nil {
		panic(fmt.Errorf("Connect to DB failed: %w \n", err))
	}

	db := mysqlConnector.NewDBConnection()
	sqlCnter = gorm.NewSQLCnter(db)
	payloadCreator = gorm.NewPayloadCreator(sqlCnter)
}

func verifyGoogleToken(idToken string) (*oauth2.Tokeninfo, error) {
	oauth2Service, err := oauth2.New(httpClient)
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

func verifyFBToken(idToken string) error {
	query := fmt.Sprintf("https://graph.facebook.com/me?access_token=%s", idToken)
	resp, err := httpClient.Get(query)

	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)

	sb := string(body)
	if resp.StatusCode != http.StatusOK {
		err = errors.New(sb)
	}

	return err
}

func CreateJWT(account string, email string) (token string, err error) {
	now := time.Now()
	jwtId := account + strconv.FormatInt(now.Unix(), 10)
	// set claims and sign
	claims := Claims{
		Account: account,
		Email:   email,
		StandardClaims: jwt.StandardClaims{
			Audience:  account,
			ExpiresAt: now.Add(24 * time.Hour).Unix(),
			Id:        jwtId,
			IssuedAt:  now.Unix(),
			Issuer:    "pupgo",
			NotBefore: now.Add(10 * time.Second).Unix(),
			Subject:   account,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString(jwtSecret)

	return token, err
}

// custom claims
type Claims struct {
	Account string `json:"account"`
	Email   string `json:"email"`
	jwt.StandardClaims
}

// jwt secret key
var jwtSecret = []byte("secret")

// validate JWT
func AuthRequired(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Header with no accesstoken",
		})
		c.Abort()
		return
	}
	token := strings.Split(auth, "Bearer ")[1]

	// parse and validate token for six things:
	// validationErrorMalformed => token is malformed
	// validationErrorUnverifiable => token could not be verified because of signing problems
	// validationErrorSignatureInvalid => signature validation failed
	// validationErrorExpired => exp validation failed
	// validationErrorNotValidYet => nbf validation failed
	// validationErrorIssuedAt => iat validation failed
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})

	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "token could not be verified because of signing problems"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "token is not yet valid before sometime"
			} else {
				message = "can not handle this token"
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": message,
		})
		c.Abort()
		return
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		c.Set("account", claims.Account)
		c.Set("email", claims.Email)
		c.Next()
	} else {
		c.Abort()
		return
	}
}
