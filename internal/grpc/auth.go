package grpc

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	isrvc "jobsearcher_auth/internal/app/service"
	"jobsearcher_auth/internal/entity"
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
func Register(s *grpc.Server, auth *authGRPC) {
	pb.RegisterAuthServer(s, auth)
}

func (u authGRPC) VerifyToken(ctx context.Context, req *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	methodName := "VerifyToken"

	ok, err := u.authSrvc.VerifyToken(ctx, req.AccessToken)
	if err != nil {
		err = status.Error(codes.Unauthenticated, err.Error())
		u.l.Error(methodName, zap.Error(err))
		return nil, err
	}
	u.l.Info(methodName)
	return &pb.VerifyTokenResponse{Status: ok}, nil
}

func (u authGRPC) AcceptToken(ctx context.Context, req *pb.AcceptTokenRequest) (*pb.AcceptTokenResponse, error) {
	methodName := "AcceptToken"

	access, err := u.linkSrvc.Claim(ctx, req.Hash)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		u.l.Error(methodName, zap.Error(err))
		return nil, err
	}

	u.l.Info(methodName)
	return &pb.AcceptTokenResponse{AccessToken: access}, nil
}

func (u authGRPC) CreateLink(ctx context.Context, req *pb.CreateLinkRequest) (*pb.CreateLinkResponse, error) {
	methodName := "CreateLink"

	user := entity.User{
		int(req.ID),
		nilToString(req.FirstName),
		nilToString(req.LastName),
		nilToString(req.Username),
		nilToString(req.PhotoUrl),
	}
	access, err := u.authSrvc.CreateToken(ctx, user)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		u.l.Error(methodName, zap.Error(err))
		return nil, err
	}
	hash, err := u.linkSrvc.Create(ctx, nilToString(req.Username), access)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		u.l.Error(methodName, zap.Error(err))
		return nil, err
	}

	u.l.Info(methodName)
	return &pb.CreateLinkResponse{Hash: hash}, nil
}
