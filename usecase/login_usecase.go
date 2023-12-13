package usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
)

type LoginUC interface {
	Login(username string, pass string) (model.TokenDetails, error) 
}

type loginUC struct {
	repo repository.LoginRepo
}

func (l *loginUC) Login(username string, pass string) (model.TokenDetails, error) {
	return l.repo.Login(username, pass)
}

func NewLoginUC(repo repository.LoginRepo) LoginUC {
	return &loginUC{
		repo: repo,
	}
}