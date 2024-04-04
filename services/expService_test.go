package services

import (
	"GO-Project/models"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestNewExperienceService(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Success", func(mt *mtest.T) {
		something := NewExperienceService(mt.Client)
		require.NotNil(t, something)
	})
}
func TestCreate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Failed, Created error", func(mt *mtest.T) {

		e := NewExperienceService(mt.Client)

		payload := &models.ExperienceDto{
			Experience: "Hello",
		}

		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		err := e.Create(context.Background(), payload)
		require.Error(t, err)
	})

	mt.Run("OK", func(mt *mtest.T) {

		e := NewExperienceService(mt.Client)

		payload := &models.ExperienceDto{
			Experience: "Hello11",
		}

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		err := e.Create(context.Background(), payload)
		require.NoError(t, err)
	})
}

func TestUpdate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mt.Run("Success", func(mt *mtest.T) {
		e := NewExperienceService(mt.Client)

		id := primitive.NewObjectID()
		payload := &models.ExperienceDto{
			Experience: "Hello00",
		}

		mt.AddMockResponses(
			mtest.CreateCursorResponse(1, "services.mock", mtest.FirstBatch, bson.D{
				{Key: "_id", Value: id},
				{Key: "experience", Value: "Hallo"},
			}),
			bson.D{
				{Key: "ok", Value: 1},
				{Key: "Value", Value: bson.D{
					{Key: "_id", Value: id},
					{Key: "experience", Value: payload.Experience},
					{Key: "updatedAt", Value: time.Now()},
				}},
			},
		)
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "Value", Value: bson.D{
				{Key: "_id", Value: id},
				{Key: "experience", Value: payload.Experience},
				{Key: "updatedAt", Value: time.Now()},
			}},
		})
		err := e.Update(id, payload)
		require.NoError(t, err)
	})

	mt.Run("Failed, Updated Id not found", func(mt *mtest.T) {
		e := NewExperienceService(mt.Client)

		testErr := mtest.CommandError{
			Message: mongo.ErrNoDocuments.Error(),
			Name:    mongo.ErrNoDocuments.Error(),
		}

		mt.AddMockResponses(mtest.CreateCommandErrorResponse(testErr))
		err := e.Update(primitive.NewObjectID(), &models.ExperienceDto{})
		require.Error(t, err)

	})
}

func TestFindById(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Success", func(mt *mtest.T) {
		e := NewExperienceService(mt.Client)
		id := primitive.NewObjectID()
		payload := &models.ExperienceDto{
			Experience: "Hello",
		}

		mt.AddMockResponses(
			mtest.CreateCursorResponse(1, "services.mock", mtest.FirstBatch, bson.D{
				{Key: "_id", Value: id},
				{Key: "experience", Value: payload.Experience},
			}))
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "Value", Value: bson.D{
				{Key: "experience", Value: payload.Experience},
			}},
		})

		payload, err := e.FindById(id)
		require.NoError(t, err, *payload)

	})

	mt.Run("Failed, FindById not found", func(mt *mtest.T) {
		e := NewExperienceService(mt.Client)

		testErr := mtest.CommandError{
			Message: mongo.ErrNoDocuments.Error(),
			Name:    mongo.ErrNoDocuments.Error(),
		}

		mt.AddMockResponses(mtest.CreateCommandErrorResponse(testErr))
		payload, nil := e.FindById(primitive.NewObjectID())
		require.Error(t, nil, payload)
	})

}

func TestFindAll(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Success: FindAll", func(mt *mtest.T) {
		e := NewExperienceService(mt.Client)

		expected := []*models.ExperienceDto{
			{Experience: "Hello"},
			{Experience: "Hello1"},
		}

		first := mtest.CreateCursorResponse(1, "services.mock", mtest.FirstBatch, bson.D{
			{Key: "experience", Value: expected[0].Experience},
		})

		second := mtest.CreateCursorResponse(0, "services.mock", mtest.NextBatch, bson.D{
			{Key: "experience", Value: expected[1].Experience},
		})

		killCursors := mtest.CreateCursorResponse(0, "services.mock", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)

		result, nil := e.FindAll()
		require.NotNil(t, result, nil)
		require.Equal(t, expected, result)
	})

	mt.Run("Failed, FindAll Not found", func(mt *mtest.T) {
		e := NewExperienceService(mt.Client)

		testErr := mtest.CommandError{
			Message: mongo.ErrNoDocuments.Error(),
			Name:    mongo.ErrNoDocuments.Error(),
		}

		mt.AddMockResponses(mtest.CreateCommandErrorResponse(testErr))
		payload, err := e.FindAll()
		require.Error(t, err, payload)
	})

	mt.Run("Find-Cursor-ALL Data Type Error", func(mt *mtest.T) {
		e := NewExperienceService(mt.Client)

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "services.mock", mtest.FirstBatch, bson.D{
			{Key: "experience", Value: "Hello"}},
		))

		_, err := e.FindAll()
		require.Error(t, err)
	})

}

func TestDelete(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mt.Run("Success", func(mt *mtest.T) {
		e := NewExperienceService(mt.Client)
		id := primitive.NewObjectID()

		mt.AddMockResponses(
			mtest.CreateCursorResponse(1, "services.mock", mtest.FirstBatch, bson.D{
				{Key: "_id", Value: id},
			}))
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
		})

		err := e.Delete(id)
		require.Error(t, err)

	})

	mt.Run("Failed", func(mt *mtest.T) {
		e := NewExperienceService(mt.Client)
		id := primitive.NewObjectID()

		testErr := mtest.CommandError{
			Message: mongo.ErrNoDocuments.Error(),
			Name:    mongo.ErrNoDocuments.Error(),
		}

		mt.AddMockResponses(mtest.CreateCommandErrorResponse(testErr))
		err := e.Delete(id)
		require.Error(t, err)
	})
}
