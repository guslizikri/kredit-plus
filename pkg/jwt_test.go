package pkg_test

import (
	"os"
	"testing"
	"time"

	"sigmatech-kredit-plus/pkg"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestVerifyToken(t *testing.T) {
	// Set secret dulu untuk test
	_ = os.Setenv("JWT_KEY", "test_secret")

	// Generate token untuk test
	tokenClaims := pkg.NewToken("consumer-id-test", "", "consumer")
	tokenClaims.RegisteredClaims.ExpiresAt = nil // optional: untuk test token tanpa expired
	tokenString, err := tokenClaims.Generate()
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// Test verify token
	claims, err := pkg.VerifyToken(tokenString)
	assert.NoError(t, err)
	assert.Equal(t, "consumer-id-test", claims.ConsumerId)
	assert.Equal(t, "consumer", claims.Role)
}

func TestVerifyToken_Invalid(t *testing.T) {
	_ = os.Setenv("JWT_KEY", "test_secret")

	invalidToken := "this.is.not.a.valid.token"
	_, err := pkg.VerifyToken(invalidToken)
	assert.Error(t, err)
}

func TestVerifyToken_Expired(t *testing.T) {
	_ = os.Setenv("JWT_KEY", "test_secret")

	// Generate expired token
	tokenClaims := pkg.NewToken("consumer-expired", "", "consumer")
	tokenClaims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(-time.Hour)) // expired 1 jam lalu
	tokenString, err := tokenClaims.Generate()
	assert.NoError(t, err)

	_, err = pkg.VerifyToken(tokenString)
	assert.Error(t, err)
}
