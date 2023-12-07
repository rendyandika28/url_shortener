package service

import (
	"context"
	"url_shortener/helper/dto"
	"url_shortener/model/web"
)

type ShortenerService interface {
	Create(ctx context.Context, request dto.ShortenerRequest) dto.ShortenerResponse
	Update(ctx context.Context, slug string, request dto.ShortenerUpdateRequest) dto.ShortenerResponse
	Delete(ctx context.Context, slug string)
	FindBySlug(ctx context.Context, slug string) dto.ShortenerResponse
	FindAll(ctx context.Context) ([]dto.ShortenerResponse, web.Pagination)
}
