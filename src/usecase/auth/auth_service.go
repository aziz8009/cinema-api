package auth

import (
	"errors"

	"github.com/aziz8009/cinema-app/src/entities"
	"github.com/aziz8009/cinema-app/src/repositories"
	"github.com/aziz8009/cinema-app/utils"
	"github.com/aziz8009/cinema-app/utils/constants"
	"github.com/labstack/echo/v4"
)

type AuthService interface {
	Login(c echo.Context, req LoginRequest) (res constants.DefaultResponse, err error)
	Register(c echo.Context, req RegisterRequest) (res constants.DefaultResponse, err error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService() *authService {
	return &authService{}
}

func (a *authService) SetUserRepo(userRepo repositories.UserRepository) *authService {
	a.userRepo = userRepo
	return a
}

func (a *authService) Validate() AuthService {
	if a.userRepo == nil {
		panic("userRepo is nil")
	}
	return a
}

func (a *authService) Login(c echo.Context, req LoginRequest) (res constants.DefaultResponse, err error) {
	ctx := c.Request().Context()

	user, err := a.userRepo.GetByEmail(ctx, req.Email)

	if err != nil {
		return res, err
	}

	isValid, err := utils.Decrypt([]byte(user.Password), []byte(req.Password))

	if !isValid {
		err = errors.New("invalid email or password")
		return res, err
	}

	token, err := utils.GenerateToken(user.ID)

	res = constants.DefaultResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data: LoginResponse{
			Token: token,
		},
		Errors: make([]string, 0),
	}

	return
}

func (a *authService) Register(c echo.Context, req RegisterRequest) (res constants.DefaultResponse, err error) {

	ctx := c.Request().Context()
	// Password Hashing
	hashedPassword, err := utils.Encrypt(req.Password)
	if err != nil {
		return
	}

	userCheck, err := a.userRepo.GetByEmail(ctx, req.Email)

	if err != nil {
		return res, err
	}

	if userCheck.ID > 0 {
		err = errors.New("user already exists email")
		return res, err
	}

	dataInsert := entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	user, err := a.userRepo.Create(ctx, dataInsert)

	res = constants.DefaultResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data:    user,
		Errors:  make([]string, 0),
	}
	return
}