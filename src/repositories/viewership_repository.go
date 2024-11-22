package repositories

import (
	"context"
	"database/sql"

	"github.com/aziz8009/cinema-app/src/entities"
)

type ViewerShipRepository interface {
	GetAll(ctx context.Context) (res []entities.View, err error)
}

type viewerShipRepository struct {
	db *sql.DB
}

func NewViewerShipRepository(db *sql.DB) ViewerShipRepository {
	if db == nil {
		panic("db is nil")
	}

	return &viewerShipRepository{db: db}
}

func (m *viewerShipRepository) GetAll(ctx context.Context) (res []entities.View, err error) {
	return
}
