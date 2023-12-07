package repository

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/url"
	"time"
	"url_shortener/helper/dto"
	"url_shortener/model/domain"
	"url_shortener/model/web"
)

type ShortenerRepositoryImpl struct{}

func NewShortenerRepository() ShortenerRepository {
	return &ShortenerRepositoryImpl{}
}

func (repository *ShortenerRepositoryImpl) Save(ctx context.Context, db *gorm.DB, shortener *domain.UrlShortener) domain.UrlShortener {
	if err := db.Create(&shortener).Error; err != nil {
		panic(err)
	}
	return *shortener
}

func (repository *ShortenerRepositoryImpl) Update(ctx context.Context, db *gorm.DB, shortener *domain.UrlShortener) domain.UrlShortener {
	if err := db.Updates(&shortener).Error; err != nil {
		panic(err)
	}
	return *shortener
}

func (repository *ShortenerRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, shortener *domain.UrlShortener) {
	if err := db.Delete(&shortener).Error; err != nil {
		panic(err)
	}
}

func (repository *ShortenerRepositoryImpl) FindBySlug(ctx context.Context, db *gorm.DB, slug string) (domain.UrlShortener, error) {
	var shortener domain.UrlShortener
	var params string

	if _, err := uuid.Parse(slug); err == nil {
		params = "id = ?"
	} else {
		params = "slug = ?"
	}

	if err := db.Take(&shortener, params, slug).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			unescape, _ := url.QueryUnescape(slug)
			panic(fiber.NewError(fiber.StatusNotFound, "There is no data with value: "+unescape))
		}
		panic(err)
	}

	if err := db.Model(&shortener).UpdateColumn("last_visited", time.Now()).Error; err != nil {
		panic(err)
	}

	return shortener, nil
}

func (repository *ShortenerRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB, query dto.ShortenerQuery) ([]domain.UrlShortener, web.AddsDatabaseResponse) {
	var shorteners []domain.UrlShortener
	offset := (query.Page - 1) * query.PerPage

	queryDb := db.Where("slug like ?", "%"+query.Search+"%").Order(query.SortBy).Offset(offset).Limit(query.PerPage)

	if err := queryDb.Find(&shorteners).Error; err != nil {
		panic(err)
	}

	count := repository.CountAll(ctx, queryDb)

	additionalResponse := web.AddsDatabaseResponse{TotalData: count}
	return shorteners, additionalResponse
}

func (repository *ShortenerRepositoryImpl) CountAll(ctx context.Context, db *gorm.DB) int64 {
	var count int64
	if err := db.Model(&domain.UrlShortener{}).Count(&count).Error; err != nil {
		panic(err)
	}

	return count
}
