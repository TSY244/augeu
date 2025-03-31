package utils

import "github.com/golang-jwt/jwt/v4"

const (
	SIGNING_KEY = "123asdQWE1qa@WS#ED,./l;'p[]"
)

// Role 0: admin, 1: member
type Info struct {
	Name string
	Role int
}

type InfoClaims struct {
	Info Info
	jwt.RegisteredClaims
}

func NewJwt(info Info) (string, error) {
	infoClaims := InfoClaims{
		Info:             info,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, infoClaims)
	ss, err := token.SignedString([]byte(SIGNING_KEY))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func ParseJwt(ss string) (Info, error) {
	t, err := jwt.ParseWithClaims(ss, &InfoClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SIGNING_KEY), nil
	})
	if err != nil {
		return Info{}, err
	}
	if claims, ok := t.Claims.(*InfoClaims); ok && t.Valid {
		return claims.Info, nil
	}
	return Info{}, nil
}
