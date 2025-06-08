package gotool

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims 自定义的 JWT Claims 结构
type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// JWTConfig JWT 配置结构
type JWTConfig struct {
	SecretKey     string        // JWT 密钥
	ExpireTime    time.Duration // 过期时间
	Issuer        string        // 签发者
	SigningMethod string        // 签名方法
}

// DefaultJWTConfig 返回默认的 JWT 配置
func DefaultJWTConfig() *JWTConfig {
	return &JWTConfig{
		SecretKey:     "your-secret-key",
		ExpireTime:    24 * time.Hour,
		Issuer:        "gotool",
		SigningMethod: "HS256",
	}
}

// GenerateToken 生成 JWT token
func GenerateToken(config *JWTConfig, userID uint, username string) (string, error) {
	if config == nil {
		config = DefaultJWTConfig()
	}

	// 创建 claims
	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.ExpireTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    config.Issuer,
		},
	}

	// 创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名 token
	tokenString, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析 JWT token
func ParseToken(config *JWTConfig, tokenString string) (*JWTClaims, error) {
	if config == nil {
		config = DefaultJWTConfig()
	}

	// 解析 token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// 验证 token
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// RefreshToken 刷新 JWT token
func RefreshToken(config *JWTConfig, tokenString string) (string, error) {
	// 解析原 token
	claims, err := ParseToken(config, tokenString)
	if err != nil {
		return "", err
	}

	// 生成新 token
	return GenerateToken(config, claims.UserID, claims.Username)
}

// ValidateToken 验证 JWT token 是否有效
func ValidateToken(config *JWTConfig, tokenString string) bool {
	_, err := ParseToken(config, tokenString)
	return err == nil
}
