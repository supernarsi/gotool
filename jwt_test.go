package gotool

import (
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	// 创建测试配置
	config := &JWTConfig{
		SecretKey:     "test-secret-key",
		ExpireTime:    1 * time.Hour,
		Issuer:        "test",
		SigningMethod: "HS256",
	}

	// 测试生成 token
	userID := uint(123)
	username := "testuser"
	token, err := GenerateToken(config, userID, username)
	if err != nil {
		t.Errorf("GenerateToken failed: %v", err)
	}

	// 测试解析 token
	claims, err := ParseToken(config, token)
	if err != nil {
		t.Errorf("ParseToken failed: %v", err)
	}

	// 验证 claims
	if claims.UserID != userID {
		t.Errorf("UserID mismatch: got %v, want %v", claims.UserID, userID)
	}
	if claims.Username != username {
		t.Errorf("Username mismatch: got %v, want %v", claims.Username, username)
	}

	// 测试 token 验证
	if !ValidateToken(config, token) {
		t.Error("ValidateToken failed for valid token")
	}

	// 测试无效 token
	if ValidateToken(config, "invalid-token") {
		t.Error("ValidateToken passed for invalid token")
	}

	// 测试 token 刷新
	newToken, err := RefreshToken(config, token)
	if err != nil {
		t.Errorf("RefreshToken failed: %v", err)
	}

	// 验证新 token
	newClaims, err := ParseToken(config, newToken)
	if err != nil {
		t.Errorf("ParseToken failed for refreshed token: %v", err)
	}

	// 验证新 token 的 claims
	if newClaims.UserID != userID {
		t.Errorf("Refreshed token UserID mismatch: got %v, want %v", newClaims.UserID, userID)
	}
	if newClaims.Username != username {
		t.Errorf("Refreshed token Username mismatch: got %v, want %v", newClaims.Username, username)
	}
}

func TestDefaultJWTConfig(t *testing.T) {
	config := DefaultJWTConfig()
	if config == nil {
		t.Error("DefaultJWTConfig returned nil")
	}
	if config.SecretKey == "" {
		t.Error("DefaultJWTConfig SecretKey is empty")
	}
	if config.ExpireTime == 0 {
		t.Error("DefaultJWTConfig ExpireTime is zero")
	}
	if config.Issuer == "" {
		t.Error("DefaultJWTConfig Issuer is empty")
	}
	if config.SigningMethod == "" {
		t.Error("DefaultJWTConfig SigningMethod is empty")
	}
}
