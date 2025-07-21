package models

import "time"

type Ad struct {
	ID          string    `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	ImageURL    string    `db:"image_url"`
	Price       float64   `db:"price"`
	AuthorID    string    `db:"author_id"`
	CreatedAt   time.Time `db:"created_at"`
}

type AdWithAuthor struct {
	Ad
	IsMine bool `db:"-"`
}

type AdsFilter struct {
	MinPrice *float64
	MaxPrice *float64
	SortBy   string // "created_at" | "price"
	SortDesc bool
	Limit    int
	Offset   int
}
