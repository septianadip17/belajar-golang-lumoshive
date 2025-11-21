package service

type userService struct{}

// constractor
func NewUserService() userService {
	return userService{}
}

func (*userService) RegisterUser() (model.user, error) {

	// logic register user
	return model.user{}, nil
}
