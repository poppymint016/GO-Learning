package handlers

import (
	"GO-Project/models"
	"GO-Project/responses"
	"GO-Project/services"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type expHandle struct {
	expService services.ExperienceService
}

func NewExperienceHandle(app fiber.Router, expService services.ExperienceService) {
	c := expHandle{
		expService: expService,
	}
	app.Post("/api/experiences", c.Create)
	app.Put("/api/:experienceId", c.Update)
	app.Get("/api/:experienceId", c.FindById)
	app.Get("/api/", c.FindAll)
	app.Delete("/api/:experienceId", c.Delete)
}

func (e expHandle) Create(c *fiber.Ctx) error {
	var exp *models.ExperienceDto

	if err := c.BodyParser(&exp); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.MessageResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	if err := e.expService.Create(c.Context(), exp); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MessageResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	return c.Status(http.StatusCreated).JSON(responses.MessageResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data:    &fiber.Map{"data": exp},
	})
}

func (e expHandle) Update(c *fiber.Ctx) error {
	experienceId := c.Params("experienceId")

	var exp *models.ExperienceDto

	objId, err := primitive.ObjectIDFromHex(experienceId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.MessageResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	if err := c.BodyParser(&exp); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.MessageResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	if err := e.expService.Update(objId, exp); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MessageResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	return c.Status(http.StatusOK).JSON(responses.MessageResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": exp},
	})
}

func (e expHandle) FindById(c *fiber.Ctx) error {
	experienceId := c.Params("experienceId")

	var exp *models.ExperienceDto

	objID, err := primitive.ObjectIDFromHex(experienceId)
	if err != nil {
		fmt.Println("0")
		slog.Error("error log id", slog.String("id", experienceId), slog.Any("err", err))
		return c.Status(http.StatusBadRequest).JSON(responses.MessageResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": "Not found"},
		})
	}

	if objID.IsZero() {
		slog.Error("error log id", slog.String("id", experienceId), slog.Any("err", err))
		return c.Status(http.StatusBadRequest).JSON(responses.MessageResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": "Not found"},
		})
	}

	exp, err = e.expService.FindById(objID)
	if err != nil {
		fmt.Println("1")
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(responses.MessageResponse{
				Status:  http.StatusNotFound,
				Message: "error",
				Data:    &fiber.Map{"data": "Experience not found"},
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(responses.MessageResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	return c.Status(http.StatusOK).JSON(responses.MessageResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": exp},
	})
}

func (e expHandle) FindAll(c *fiber.Ctx) error {
	exp, err := e.expService.FindAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MessageResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	return c.Status(http.StatusOK).JSON(responses.MessageResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": exp},
	})
}

func (e expHandle) Delete(c *fiber.Ctx) error {
	experienceId := c.Params("experienceId")

	objId, err := primitive.ObjectIDFromHex(experienceId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.MessageResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	if err := e.expService.Delete(objId); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MessageResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	return c.Status(http.StatusOK).JSON(responses.MessageResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": "Experience successfully deleted!"},
	})
}
