package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/aziz8009/cinema-app/src/entities"
	"github.com/aziz8009/cinema-app/utils/constants"
)

type MovieRepository interface {
	GetAll(ctx context.Context, req constants.MoviesRequest) (res []entities.Movie, count int64, err error)
	GetDetailById(ctx context.Context, id uint) (res entities.Movie, err error)
	Create(ctx context.Context, req entities.Movie) (res entities.Movie, err error)
	Update(ctx context.Context, id uint, req entities.Movie) (res entities.Movie, err error)
	GetMostViewed(ctx context.Context) (res entities.Movie, err error)
	GetMostViewedByGenre(ctx context.Context) (res []entities.Movie, err error)
}

type movieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) MovieRepository {
	if db == nil {
		panic("db is nil")
	}

	return &movieRepository{db: db}
}

func (m *movieRepository) GetAll(ctx context.Context, req constants.MoviesRequest) (res []entities.Movie, count int64, err error) {
	var (
		// offset    = (req.Page - 1) * req.Limit
		args      []interface{}
		countArgs []interface{}
		filters   []string
	)

	query := `
		SELECT m.id, m.name, m.description, m.duration, m.artists, m.genres, m.watch_url, m.views, u.name AS author, m.created_at, m.published
		FROM movies m
		LEFT JOIN users u ON u.id = m.created_by
		WHERE m.deleted_at IS NULL`

	countQuery := `
		SELECT COUNT(*)
		FROM movies m
		LEFT JOIN users u ON u.id = m.created_by
		WHERE m.deleted_at IS NULL`

	// Apply status filters
	if req.Status != nil {
		if *req.Status {
			filters = append(filters, "m.published IS NOT NULL AND m.published != ''")
		} else {
			filters = append(filters, "m.published IS NULL OR m.published = ''")
		}
	}

	// Apply keyword filters
	if req.Keyword != "" {
		filters = append(filters, fmt.Sprintf("(LOWER(m.name) LIKE $%d OR LOWER(m.artists) LIKE $%d)", len(args)+1, len(args)+1))
		keyword := "%" + strings.ToLower(req.Keyword) + "%"
		args = append(args, keyword)
		countArgs = append(countArgs, keyword)
	}

	// Combine filters into WHERE clause
	if len(filters) > 0 {
		filterClause := " AND " + strings.Join(filters, " AND ")
		query += filterClause
		countQuery += filterClause
	}

	// Append ordering and pagination
	// query += fmt.Sprintf(" ORDER BY m.created_at DESC LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	// args = append(args, req.Limit, offset)
	// Execute count query
	err = m.db.QueryRowContext(ctx, countQuery, countArgs...).Scan(&count)
	if err != nil {
		return nil, 0, fmt.Errorf("error fetching total count: %w", err)
	}

	// Execute main query
	rows, err := m.db.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, 0, fmt.Errorf("error fetching paginated data: %w", err)
	}
	defer rows.Close()

	// Scan rows into results
	for rows.Next() {
		var movie entities.Movie
		err = rows.Scan(
			&movie.ID, &movie.Name, &movie.Description, &movie.Duration,
			&movie.Artists, &movie.Genres, &movie.WatchURL, &movie.Views,
			&movie.Author, &movie.CreatedAt, &movie.Published,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("error scanning movie data: %w", err)
		}
		res = append(res, movie)
	}

	// Check for errors during row iteration
	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating rows: %w", err)
	}

	return
}

func (m movieRepository) GetDetailById(ctx context.Context, id uint) (res entities.Movie, err error) {

	rows := m.db.QueryRowContext(ctx,
		`SELECT m.id, m.name, m.description, m.duration, m.artists, m.genres, m.watch_url, m.views, u.name AS author, m.created_at, m.published
		FROM movies m
		LEFT JOIN users u ON u.id = m.created_by
		WHERE m.deleted_at IS NULL AND m.id = ?`,
		id,
	)

	err = rows.Scan(
		&res.ID, &res.Name, &res.Description, &res.Duration,
		&res.Artists, &res.Genres, &res.WatchURL, &res.Views,
		&res.Author, &res.CreatedAt, &res.Published,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return res, fmt.Errorf("movie not found with id: %s", id)
		}
		return res, fmt.Errorf("error scanning movie data: %v", err)
	}

	return
}

func (m *movieRepository) Create(ctx context.Context, req entities.Movie) (res entities.Movie, err error) {

	query, args, err := sq.Insert("movies").
		Columns("name", "description", "duration", "artists", "genres", "watch_url", "created_by", "created_at").
		Values(req.Name, req.Description, req.Duration, req.Artists, req.Genres, req.WatchURL, req.CreatedBy, req.CreatedAt).
		ToSql()

	if err != nil {
		return res, err
	}

	// Execute the query
	result, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return res, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.ID = uint(lastID)

	return
}

func (m *movieRepository) Update(ctx context.Context, id uint, req entities.Movie) (res entities.Movie, err error) {

	query, args, err := sq.Update("movies").
		Set("name", req.Name).
		Set("description", req.Description).
		Set("duration", req.Duration).
		Set("artists", req.Artists).
		Set("genres", req.Genres).
		Set("watch_url", req.WatchURL).
		Set("updated_by", req.UpdatedBy).
		Set("updated_at", req.UpdatedAt).
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return res, err
	}

	updated, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return res, err
	}

	lastID, err := updated.LastInsertId()
	if err != nil {
		return res, err
	}

	res.ID = uint(lastID)

	return
}

func (m *movieRepository) GetMostViewed(ctx context.Context) (res entities.Movie, err error) {

	query := `SELECT 
		m.id, 
		m.name, 
		m.description, 
		m.artists, 
		m.genres, 
		m.watch_url, 
		COUNT(vt.movie_id) AS views,
		SUM(vt.view_duration) AS total_view_duration,
		u.name AS author, 
		m.published,
		m.created_at,
		m.updated_at
	FROM viewership_tracking  as vt
	JOIN movies m ON vt.movie_id = m.id
	JOIN users u ON u.id = m.created_by
	GROUP BY m.id
	ORDER BY views DESC
	LIMIT ?
	`
	rows := m.db.QueryRowContext(ctx, query, 1)

	err = rows.Scan(
		&res.ID, &res.Name, &res.Description,
		&res.Artists, &res.Genres, &res.WatchURL, &res.Views, &res.ViewTotalDuration,
		&res.Author, &res.Published, &res.CreatedAt, &res.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return res, fmt.Errorf("most viewed movie not found")
		}
		return res, fmt.Errorf("error scanning most viewed movie data: %v", err)
	}

	return
}

func (m *movieRepository) GetMostViewedByGenre(ctx context.Context) (res []entities.Movie, err error) {

	query := `SELECT 
		m.id, 
		m.name, 
		m.description, 
		m.artists, 
		m.genres, 
		m.watch_url, 
		COUNT(vt.movie_id) AS views,
		SUM(vt.view_duration) AS total_view_duration,
		u.name AS author, 
		m.published,
		m.created_at,
		m.updated_at
	FROM viewership_tracking AS vt
	JOIN movies m ON vt.movie_id = m.id
	JOIN users u ON u.id = m.created_by
	GROUP BY m.id ORDER BY views DESC LIMIT 1`

	rows, err := m.db.QueryContext(ctx, query)

	if err != nil {
		return res, fmt.Errorf("error fetching paginated data: %w", err)
	}
	defer rows.Close()

	// Scan rows into results
	for rows.Next() {
		var movie entities.Movie
		err = rows.Scan(
			&movie.ID, &movie.Name, &movie.Description,
			&movie.Artists, &movie.Genres, &movie.WatchURL, &movie.Views, &movie.ViewTotalDuration,
			&movie.Author, &movie.Published, &movie.CreatedAt, &movie.UpdatedAt,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return res, fmt.Errorf("most viewed movie not found")
			}
			return res, fmt.Errorf("error scanning most viewed movie data: %v", err)
		}

		res = append(res, movie)
	}

	// Check for errors during row iteration
	if err = rows.Err(); err != nil {
		return res, fmt.Errorf("error iterating rows: %w", err)
	}

	return
}
