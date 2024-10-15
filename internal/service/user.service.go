package service

import "go-tip/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

// us := &UserService{}
func (us *UserService) GetUser() string {
	return us.userRepo.GetUser()
}