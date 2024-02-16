package handlers

import (
	"GO-Project/models"
	mock_services "GO-Project/services/mocks"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestNewExperienceHandle(t *testing.T) {
	fiber := fiber.New()
	defer func() {
		_ = fiber.Shutdown()
	}()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_services.NewMockExperienceService(ctrl)

	NewExperienceHandle(fiber, mockStore)
}

func TestCreate(t *testing.T) {
	f := fiber.New()
	defer func() {
		_ = f.Shutdown()
	}()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_services.NewMockExperienceService(ctrl)

	NewExperienceHandle(f, mockStore)

	mockStore.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(nil)

	payload := &models.ExperienceDto{
		Experience: "Hello",
	}

	data, err := json.Marshal(payload)
	require.NoError(t, err)

	req := httptest.NewRequest(fiber.MethodPost, "/api/experiences", bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")

	response, err := f.Test(req)
	require.NoError(t, err)
	defer func() { _ = response.Body.Close() }()

	require.Equal(t, 201, response.StatusCode)
}

func TestUpdate(t *testing.T) {
	f := fiber.New()
	defer func() {
		_ = f.Shutdown()
	}()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_services.NewMockExperienceService(ctrl)
	NewExperienceHandle(f, mockStore)

	payload := &models.ExperienceDto{
		Experience: "Helloooooo",
	}

	data, err := json.Marshal(payload)
	require.NoError(t, err)

	mockId := primitive.NewObjectID()
	mockStore.EXPECT().Update(mockId, payload).Times(1).Return(nil)

	url := "/api/" + mockId.Hex()
	req := httptest.NewRequest(fiber.MethodPut, url, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")

	response, err := f.Test(req)
	require.NoError(t, err)
	defer func() { _ = response.Body.Close() }()

	require.Equal(t, 200, response.StatusCode)

}

func TestFindById(t *testing.T) {
	f := fiber.New()
	defer func() {
		_ = f.Shutdown()
	}()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_services.NewMockExperienceService(ctrl)
	NewExperienceHandle(f, mockStore)

	payload := &models.ExperienceDto{
		Experience: "Hellooo",
	}

	data, err := json.Marshal(payload)
	require.NoError(t, err)

	mockId := primitive.NewObjectID()
	mockStore.EXPECT().FindById(mockId).Times(1).Return(payload, err)

	url := "/api/" + mockId.Hex()
	req := httptest.NewRequest(fiber.MethodGet, url, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")

	response, err := f.Test(req)
	require.NoError(t, err)
	defer func() { _ = response.Body.Close() }()

	require.Equal(t, 200, response.StatusCode)
}

func TestFindAll(t *testing.T) {
	f := fiber.New()
	defer func() {
		_ = f.Shutdown()
	}()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_services.NewMockExperienceService(ctrl)
	NewExperienceHandle(f, mockStore)

	payload := []*models.ExperienceDto{
		{Experience: "Hello"},
	}

	data, err := json.Marshal(payload)
	require.NoError(t, err)

	mockStore.EXPECT().FindAll().Times(1).Return(payload, err)

	req := httptest.NewRequest(fiber.MethodGet, "/api/", bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")

	response, err := f.Test(req)
	require.NoError(t, err)
	defer func() { _ = response.Body.Close() }()

	require.Equal(t, 200, response.StatusCode)

}

func TestDelete(t *testing.T) {
	f := fiber.New()
	defer func() {
		_ = f.Shutdown()
	}()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_services.NewMockExperienceService(ctrl)
	NewExperienceHandle(f, mockStore)

	payload := &models.ExperienceDto{
		Experience: "Hellooo",
	}

	data, err := json.Marshal(payload)
	require.NoError(t, err)

	mockId := primitive.NewObjectID()
	mockStore.EXPECT().Delete(mockId).Times(1).Return(err)

	url := "/api/" + mockId.Hex()
	req := httptest.NewRequest(fiber.MethodDelete, url, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")

	response, err := f.Test(req)
	require.NoError(t, err)
	defer func() { _ = response.Body.Close() }()

	require.Equal(t, 200, response.StatusCode)
}
