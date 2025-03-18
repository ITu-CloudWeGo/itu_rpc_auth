package main

import (
	"context"
	auth_service "github.com/ITu-CloudWeGo/itu_rpc_auth/rpc/kitex_gen/auth_service"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// CreateAuth implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) CreateAuth(ctx context.Context, req *auth_service.CreateAuthReq) (resp *auth_service.CreateAuthResp, err error) {
	// TODO: Your code here...
	return
}

// UpdateAuth implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) UpdateAuth(ctx context.Context, req *auth_service.UpdateAuthReq) (resp *auth_service.UpdateAuthResp, err error) {
	// TODO: Your code here...
	return
}

// Email implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) Email(ctx context.Context, req *auth_service.EmailReq) (resp *auth_service.EmailResp, err error) {
	// TODO: Your code here...
	return
}
