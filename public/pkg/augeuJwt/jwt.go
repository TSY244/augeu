package augeuJwt

import "github.com/golang-jwt/jwt/v4"

const (
	SIGNING_KEY = "1drftvgybhnjkDERFTGYIJMKO@#$%^&*"
)

const (
	RoleUser  = 0
	RoleAgent = 1
)

// Role 0: User, 1: Agent
type Info struct {
	Role       int
	UserInfo   UserInfo
	ClientInfo AgentInfo
}

type UserInfo struct {
	Name string
}

type AgentInfo struct {
	ClientId string
	Uuid     string // windows uuid
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
