package jwt

import (
	"errors"
	"fmt"
	"github.com/ITu-CloudWeGo/itu_rpc_auth/rpc/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	Uid uint64 `json:"uid"`
	jwt.RegisteredClaims
}

var (
	globalConfig *config.Config
)

func init() {
	var err error
	globalConfig, err = config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
}

func GenerateAccessToken(uid uint64) (string, int64, error) {

	expire := time.Now().Add(time.Duration(globalConfig.TokenExpiry.AccessTokenExpiry) * time.Second).Unix()

	claims := CustomClaims{
		Uid: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expire, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(globalConfig.Secrets.AccessTokenSecret))
	if err != nil {
		return "", 0, err
	}

	return signedToken, expire, nil
}

func GenerateRefreshToken(uid uint64) (string, int64, error) {
	expire := time.Now().Add(time.Duration(globalConfig.TokenExpiry.RefreshTokenExpiry) * time.Second).Unix()

	claims := CustomClaims{
		Uid: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expire, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(globalConfig.Secrets.RefreshTokenSecret))
	if err != nil {
		return "", 0, err
	}

	return signedToken, expire, nil
}
func ValidateRefreshToken(refreshToken string) (uint64, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(globalConfig.Secrets.RefreshTokenSecret), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.Uid, nil
	}

	return 0, errors.New("invalid refresh token")
}

func CheckAccessToken(accessToken string) (time.Time, error) {
	token, err := jwt.ParseWithClaims(accessToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(globalConfig.Secrets.AccessTokenSecret), nil
	})
	if err != nil {
		return time.Time{}, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.ExpiresAt.Time, nil
	}

	return time.Time{}, errors.New("invalid access token")
}

func CheckRefreshToken(refresh string) (time.Time, error) {
	token, err := jwt.ParseWithClaims(refresh, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(globalConfig.Secrets.RefreshTokenSecret), nil
	})
	if err != nil {
		return time.Time{}, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.ExpiresAt.Time, nil
	}

	return time.Time{}, errors.New("invalid access token")
}
