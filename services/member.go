package services

import (
	"line-wallet/config"
	"line-wallet/models"
	"line-wallet/repository"
)

type MemberService struct {
	Conf config.Config
	Repo repository.Repository
}

type IMemberService interface {
	FindMemberByLineUserID(line_user_id string) (*models.Member, error)
}

func (s MemberService) FindMemberByLineUserID(line_user_id string) (*models.Member, error) {
	member, err := s.Repo.Member.FindMemberByLineUserID(line_user_id)
	if err != nil {
		return nil, err
	}
	return member, nil
}
