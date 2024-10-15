package repo

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

// ur := &UserRepo{}
func (ur *UserRepo) GetUser() string {
	return "Ngon roi! 123"
}