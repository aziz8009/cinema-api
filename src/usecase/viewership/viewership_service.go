package viewership

import (
	"github.com/aziz8009/cinema-app/src/repositories"
	"github.com/aziz8009/cinema-app/utils/constants"
	"github.com/labstack/echo/v4"
)

type ViewerShipService interface {
	GetAll(c echo.Context) (res constants.DefaultResponse, err error)
}

type viewerShipService struct {
	viewerShipRepo repositories.ViewerShipRepository
}

func NewViewerShipService() *viewerShipService {
	return &viewerShipService{}
}

func (m *viewerShipService) SetViewerShipRepo(viewerShipRepo repositories.ViewerShipRepository) *viewerShipService {
	m.viewerShipRepo = viewerShipRepo

	return m
}

func (m *viewerShipService) Validate() ViewerShipService {
	if m.viewerShipRepo == nil {
		panic("movieRepo is nil")
	}

	return m
}

func (m *viewerShipService) GetAll(c echo.Context) (res constants.DefaultResponse, err error) {
	return
}
