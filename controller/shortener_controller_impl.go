package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"url_shortener/helper/dto"
	"url_shortener/model/web"
	"url_shortener/service"
)

type ShortenerControllerImpl struct {
	ShortenerService service.ShortenerService
}

func NewShortenerController(shortenerService service.ShortenerService) *ShortenerControllerImpl {
	return &ShortenerControllerImpl{ShortenerService: shortenerService}
}

func (controller ShortenerControllerImpl) Create(ctx *fiber.Ctx) error {
	request := new(dto.ShortenerRequest)
	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrUnprocessableEntity
	}

	shortenerResponse := controller.ShortenerService.Create(context.Background(), *request)

	webResponse := web.Response{
		Code:   201,
		Status: "OK",
		Data:   shortenerResponse,
	}

	return ctx.JSON(webResponse)
}

func (controller ShortenerControllerImpl) Update(ctx *fiber.Ctx) error {
	request := new(dto.ShortenerUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrUnprocessableEntity
	}

	slug := ctx.Params("slug")

	shortenerResponse := controller.ShortenerService.Update(context.Background(), slug, *request)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   shortenerResponse,
	}

	return ctx.JSON(webResponse)
}

func (controller ShortenerControllerImpl) Delete(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	controller.ShortenerService.Delete(context.Background(), slug)

	webResponse := web.Response{
		Code:   204,
		Status: "OK",
	}

	return ctx.JSON(webResponse)
}

func (controller ShortenerControllerImpl) FindBySlug(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")

	shortenerResponse := controller.ShortenerService.FindBySlug(context.Background(), slug)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   shortenerResponse,
	}

	return ctx.JSON(webResponse)
}

func (controller ShortenerControllerImpl) FindAll(ctx *fiber.Ctx) error {
	query := new(dto.ShortenerQuery)

	if err := ctx.QueryParser(query); err != nil {
		return err
	}

	fiberCtx := context.WithValue(context.Background(), "query", *query)
	shortenerResponses, meta := controller.ShortenerService.FindAll(fiberCtx)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   shortenerResponses,
		Meta:   meta,
	}

	return ctx.JSON(webResponse)
}
