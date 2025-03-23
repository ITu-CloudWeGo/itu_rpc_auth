package main

import (
	"context"
	auth_service "github.com/ITu-CloudWeGo/itu_rpc_auth/rpc/kitex_gen/auth_service"
	"github.com/ITu-CloudWeGo/itu_rpc_auth/rpc/service/email"
	"github.com/ITu-CloudWeGo/itu_rpc_auth/rpc/service/jwt"
	"log"
	"time"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// GenerateToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) GenerateToken(ctx context.Context, req *auth_service.GenerateTokenRequest) (resp *auth_service.GenerateTokenResponse, err error) {
	// TODO: Your code here...
	accessToken, accessExpire, err := jwt.GenerateAccessToken(req.Uid)
	if err != nil {
		log.Printf("generate access token failed: %v", err)
		return nil, err
	}

	refreshToken, refreshExpire, err := jwt.GenerateRefreshToken(req.Uid)
	if err != nil {
		log.Printf("generate refresh token failed: %v", err)
		return nil, err
	}

	return &auth_service.GenerateTokenResponse{
		AccessToken:        accessToken,
		RefreshToken:       refreshToken,
		AccessTokenExpire:  accessExpire,
		RefreshTokenExpire: refreshExpire,
	}, nil
}

// RefreshToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) RefreshToken(ctx context.Context, req *auth_service.RefreshTokenRequest) (resp *auth_service.RefreshTokenResponse, err error) {
	// TODO: Your code here...
	uid, err := jwt.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		log.Printf("validate refresh token failed: %v", err)
		return nil, err
	}

	accessToken, accessExpire, err := jwt.GenerateAccessToken(uid)
	if err != nil {
		log.Printf("generate access token failed: %v", err)
		return nil, err
	}

	refreshToken, refreshExpire, err := jwt.GenerateRefreshToken(uid)
	if err != nil {
		log.Printf("generate refresh token failed: %v", err)
		return nil, err
	}

	return &auth_service.RefreshTokenResponse{
		AccessToken:        accessToken,
		RefreshToken:       refreshToken,
		AccessTokenExpire:  accessExpire,
		RefreshTokenExpire: refreshExpire,
	}, nil
}

// CheckAccessToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) CheckAccessToken(ctx context.Context, req *auth_service.CheckAccessTokenRequest) (resp *auth_service.CheckAccessTokenResponse, err error) {
	// TODO: Your code here...
	accessTokenExpire, err := jwt.CheckAccessToken(req.AccessToken)
	if err != nil {
		log.Printf("validate access token failed: %v", err)
		return nil, err
	}
	isExpired := time.Now().After(accessTokenExpire)

	return &auth_service.CheckAccessTokenResponse{
		CheckAccessExpiresAt: !isExpired, //true 表示过期
	}, nil
}

// CheckRefreshToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) CheckRefreshToken(ctx context.Context, req *auth_service.CheckRefreshTokenRequest) (resp *auth_service.CheckRefreshTokenResponse, err error) {
	// TODO: Your code here...
	refreshTokenExpire, err := jwt.CheckRefreshToken(req.RefreshToken)
	if err != nil {
		log.Printf("validate refresh token failed: %v", err)
		return nil, err
	}
	isExpired := time.Now().After(refreshTokenExpire)

	return &auth_service.CheckRefreshTokenResponse{
		CheckRefreshExpiresAt: !isExpired, //true 表示过期
	}, nil
}

// Email implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) Email(ctx context.Context, req *auth_service.EmailRequest) (resp *auth_service.EmailResponse, err error) {
	// TODO: Your code here...
	captcha, err := email.SendCaptchaEmail(req.Email)
	if err != nil {
		log.Printf("发送验证码邮件失败: email=%s, error=%v", req.Email, err)
		return nil, err
	}

	return &auth_service.EmailResponse{
		Captcha: int64(captcha),
	}, nil
}
