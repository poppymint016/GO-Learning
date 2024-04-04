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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
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
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Response-Success", func(mt *mtest.T) {
		NewExperienceHandle(f, mockStore)

		mockStore.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(nil)

		payload := &models.ExperienceDto{
			Experience: "Hello",
		}

		data, err := json.Marshal(payload)
		require.NoError(t, err)

		req := httptest.NewRequest(fiber.MethodPost, "/api/experiences", bytes.NewReader(data))
		req.Header.Set("Content-Type", "application/json")
		res, err := f.Test(req)
		require.NoError(t, err)
		defer func() { _ = res.Body.Close() }()
		require.Equal(t, fiber.StatusCreated, res.StatusCode)
	})

	mt.Run("Response-BadRequest", func(mt *mtest.T) {
		NewExperienceHandle(f, mockStore)

		req := httptest.NewRequest(fiber.MethodPost, "/api/experiences", nil)
		req.Header.Set("Content-Type", "application/json")
		res, err := f.Test(req)
		defer func() { _ = res.Body.Close() }()
		require.Nil(t, err)
		require.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})

	mt.Run("Response-internalserver-error", func(mt *mtest.T) {
		NewExperienceHandle(f, mockStore)

		mockStore.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(mongo.ErrUnacknowledgedWrite)

		payload := &models.ExperienceDto{
			Experience: "Hello",
		}

		data, err := json.Marshal(payload)
		require.NoError(t, err)

		req := httptest.NewRequest(fiber.MethodPost, "/api/experiences", bytes.NewReader(data))
		req.Header.Set("Content-Type", "application/json")
		res, err := f.Test(req)
		defer func() { _ = res.Body.Close() }()
		require.Nil(t, err)
		require.Equal(t, fiber.StatusInternalServerError, res.StatusCode)
	})

}

func TestUpdate(t *testing.T) {
	f := fiber.New()
	defer func() {
		_ = f.Shutdown()
	}()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_services.NewMockExperienceService(ctrl)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Response-Success", func(mt *mtest.T) {
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

		res, err := f.Test(req)
		require.NoError(t, err)
		defer func() { _ = res.Body.Close() }()
		require.Equal(t, fiber.StatusOK, res.StatusCode)

	})

	mt.Run("Response-NotFound-FormatID-Error", func(mt *mtest.T) {
		NewExperienceHandle(f, mockStore)

		req := httptest.NewRequest(fiber.MethodPut, "/api/yyyyyy", nil)
		res, err := f.Test(req)
		defer func() { _ = res.Body.Close() }()
		require.Nil(t, err)
		require.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})

	mt.Run("Response-BadRequest-BodyParser", func(mt *mtest.T) {
		NewExperienceHandle(f, mockStore)

		mockId := primitive.NewObjectID()
		url := "/api/" + mockId.Hex()

		req := httptest.NewRequest(fiber.MethodPut, url, nil)
		req.Header.Set("Content-Type", "application/json")

		res, err := f.Test(req)
		defer func() { _ = res.Body.Close() }()
		require.Nil(t, err)
		require.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})

	mt.Run("Response-internalserver-error", func(mt *mtest.T) {
		NewExperienceHandle(f, mockStore)

		payload := &models.ExperienceDto{
			Experience: "Hello",
		}

		data, err := json.Marshal(payload)
		require.NoError(t, err)

		mockId := primitive.NewObjectID()
		mockStore.EXPECT().Update(mockId, payload).Times(1).Return(mongo.ErrUnacknowledgedWrite)

		url := "/api/" + mockId.Hex()
		req := httptest.NewRequest(fiber.MethodPut, url, bytes.NewReader(data))
		req.Header.Set("Content-Type", "application/json")

		res, err := f.Test(req)
		defer func() { _ = res.Body.Close() }()
		require.Nil(t, err)
		require.Equal(t, fiber.StatusInternalServerError, res.StatusCode)
	})

}

func TestFindById(t *testing.T) {
	f := fiber.New()
	defer func() {
		_ = f.Shutdown()
	}()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_services.NewMockExperienceService(ctrl)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Response-Success", func(mt *mtest.T) {
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

		res, err := f.Test(req)
		require.NoError(t, err)
		defer func() { _ = res.Body.Close() }()

		require.Equal(t, fiber.StatusOK, res.StatusCode)
	})

	mt.Run("Response-NotFound-FormatID-Error", func(mt *mtest.T) {
		NewExperienceHandle(f, mockStore)

		req := httptest.NewRequest(fiber.MethodGet, "/api/yyyyyy", nil)
		res, err := f.Test(req)
		defer func() { _ = res.Body.Close() }()
		require.Nil(t, err)
		require.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})

	mt.Run("Response-NotFound-FormatID-Error", func(mt *mtest.T) {
		NewExperienceHandle(f, mockStore)

		req := httptest.NewRequest(fiber.MethodGet, "/api/000000000000000000000000", nil)
		res, err := f.Test(req)
		defer func() { _ = res.Body.Close() }()
		require.Nil(t, err)
		require.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})

	mt.Run("Response-Documents Not found", func(mt *mtest.T) {
		NewExperienceHandle(f, mockStore)

		payload := &models.ExperienceDto{
			Experience: "Hello",
		}

		data, err := json.Marshal(payload)
		require.NoError(t, err)

		mockId := primitive.NewObjectID()
		mockStore.EXPECT().FindById(mockId).Times(1).Return(err, mongo.ErrNoDocuments)

		url := "/api/" + mockId.Hex()
		req := httptest.NewRequest(fiber.MethodGet, url, bytes.NewReader(data))
		req.Header.Set("Content-Type", "application/json")

		res, err := f.Test(req)
		defer func() { _ = res.Body.Close() }()
		require.Nil(t, err)
		require.Equal(t, fiber.StatusNotFound, res.StatusCode)

	})

	mt.Run("Response-internalserver-error", func(mt *mtest.T) {
		NewExperienceHandle(f, mockStore)

		payload := &models.ExperienceDto{
			Experience: "Hello",
		}

		data, err := json.Marshal(payload)
		require.NoError(t, err)

		mockId := primitive.NewObjectID()
		mockStore.EXPECT().FindById(mockId).Times(1).Return(err, mongo.ErrUnacknowledgedWrite)

		url := "/api/" + mockId.Hex()
		req := httptest.NewRequest(fiber.MethodGet, url, bytes.NewReader(data))
		req.Header.Set("Content-Type", "application/json")

		res, err := f.Test(req)
		defer func() { _ = res.Body.Close() }()
		require.Nil(t, err)
		require.Equal(t, fiber.StatusInternalServerError, res.StatusCode)
	})

}

func TestFindAll(t *testing.T) {
	f := fiber.New()
	defer func() {
		_ = f.Shutdown()
	}()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_services.NewMockExperienceService(ctrl)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Response-Success", func(mt *mtest.T) {
		NewExperienceHandle(f, mockStore)

		payload := []*models.ExperienceDto{
			{Experience: "Hello"},
		}

		data, err := json.Marshal(payload)
		require.NoError(t, err)

		mockStore.EXPECT().FindAll().Times(1).Return(payload, err)

		req := httptest.NewRequest(fiber.MethodGet, "/api/", bytes.NewReader(data))
		req.Header.Set("Content-Type", "application/json")

		res, err := f.Test(req)
		require.NoError(t, err)
		defer func() { _ = res.Body.Close() }()

		require.Equal(t, fiber.StatusOK, res.StatusCode)
	})

	mt.Run("Response-internalserver-error", func(mt *mtest.T) {
		NewExperienceHandle(f, mockStore)

		payload := &models.ExperienceDto{
			Experience: "Hello",
		}

		data, err := json.Marshal(payload)
		require.NoError(t, err)

		mockStore.EXPECT().FindAll().Times(1).Return(err, mongo.ErrUnacknowledgedWrite)

		req := httptest.NewRequest(fiber.MethodGet, "/api/", bytes.NewReader(data))
		req.Header.Set("Content-Type", "application/json")

		res, err := f.Test(req)
		defer func() { _ = res.Body.Close() }()
		require.Nil(t, err)
		require.Equal(t, fiber.StatusInternalServerError, res.StatusCode)
	})

}

func TestDelete(t *testing.T) {
	f := fiber.New()
	defer func() {
		_ = f.Shutdown()
	}()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_services.NewMockExperienceService(ctrl)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Response-Success", func(mt *mtest.T) {
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

		res, err := f.Test(req)
		require.NoError(t, err)
		defer func() { _ = res.Body.Close() }()

		require.Equal(t, fiber.StatusOK, res.StatusCode)
	})

	mt.Run("Response-NotFound-FormatID-Error", func(mt *mtest.T) {
		NewExperienceHandle(f, mockStore)

		req := httptest.NewRequest(fiber.MethodDelete, "/api/yyyyyy", nil)
		res, err := f.Test(req)
		defer func() { _ = res.Body.Close() }()
		require.Nil(t, err)
		require.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})

	mt.Run("Response-internalserver-error", func(mt *mtest.T) {
		NewExperienceHandle(f, mockStore)

		payload := &models.ExperienceDto{
			Experience: "Hello",
		}

		data, err := json.Marshal(payload)
		require.NoError(t, err)

		mockId := primitive.NewObjectID()
		mockStore.EXPECT().Delete(mockId).Times(1).Return(mongo.ErrUnacknowledgedWrite)

		url := "/api/" + mockId.Hex()
		req := httptest.NewRequest(fiber.MethodDelete, url, bytes.NewReader(data))
		req.Header.Set("Content-Type", "application/json")

		res, err := f.Test(req)
		defer func() { _ = res.Body.Close() }()
		require.Nil(t, err)
		require.Equal(t, fiber.StatusInternalServerError, res.StatusCode)
	})

}
