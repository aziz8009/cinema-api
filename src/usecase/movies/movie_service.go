package movies

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aziz8009/cinema-app/src/entities"
	"github.com/aziz8009/cinema-app/src/repositories"
	"github.com/aziz8009/cinema-app/utils"
	"github.com/aziz8009/cinema-app/utils/constants"
	"github.com/labstack/echo/v4"
)

type MovieService interface {
	GetAll(c echo.Context) (res constants.DefaultResponse, err error)
	GetDetailById(c echo.Context) (res constants.DefaultResponse, err error)
	Create(c echo.Context) (res constants.DefaultResponse, err error)
	Update(c echo.Context) (res constants.DefaultResponse, err error)
	GetMostViewed(c echo.Context) (res constants.DefaultResponse, err error)
	GetMostViewedByGenre(c echo.Context) (res constants.DefaultResponse, err error)
}

type movieService struct {
	movieRepo repositories.MovieRepository
}

func NewMovieService() *movieService {
	return &movieService{}
}

func (m *movieService) SetMovieRepo(movieRepo repositories.MovieRepository) *movieService {
	m.movieRepo = movieRepo

	return m
}

func (m *movieService) Validate() MovieService {
	if m.movieRepo == nil {
		panic("movieRepo is nil")
	}

	return m
}

func (m *movieService) GetAll(c echo.Context) (res constants.DefaultResponse, err error) {

	var (
		ctx        = c.Request().Context()
		params     constants.MoviesRequest
		pageParam  = c.QueryParam("page")
		limitParam = c.QueryParam("limit")
		status     = c.QueryParam("status")
		keyword    = c.QueryParam("keyword")
	)

	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}

	params.Page = uint(page)

	limit, err := strconv.Atoi(limitParam)

	if err != nil || limit <= 0 {
		limit = 10
	}

	params.Limit = uint(limit)

	if status != "" {
		parsedStatus, err := strconv.ParseBool(status)

		if err != nil {
			parsedStatus = false

		}

		params.Status = &parsedStatus
	}

	if keyword != "" {
		params.Keyword = keyword
	}

	if err = utils.ValidateRequest(c, &params); err != nil {
		return
	}

	movies, count, err := m.movieRepo.GetAll(ctx, params)
	fmt.Println("count : ", count)

	res = constants.DefaultResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data:    movies,
		Errors:  make([]string, 0),
	}

	return
}

func (m *movieService) GetDetailById(c echo.Context) (res constants.DefaultResponse, err error) {
	var (
		ctx = c.Request().Context()
		id  = c.Param("id")
	)

	idParsed, err := strconv.Atoi(id)
	if err != nil || idParsed <= 0 {
		idParsed = 0
	}

	movies, err := m.movieRepo.GetDetailById(ctx, uint(idParsed))

	res = constants.DefaultResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data:    movies,
		Errors:  make([]string, 0),
	}

	return

}

func (m *movieService) Create(c echo.Context) (res constants.DefaultResponse, err error) {

	var (
		ctx = c.Request().Context()
		now = time.Now().UTC()
		req MovieRequest
	)

	authData, err := utils.AuthData(c)
	if err != nil {
		return
	}

	if err = utils.ValidateRequest(c, &req); err != nil {
		return
	}

	dataInsert := entities.Movie{
		Name:        req.Title,
		Description: req.Description,
		Duration:    req.Duration,
		Artists:     req.Artists,
		Genres:      req.Genres,
		WatchURL:    req.WatchURL,
		CreatedBy:   int(authData.ID),
		CreatedAt:   now,
	}

	movie, err := m.movieRepo.Create(ctx, dataInsert)

	dataInsert.ID = movie.ID

	res = constants.DefaultResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data:    dataInsert,
		Errors:  make([]string, 0),
	}
	return
}

func (m *movieService) Update(c echo.Context) (res constants.DefaultResponse, err error) {

	var (
		req MovieRequest
		ctx = c.Request().Context()
		id  = c.Param("id")
		now = time.Now().UTC()
	)

	authData, err := utils.AuthData(c)
	if err != nil {
		return
	}

	if err = utils.ValidateRequest(c, &req); err != nil {
		return
	}

	idParse, err := strconv.Atoi(id)

	if err != nil {
		return
	}

	dataUpdate := entities.Movie{
		Name:        req.Title,
		Description: req.Description,
		Duration:    req.Duration,
		Artists:     req.Artists,
		Genres:      req.Genres,
		WatchURL:    req.WatchURL,
		Published:   req.Published,
		UpdatedBy:   int(authData.ID),
		UpdatedAt:   now,
	}

	movieUpdated, err := m.movieRepo.Update(ctx, uint(idParse), dataUpdate)

	dataUpdate.ID = movieUpdated.ID
	res = constants.DefaultResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data:    dataUpdate,
		Errors:  make([]string, 0),
	}
	return
}

func (m *movieService) GetMostViewed(c echo.Context) (res constants.DefaultResponse, err error) {

	var (
		ctx = c.Request().Context()
	)

	movies, err := m.movieRepo.GetMostViewed(ctx)

	res = constants.DefaultResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data:    movies,
		Errors:  make([]string, 0),
	}

	return
}

func (m *movieService) GetMostViewedByGenre(c echo.Context) (res constants.DefaultResponse, err error) {

	var (
		ctx = c.Request().Context()
	)

	movies, err := m.movieRepo.GetMostViewedByGenre(ctx)

	res = constants.DefaultResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data:    movies,
		Errors:  make([]string, 0),
	}

	return
}
