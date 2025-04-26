package utils

import (
	"errors"
	"glossika-assignment/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWT 配置
var jwtConfig config.JWTConfig

// InitJWT 初始化 JWT 配置
func InitJWT(cfg config.JWTConfig) {
	jwtConfig = cfg
}

// Claims 定義 JWT 聲明
type Claims struct {
	UserID string `json:"userId"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT 令牌
func GenerateToken(userID string) (string, time.Time, error) {
	// 使用配置中的過期時間
	expirationTime := time.Now().Add(time.Duration(jwtConfig.ExpireMinutes) * time.Minute)

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用配置中的密鑰
	tokenString, err := token.SignedString([]byte(jwtConfig.Secret))

	return tokenString, expirationTime, err
}

// ParseToken 解析和驗證 JWT 令牌
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 使用配置中的密鑰
		return []byte(jwtConfig.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
