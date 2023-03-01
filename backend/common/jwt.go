package common

import (
	"atm/global"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claim struct {
	Uid int64 `json:"uid"`
	jwt.RegisteredClaims
}

func GenToken(uid int64) (string, error) {
	signingKey := []byte(global.Conf.JwtConf.SigningKey)
	expireTime := global.Conf.JwtConf.ExpireTime
	claim := Claim{uid, jwt.RegisteredClaims{
		Issuer: "atm",
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Duration(expireTime) * time.Second)},
	}}
	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claim).SignedString(signingKey)
	return token, err
}

func VerifyToken(token string) (int64, error) {
	t, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Conf.JwtConf.SigningKey), nil
	})
	claims, ok := t.Claims.(*Claim)
	if ok && t.Valid {
		return claims.Uid, nil
	}
	return 0, err
}
