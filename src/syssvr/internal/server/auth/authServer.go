// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"github.com/i-Things/things/src/syssvr/internal/logic/auth"
	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"
)

type AuthServer struct {
	svcCtx *svc.ServiceContext
	sys.UnimplementedAuthServer
}

func NewAuthServer(svcCtx *svc.ServiceContext) *AuthServer {
	return &AuthServer{
		svcCtx: svcCtx,
	}
}

func (s *AuthServer) AuthApiCheck(ctx context.Context, in *sys.CheckAuthReq) (*sys.Response, error) {
	l := authlogic.NewAuthApiCheckLogic(ctx, s.svcCtx)
	return l.AuthApiCheck(in)
}

func (s *AuthServer) AuthApiMultiUpdate(ctx context.Context, in *sys.AuthApiMultiUpdateReq) (*sys.Response, error) {
	l := authlogic.NewAuthApiMultiUpdateLogic(ctx, s.svcCtx)
	return l.AuthApiMultiUpdate(in)
}

func (s *AuthServer) AuthApiIndex(ctx context.Context, in *sys.AuthApiIndexReq) (*sys.AuthApiIndexResp, error) {
	l := authlogic.NewAuthApiIndexLogic(ctx, s.svcCtx)
	return l.AuthApiIndex(in)
}