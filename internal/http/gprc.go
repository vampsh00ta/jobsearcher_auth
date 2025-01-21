package http

import (
	"context"
	"google.golang.org/grpc"
	isrvc "jobsearcher_user/internal/app/service"
	"jobsearcher_user/internal/entity"
	"jobsearcher_user/internal/pb"
)

type authGRPC struct {
	pb.UnimplementedAuthServer

	authSrvc isrvc.Auth
	linkSrvc isrvc.Link
}

func NewAuthGRPC(authSrvc isrvc.Auth, linkSrvc isrvc.Link) *authGRPC {
	return &authGRPC{authSrvc: authSrvc, linkSrvc: linkSrvc}
}

func Register(s *grpc.Server, auth *authGRPC) {
	pb.RegisterAuthServer(s, auth)
}

func (u authGRPC) VerifyToken(ctx context.Context, req *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	ok, err := u.authSrvc.VerifyToken(ctx, req.AccessToken)
	if err != nil {

		return nil, err
	}
	return &pb.VerifyTokenResponse{Status: ok}, nil
}

func (u authGRPC) AcceptToken(ctx context.Context, req *pb.AcceptTokenRequest) (*pb.AcceptTokenResponse, error) {

	access, err := u.linkSrvc.Claim(ctx, req.Hash)
	if err != nil {
		return nil, err
	}

	return &pb.AcceptTokenResponse{AccessToken: access}, nil
}

func (u authGRPC) CreateLink(ctx context.Context, req *pb.CreateLinkRequest) (*pb.CreateLinkResponse, error) {

	user := entity.User{
		int(req.ID),
		nilToString(req.FirstName),
		nilToString(req.LastName),
		nilToString(req.Username),
		nilToString(req.PhotoUrl),
	}
	access, err := u.authSrvc.CreateToken(ctx, user)
	if err != nil {
		return nil, err
	}
	hash, err := u.linkSrvc.Create(ctx, nilToString(req.Username), access)
	if err != nil {
		return nil, err
	}
	return &pb.CreateLinkResponse{Hash: hash}, nil
}
