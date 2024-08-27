package service

import "my-ecom.com/api/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func IUserService() *UserService {
	return &UserService{
		userRepo: repo.IUserRepo(),
	}
}

func (us *UserService) GetInfoUser() string {
	return us.userRepo.GetInfoUser()
}
