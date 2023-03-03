package handler

import (
	"fmt"
	"net/http"

	"github.com/amirhnajafiz/authX/internal/model"
	"github.com/amirhnajafiz/authX/internal/port/http/request"
	"github.com/amirhnajafiz/authX/internal/port/http/response"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// CreateApp for a user.
func (h *Handler) CreateApp(ctx *fiber.Ctx) error {
	userRequest := new(request.NewApp)

	if err := ctx.BodyParser(&userRequest); err != nil {
		h.Logger.Info("body parsing failed", zap.Error(err))

		return fiber.ErrBadRequest
	}

	user, err := h.Repository.Users.GetByEmail(ctx.Locals("email").(string))
	if err != nil {
		h.Logger.Info("user not found", zap.String("email", ctx.Locals("email").(string)))

		return fiber.ErrBadRequest
	}

	appInstance := model.App{
		Name:   userRequest.Name,
		AppKey: uuid.NewString()[:10],
		UserID: user.ID,
	}

	if err := h.Repository.Apps.Create(&appInstance); err != nil {
		h.Logger.Error("failed to create app instance", zap.Error(err))

		return fiber.ErrInternalServerError
	}

	return ctx.SendStatus(http.StatusCreated)
}

// GetUserApps returns all apps.
func (h *Handler) GetUserApps(ctx *fiber.Ctx) error {
	var list []response.App

	user, err := h.Repository.Users.GetByEmail(ctx.Locals("email").(string))
	if err != nil {
		h.Logger.Info("user not found", zap.String("email", ctx.Locals("email").(string)))

		return fiber.ErrBadRequest
	}

	apps, err := h.Repository.Apps.GetByUserID(user.ID)
	if err != nil {
		h.Logger.Error("failed to get apps", zap.Error(err))

		return fiber.ErrInternalServerError
	}

	for _, app := range apps {
		tmp := response.App{
			Name:      app.Name,
			AppKey:    app.AppKey,
			URL:       fmt.Sprintf("%s/api/app/%s/client", ctx.Hostname(), app.AppKey),
			CreatedAt: app.CreatedAt,
		}

		list = append(list, tmp)
	}

	return ctx.JSON(list)
}

// GetSingleApp of a user.
func (h *Handler) GetSingleApp(ctx *fiber.Ctx) error {
	app, err := h.Repository.Apps.GetByKey(ctx.Params("app_key"))
	if err != nil {
		h.Logger.Error("app not found", zap.String("id", ctx.Params("app_key")))

		return fiber.ErrNotFound
	}

	var list []response.Client

	clients, err := h.Repository.Clients.GetAppClients(app.ID)
	if err != nil {
		h.Logger.Error("cannot get clients", zap.Error(err))

		return fiber.ErrInternalServerError
	}

	for _, client := range clients {
		tmp := response.Client{
			Claims:    client.Credentials,
			CreatedAt: client.CreatedAt,
		}

		list = append(list, tmp)
	}

	return ctx.JSON(list)
}
