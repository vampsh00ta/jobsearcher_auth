package grpc

import (
	"go.uber.org/zap"
	isrvc "jobsearcher_auth/internal/app/service"
	"jobsearcher_auth/internal/pb"
)

type authGRPC struct {
	pb.UnimplementedAuthServer

	l        *zap.Logger
	authSrvc isrvc.Auth
	linkSrvc isrvc.Link
}

func New(authSrvc isrvc.Auth, linkSrvc isrvc.Link, l *zap.Logger) *authGRPC {
	return &authGRPC{
		authSrvc: authSrvc,
		linkSrvc: linkSrvc,
		l:        l,
	}
}
