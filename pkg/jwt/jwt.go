package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个UserId字段，所以要自定义结构体
// 如果想要保存更多的信息，都可以添加到这个结构体中
type MyClaims struct {
	UserId uint64 `json:"user_id"`
	jwt.StandardClaims
}

var mySecret = []byte("www.jsxz.com")

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

// GenToken 生成access token 和 refresh token
func GenToken(userId uint64) (aToken, rToken string, err error) {
	myClaims := MyClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
			Issuer:    "jsxz", //签发人
		},
	}

	//加密并获的完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodES256, myClaims).SignedString(mySecret)
	if err != nil {
		return
	}

	//refresh token 不需要任何自定字段
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		Issuer:    "jsxz",
	}).SignedString(mySecret)

	return
}

// 解析jwt
func ParseToken(tokenString string) (claims *MyClaims, err error) {
	var token *jwt.Token
	claims = new(MyClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return
	}

	if !token.Valid {
		err = errors.New("invalid token")
	}

	return
}

// 更新token
func RefreshToken(aToken, rToken string) (newRToken, newAToken string, err error) {
	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
		return
	}

	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserId)
	}

	return
}
