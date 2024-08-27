package repo

type UserRepo struct{}

func IUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetInfoUser() string {
	return "Toan"
}
