package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"sigmatech-kredit-plus/internal/middleware"
	"sigmatech-kredit-plus/pkg"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func generateTestToken(consumerID, adminID, role string) string {
	_ = os.Setenv("JWT_KEY", "test_secret")
	claims := pkg.NewToken(consumerID, adminID, role)
	claims.RegisteredClaims.ExpiresAt = nil // optional → biar nggak expired di test
	token, _ := claims.Generate()
	return token
}

func setupTestRouter(role string) (*gin.Engine, *httptest.ResponseRecorder, *http.Request) {
	r := gin.New()
	r.Use(middleware.Auth(role))
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})
	return r, httptest.NewRecorder(), nil
}

func TestAuthMiddleware_ValidTokenAndRole(t *testing.T) {
	token := generateTestToken("consumer-id-test", "", "consumer")
	r, w, _ := setupTestRouter("consumer")

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "success")
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	r, w, _ := setupTestRouter("consumer")

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalid.token.here")

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "invalid or expired token")
}

func TestAuthMiddleware_InvalidRole(t *testing.T) {
	token := generateTestToken("consumer-id-test", "", "consumer")
	r, w, _ := setupTestRouter("admin") // expecting "admin", got "consumer"

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "you not have permission")
}

func TestAuthMiddleware_MissingHeader(t *testing.T) {
	r, w, _ := setupTestRouter("consumer")

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "authorization header missing")
}

func TestAuthMiddleware_InvalidHeaderType(t *testing.T) {
	r, w, _ := setupTestRouter("consumer")

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Token abc.def.ghi")

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "invalid header type")
}
