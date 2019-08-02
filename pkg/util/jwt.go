package util

import (
	"encoding/json"
	"gin-blog/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	UserId int64
	jwt.StandardClaims
}

func (claims *Claims) SetUser(c *gin.Context) (err error) {
	claimsData, err := json.Marshal(claims)
	c.Set("user", string(claimsData))
	return
}

func GetUser(c *gin.Context) *Claims {
	userData, _ := c.Get("user")
	user :=  &Claims{}

	if userData == nil {
		return user
	}
	err := json.Unmarshal([]byte(userData.(string)), user)
	if err != nil {
		return user
	}
	return user
}


func GenerateToken(username string, id int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(setting.ExpireTime * time.Hour)
	claims := Claims{
		username,
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
