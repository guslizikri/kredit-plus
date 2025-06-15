package pkg

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type claims struct {
	ConsumerId string `json:"consumerId"`
	AdminId    string `json:"adminId"`
	Role       string `json:"role"`
	jwt.RegisteredClaims
}

func NewToken(consumerId, adminId, role string) *claims {
	return &claims{
		ConsumerId: consumerId,
		AdminId:    adminId,
		Role:       role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "FWG023",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 48)),
		},
	}
}

func (c *claims) Generate() (string, error) {
	secret := os.Getenv("JWT_KEY")
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return tokens.SignedString([]byte(secret))
}

func VerifyToken(token string) (*claims, error) {
	secret := os.Getenv("JWT_KEY")
	data, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claimData := data.Claims.(*claims)
	return claimData, nil

}
