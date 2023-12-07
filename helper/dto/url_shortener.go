package dto

import (
	"github.com/google/uuid"
	"time"
	"url_shortener/model/domain"
	"url_shortener/model/web"
)

func ToShortenerResponse(shortener domain.UrlShortener) ShortenerResponse {
	return ShortenerResponse{
		ID:          shortener.ID,
		Slug:        shortener.Slug,
		Url:         shortener.Url,
		LastVisited: shortener.LastVisited,
		CreatedAt:   shortener.CreatedAt,
		UpdatedAt:   shortener.UpdatedAt,
	}
}

func ToShortenerResponses(categories []domain.UrlShortener) []ShortenerResponse {
	shortenerResponses := []ShortenerResponse{}
	for _, shortener := range categories {
		shortenerResponses = append(shortenerResponses, ToShortenerResponse(shortener))
	}
	return shortenerResponses
}

type ShortenerQuery struct {
	web.QueryRequest
	LastVisited string `query:"last_visited"`
}

type ShortenerRequest struct {
	Slug string `json:"slug" validate:"required,min=2,max=255,slug"`
	Url  string `json:"url" validate:"required,url"`
}

type ShortenerResponse struct {
	ID          uuid.UUID `json:"id"`
	Slug        string    `json:"slug"`
	Url         string    `json:"url"`
	LastVisited time.Time `json:"last_visited"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ShortenerUpdateRequest struct {
	Slug string `json:"slug" validate:"omitempty,min=2,max=255,slug"`
	Url  string `json:"url" validate:"omitempty,url"`
}
