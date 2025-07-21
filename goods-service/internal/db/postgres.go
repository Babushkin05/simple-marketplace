package db

import (
	"fmt"
	"log"
	"strings"

	"github.com/Babushkin05/simple-marketplace/goods-service/internal/models"
	"github.com/jmoiron/sqlx"
)

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) CreateAd(ad *models.Ad) (*models.Ad, error) {
	query := `
		INSERT INTO ads (title, description, image_url, price, author_login)
		VALUES (:title, :description, :image_url, :price, :author_login)
		RETURNING id, created_at;
	`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, fmt.Errorf("prepare query: %w", err)
	}

	err = stmt.Get(ad, ad)
	if err != nil {
		return nil, fmt.Errorf("insert ad: %w", err)
	}

	return ad, nil
}

func (r *PostgresRepo) ListAds(filter models.AdsFilter) ([]*models.Ad, error) {
	base := `
		SELECT 
			id, title, description, image_url, price, author_login, created_at
		FROM ads
	`
	var args []interface{}
	var whereClauses []string
	var orderBy string = "created_at"

	if filter.MinPrice != nil && *filter.MinPrice != 0 {
		whereClauses = append(whereClauses, "price >= ?")
		args = append(args, *filter.MinPrice)
	}
	if filter.MaxPrice != nil && *filter.MaxPrice != 0 {
		whereClauses = append(whereClauses, "price <= ?")
		args = append(args, *filter.MaxPrice)
	}

	if filter.SortBy == "price" {
		orderBy = "price"
	}
	if filter.SortDesc {
		orderBy += " DESC"
	} else {
		orderBy += " ASC"
	}

	// Pagination
	limit := filter.Limit
	offset := filter.Offset

	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	query := base
	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}
	query += fmt.Sprintf(" ORDER BY %s LIMIT %d OFFSET %d", orderBy, limit, offset)
	var ads []*models.Ad
	if err := r.db.Select(&ads, r.db.Rebind(query), args...); err != nil {
		return nil, fmt.Errorf("list ads: %w", err)
	}

	log.Println(len(ads))

	return ads, nil
}
