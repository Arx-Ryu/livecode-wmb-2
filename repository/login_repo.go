package repository

import (
	"errors"
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/utils"
)

type LoginRepo interface {
	Login(username string, pass string) (model.TokenDetails, error) 
}

type loginRepo struct {
	token utils.Token
}

func (l *loginRepo) Login(username string, pass string) (model.TokenDetails, error) {
	var user model.Credential
	var token *model.TokenDetails
	if username == "admin@gmail.com" && pass == "12345678" {
		token, err := l.token.CreateAccessToken(&user)
		if err != nil {
			return *token, err
		}
		l.token.StoreAccessToken(username, token)
		return *token, nil
	}
	err := errors.New("wrong username or password")
	return *token, err
}

func NewLoginRepository(token utils.Token) LoginRepo {
	repo := new(loginRepo)
	repo.token = token
	return repo
}