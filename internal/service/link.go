package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"jobsearcher_user/config"
	ipsql "jobsearcher_user/internal/app/repository/psql"
	isrvc "jobsearcher_user/internal/app/service"
	"jobsearcher_user/pkg/client"
)

type link struct {
	repLink ipsql.Link
	tx      client.Manager

	cfg *config.Config
}

func NewLink(cfg *config.Config, repLink ipsql.Link, tx client.Manager) isrvc.Link {
	return &link{
		cfg:     cfg,
		repLink: repLink,
		tx:      tx,
	}
}
func (l link) createHash(username string) (string, error) {
	mac := hmac.New(sha256.New, []byte(l.cfg.APIToken))
	_, err := mac.Write([]byte(username))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}
func (l link) Create(ctx context.Context, username, accessToken string) (string, error) {
	var err error
	hash, err := l.createHash(username)
	if err != nil {
		return "", err
	}
	if err = l.repLink.Create(ctx, hash, accessToken); err != nil {
		return "", err
	}
	return hash, err
}

func (l link) Claim(ctx context.Context, hash string) (string, error) {
	var err error
	var accessToken string
	err = l.tx.Do(ctx, func(ctx context.Context) error {
		var txErr error
		accessToken, txErr = l.repLink.GetAccessToken(ctx, hash)
		if txErr != nil {
			return txErr
		}
		if txErr = l.repLink.Delete(ctx, hash); err != nil {
			return txErr
		}
		return nil

	})

	return accessToken, err
}
