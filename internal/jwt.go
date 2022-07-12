package internal

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

const (
	// KEY ...
	KEY string = "JWT-TOKEN"
	// ExprieTime 过期时间30minutes
	ExprieTime int = 60 * 24
	// ISSUER 颁发者
	ISSUER = "mammoth"
)

type MyClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func (c MyClaims) Valid() error {
	vErr := new(jwt.ValidationError)
	now := jwt.TimeFunc().Unix()

	// The claims below are optional, by default, so if they are set to the
	// default value in Go, let's not fail the verification for them.
	if c.VerifyExpiresAt(now, false) == false {
		delta := time.Unix(now, 0).Sub(time.Unix(c.ExpiresAt, 0))
		vErr.Inner = fmt.Errorf("token is expired by %v", delta)
		vErr.Errors |= jwt.ValidationErrorExpired
	}

	if c.VerifyIssuedAt(now, false) == false {
		vErr.Inner = fmt.Errorf("token used before issued")
		vErr.Errors |= jwt.ValidationErrorIssuedAt
	}

	if c.VerifyNotBefore(now, false) == false {
		vErr.Inner = fmt.Errorf("token is not valid yet")
		vErr.Errors |= jwt.ValidationErrorNotValidYet
	}

	if vErr.Errors == 0 {
		return nil
	}

	return vErr
}

// CreateToken 用于签发jwt token, role角色名，expireSecond过期时间, eg:user/admin来区分前后台用户
func CreateToken(userID uint) (string, error) {
	now := time.Now()
	claims := MyClaims{
		userID,
		jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			NotBefore: now.Unix(),
			ExpiresAt: now.Add(time.Second * time.Duration(ExprieTime)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(KEY))
	if err != nil {
		return "", err
	}
	return t, nil
}

func ParseToken(tokenString string) (*MyClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("auth_secret")), nil
	})

	if token == nil {
		return nil, fmt.Errorf("invalid token")
	}

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, fmt.Errorf("token expired")
			}
		}
		return nil, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func GetMyClaims(c *ServiceContext) (*MyClaims, error) {
	token := c.Context.Request.Header.Get("Authorization")
	if token == "" {
		return nil, errors.New("请求头Authorization不能为空")
	}
	myClaims, err := ParseToken(token)
	if err != nil {
		return nil, err
	}
	return myClaims, nil
}
