package api

import (
	"log"
	"net/http"
	"todo/models"

	"todo/app/database"
	"todo/constants"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	apiRepository database.Repository
}

func NewApiHandler(todoRepository database.Repository) *Handler {
	return &Handler{
		apiRepository: todoRepository,
	}
}
func (h *Handler) CreatePatient(ctx *fiber.Ctx) error {
	request := new(models.CreatePatientRequest)
	if err := ctx.BodyParser(&request); err != nil {
		return models.Response(constants.StatusCodeBadRequest, constants.BadRequestMessage, err.Error()).SendResponse(ctx, http.StatusBadRequest)
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Printf("Validation error: Field '%s' failed on the '%s' tag", err.Field(), err.Tag())
		}
		return models.Response(constants.StatusCodeBadRequest, nil, constants.BadRequestMessage).SendResponse(ctx, http.StatusBadRequest)
	}

	if err := h.apiRepository.CreatePatient(ctx.Context(), request); err != nil {
		return models.Response(constants.StatusCodeSystemError, nil, constants.StatusCodeSystemErrorMessage).SendResponse(ctx, http.StatusInternalServerError)
	}
	return models.ResponseSuccess(constants.StatusCodeSuccess, constants.SuccessMessage, nil).SendResponseSuccess(ctx, http.StatusOK)
}

func (h *Handler) UpdatePatient(ctx *fiber.Ctx) error {
	request := new(models.UpdatePatientRequest)
	if err := ctx.BodyParser(&request); err != nil {
		return models.Response(constants.StatusCodeBadRequest, constants.BadRequestMessage, err.Error()).SendResponse(ctx, http.StatusBadRequest)
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		// Log the validation error
		for _, err := range err.(validator.ValidationErrors) {
			log.Printf("Validation error: Field '%s' failed on the '%s' tag", err.Field(), err.Tag())
		}
		return models.Response(constants.StatusCodeBadRequest, nil, constants.BadRequestMessage).SendResponse(ctx, http.StatusBadRequest)
	}

	if err := h.apiRepository.UpdatePatient(ctx.Context(), request); err != nil {
		return models.Response(constants.StatusCodeSystemError, nil, constants.StatusCodeSystemErrorMessage).SendResponse(ctx, http.StatusInternalServerError)
	}
	return models.ResponseSuccess(constants.StatusCodeSuccess, constants.SuccessMessage, nil).SendResponseSuccess(ctx, http.StatusOK)
}

func (h *Handler) ReadPatient(ctx *fiber.Ctx) error {
	request := new(models.ResponseReadPatient)
	if err := ctx.BodyParser(&request); err != nil {
		return models.Response(constants.StatusCodeBadRequest, constants.BadRequestMessage, err.Error()).SendResponse(ctx, http.StatusBadRequest)
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return models.Response(constants.StatusCodeBadRequest, nil, constants.BadRequestMessage).SendResponse(ctx, http.StatusBadRequest)
	}
	data, err := h.apiRepository.ReadPatient(ctx.Context(), request)
	if err != nil {
		log.Println(err.Error())
		return models.Response(constants.StatusCodeBadRequest, data, err.Error()).SendResponse(ctx, http.StatusBadRequest)
	}
	return models.ResponseSuccess(constants.StatusCodeSuccess, constants.SuccessMessage, data).SendResponseSuccess(ctx, http.StatusOK)
}

func (h *Handler) ReadPatientAll(ctx *fiber.Ctx) error {
	data, err := h.apiRepository.ReadPatientAll(ctx.Context())
	if err != nil {
		return models.Response(constants.StatusCodeSystemError, nil, constants.StatusCodeSystemErrorMessage).SendResponse(ctx, http.StatusInternalServerError)
	}
	return models.ResponseSuccess(constants.StatusCodeSuccess, constants.SuccessMessage, data).SendResponseSuccess(ctx, http.StatusOK)
}
