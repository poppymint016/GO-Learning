package services

import (
	"GO-Project/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExperienceService interface {
	Create(ctx context.Context, experience *models.ExperienceDto) error
	Update(id primitive.ObjectID, experience *models.ExperienceDto) error
	FindById(id primitive.ObjectID) (*models.ExperienceDto, error)
	FindAll() ([]*models.ExperienceDto, error)
	Delete(id primitive.ObjectID) error
}

type experienceServiceImpl struct {
	expCollection *mongo.Collection
	ctx           context.Context
}

func NewExperienceService(client *mongo.Client) ExperienceService {
	collection := client.Database("TODOLIST").Collection("experience")
	return &experienceServiceImpl{
		expCollection: collection,
		ctx:           context.TODO(),
	}
}

func (e *experienceServiceImpl) Create(ctx context.Context, exp *models.ExperienceDto) error {
	var experience = models.Experience{
		ID:         primitive.NewObjectID(),
		Experience: exp.Experience,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Time{},
	}
	_, err := e.expCollection.InsertOne(ctx, experience)
	return err
}

func (e *experienceServiceImpl) Update(id primitive.ObjectID, exp *models.ExperienceDto) error {
	var experience models.ExperienceDto

	err := e.expCollection.FindOne(e.ctx, bson.M{"_id": id}).Decode(&experience)
	if err != nil {
		return err
	}

	exp.UpdatedAt = time.Now()

	update := bson.M{
		"experience": exp.Experience,
		"updatedAt":  exp.UpdatedAt,
	}
	fmt.Println(update)

	_, err = e.expCollection.UpdateOne(e.ctx, bson.M{"_id": id}, bson.M{"$set": update})
	return err
}

func (e *experienceServiceImpl) FindById(id primitive.ObjectID) (*models.ExperienceDto, error) {
	var experience models.ExperienceDto

	err := e.expCollection.FindOne(e.ctx, bson.M{"_id": id}).Decode(&experience)
	if err != nil {
		return nil, err
	}

	return &experience, nil
}

func (e *experienceServiceImpl) FindAll() ([]*models.ExperienceDto, error) {
	cursor, err := e.expCollection.Find(e.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(e.ctx)

	var experiences []*models.ExperienceDto
	err = cursor.All(e.ctx, &experiences)
	if err != nil {
		return nil, err
	}

	return experiences, nil
}

func (e *experienceServiceImpl) Delete(id primitive.ObjectID) error {
	var experience models.ExperienceDto

	err := e.expCollection.FindOne(e.ctx, bson.M{"_id": id}).Decode(&experience)
	if err != nil {
		return err
	}

	_, err = e.expCollection.DeleteOne(e.ctx, bson.M{"_id": id})
	return err
}
