package common

import (
	"crypto/rsa"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/context"
	"go.uber.org/zap"
)

// AppClaims provides custom claim for JWT
type AppClaims struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

//using asymmetric crypto/RSA keys
const (
	//openssl genrsa -out app.rsa 1024
	privKeyPath = "keys/app.rsa"
	//openssl rsa-in app.rsa -pubout > app.rsa.pub
	pubKeyPath = "keys/app.rsa.pub"
)

//private key for signing and public key for verification
var (

	//verifyKey, signKey []byte
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

//Read the key files before starting http handlers
func initKeys(log *zap.Logger) error {
	var err error

	signBytes, err := ioutil.ReadFile(privKeyPath)

	if err != nil {
		log.Warn("[initKeys]: Error has occurred reading private Key File")
		return err
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Warn("[initKeys]: Error has occurred while parsing private key File")
		return err
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)

	if err != nil {
		log.Warn("[initKeys]: Error has occurred while reading public key File")
		return err
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Warn("[initKeys]: Error has occurred while parsing private key File")
		return err
	}

	return nil

}

//GenerateJWT Tokens
func GenerateJWT(name, role string) (string, error) {
	//create a signer for rsa 256

	// Create the Claims
	claims := AppClaims{
		name,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			Issuer:    "admin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	ss, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

//Authorize Middleware for validating JWT tokens
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Get token from request
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})

	if err != nil {
		switch err.(type) {

		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired: //JWT expired
				DisplayAppError(
					w,
					err,
					"Access Token is expired, get a new Token",
					401,
				)
				return

			default:
				DisplayAppError(w,
					err,
					"Error while parsing the Access Token!",
					500,
				)
				return
			}

		default:
			DisplayAppError(w,
				err,
				"Error while parsing Access Token!",
				500)
			return
		}

	}
	if token.Valid {
		//Set user name to HTTP context
		context.Set(r, "user", token.Claims.(*AppClaims).UserName)
		next(w, r)
	} else {
		DisplayAppError(
			w,
			err,
			"Invalid Access Token",
			401,
		)
	}
}

//JwtAuthorize Middleware for validating JWT tokens
func JwtAuthorize(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get token from request
		token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			return verifyKey, nil
		})

		if err != nil {
			switch err.(type) {

			case *jwt.ValidationError: // JWT validation error
				vErr := err.(*jwt.ValidationError)

				switch vErr.Errors {
				case jwt.ValidationErrorExpired: //JWT expired
					DisplayAppError(
						w,
						err,
						"Access Token is expired, get a new Token",
						401,
					)
					return

				default:
					DisplayAppError(w,
						err,
						"Error while parsing the Access Token!",
						500,
					)
					return
				}

			default:
				DisplayAppError(w,
					err,
					"Error while parsing Access Token!",
					500)
				return
			}

		}
		if token.Valid {
			//Set user name to HTTP context
			context.Set(r, "user", token.Claims.(*AppClaims).UserName)
			next.ServeHTTP(w, r)
		} else {
			DisplayAppError(
				w,
				err,
				"Invalid Access Token",
				401,
			)
		}

	})
}
