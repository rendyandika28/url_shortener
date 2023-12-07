package repository

import (
	"context"
	"gorm.io/gorm"
	"url_shortener/helper/dto"
	"url_shortener/model/domain"
	"url_shortener/model/web"
)

type ShortenerRepository interface {
	Save(ctx context.Context, db *gorm.DB, shortener *domain.UrlShortener) domain.UrlShortener
	Update(ctx context.Context, db *gorm.DB, shortener *domain.UrlShortener) domain.UrlShortener
	Delete(ctx context.Context, db *gorm.DB, shortener *domain.UrlShortener)
	FindBySlug(ctx context.Context, db *gorm.DB, slug string) (domain.UrlShortener, error)
	FindAll(ctx context.Context, db *gorm.DB, query dto.ShortenerQuery) ([]domain.UrlShortener, web.AddsDatabaseResponse)
	CountAll(ctx context.Context, db *gorm.DB) int64
}
