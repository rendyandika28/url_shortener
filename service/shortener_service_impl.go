package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"url_shortener/helper"
	"url_shortener/helper/dto"
	"url_shortener/model/domain"
	"url_shortener/model/web"
	"url_shortener/repository"
)

type ShortenerServiceImpl struct {
	ShortenerRepository repository.ShortenerRepository
	DB                  *gorm.DB
	Validate            *validator.Validate
}

func NewShortenerService(shortenerRepository repository.ShortenerRepository, DB *gorm.DB, validate *validator.Validate) ShortenerService {
	return &ShortenerServiceImpl{ShortenerRepository: shortenerRepository, DB: DB, Validate: validate}
}

func (service *ShortenerServiceImpl) Create(ctx context.Context, request dto.ShortenerRequest) dto.ShortenerResponse {
	if err := service.Validate.Struct(request); err != nil {
		panic(err)
	}

	shortener := &domain.UrlShortener{
		Slug: request.Slug,
		Url:  request.Url,
	}

	service.ShortenerRepository.Save(ctx, service.DB, shortener)
	return dto.ToShortenerResponse(*shortener)
}

func (service *ShortenerServiceImpl) Update(ctx context.Context, slug string, request dto.ShortenerUpdateRequest) dto.ShortenerResponse {
	if err := service.Validate.Struct(request); err != nil {
		panic(err)
	}

	response := service.FindBySlug(ctx, slug)

	shortener := &domain.UrlShortener{
		ID: response.ID,
		Slug: func() string {
			if request.Slug != "" {
				return request.Slug
			}
			return response.Slug
		}(),
		Url: func() string {
			if request.Url != "" {
				return request.Url
			}
			return response.Url
		}(),
	}

	service.ShortenerRepository.Update(ctx, service.DB, shortener)
	return dto.ToShortenerResponse(*shortener)
}

func (service *ShortenerServiceImpl) Delete(ctx context.Context, slug string) {
	err := service.DB.Transaction(func(tx *gorm.DB) error {
		shortener, err := service.ShortenerRepository.FindBySlug(ctx, service.DB, slug)
		if err != nil {
			return err
		}
		service.ShortenerRepository.Delete(ctx, service.DB, &shortener)
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func (service *ShortenerServiceImpl) FindBySlug(ctx context.Context, slug string) dto.ShortenerResponse {
	var shortener domain.UrlShortener
	err := service.DB.Transaction(func(tx *gorm.DB) error {
		shortener, _ = service.ShortenerRepository.FindBySlug(ctx, service.DB, slug)
		return nil
	})
	if err != nil {
		panic(err)
	}

	return dto.ToShortenerResponse(shortener)
}

func (service *ShortenerServiceImpl) FindAll(ctx context.Context) ([]dto.ShortenerResponse, web.Pagination) {
	query := ctx.Value("query").(dto.ShortenerQuery)

	query.SortBy = helper.GetOrderQuery(query.SortBy)
	query.Page = helper.SetDefaultValueQuery(query.Page, 1).(int)
	query.PerPage = helper.SetDefaultValueQuery(query.PerPage, 10).(int)

	shorteners, additionalResponse := service.ShortenerRepository.FindAll(ctx, service.DB, query)
	totalPage := int((additionalResponse.TotalData + int64(query.PerPage) - 1) / int64(query.PerPage))

	meta := web.Pagination{
		Page:      query.Page,
		PerPage:   query.PerPage,
		TotalData: additionalResponse.TotalData,
		TotalPage: totalPage,
	}

	return dto.ToShortenerResponses(shorteners), meta
}
