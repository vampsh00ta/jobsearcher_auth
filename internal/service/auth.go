package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"jobsearcher_user/config"
	isrvc "jobsearcher_user/internal/app/service"
	//"github.com/dgrijalva/jwt-go"
	"jobsearcher_user/internal/entity"
	"time"
)

type auth struct {
	cfg *config.Config
}

type jwtClaim struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	PhotoUrl  string `json:"photo_url"`
	jwt.RegisteredClaims
}

func NewAuth(cfg *config.Config) isrvc.Auth {
	return &auth{
		cfg: cfg,
	}
}
func (a auth) CreateToken(ctx context.Context, user entity.User) (string, error) {
	userClaims := jwtClaim{
		user.ID,
		user.FirstName,
		user.LastName,
		user.Username,
		user.PhotoUrl,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	access, err := token.SignedString([]byte(a.cfg.APIToken))
	if err != nil {
		return "", err
	}
	return access, nil

}
func (a auth) VerifyToken(_ context.Context, accessToken string) (bool, error) {
	token, err := jwt.ParseWithClaims(accessToken, &jwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.cfg.APIToken), nil
	})
	if err != nil {
		return false, err
	}
	_, ok := token.Claims.(*jwtClaim)
	if !ok {
		return false, errors.New("wrong or expired data")
	}
	return true, nil

}
