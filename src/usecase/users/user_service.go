package users

import (
	"github.com/aziz8009/cinema-app/src/entities"
	"github.com/aziz8009/cinema-app/src/repositories"
	"github.com/aziz8009/cinema-app/utils/constants"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	GetAllUsers(c echo.Context) (res constants.DefaultResponse, err error)
	CreateUser(c echo.Context, user UserReq) (res constants.DefaultResponse, err error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService() *userService {
	return &userService{}
}

func (s *userService) SetUserRepo(repo repositories.UserRepository) *userService {
	s.userRepo = repo
	return s
}

func (s *userService) Validate() UserService {
	if s.userRepo == nil {
		panic("userRepo is nil")
	}
	return s
}

func (s *userService) GetAllUsers(c echo.Context) (res constants.DefaultResponse, err error) {
	ctx := c.Request().Context()

	data, err := s.userRepo.GetAll(ctx)

	if err != nil {
		return
	}

	res = constants.DefaultResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data:    data,
		Errors:  make([]string, 0),
	}
	return

}

func (s *userService) CreateUser(c echo.Context, req UserReq) (res constants.DefaultResponse, err error) {

	ctx := c.Request().Context()
	dataInsert := entities.User{
		ID:    req.ID,
		Name:  req.Name,
		Email: req.Email,
	}

	lastInserId, err := s.userRepo.Create(ctx, dataInsert)

	res = constants.DefaultResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data:    lastInserId,
		Errors:  make([]string, 0),
	}
	return
}
