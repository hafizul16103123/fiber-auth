package repository

import (
	"context"

	"fiber-app/src/common"
	"fiber-app/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository interface {
	GetAllBooks(ctx context.Context) ([]models.Book, error)
	GetBookByID(ctx context.Context, id string) (*models.Book, error)
	CreateBook(ctx context.Context, book *models.Book) (*mongo.InsertOneResult, error)
	UpdateBook(ctx context.Context, id string, updateData map[string]interface{}) (*mongo.UpdateResult, error)
	DeleteBook(ctx context.Context, id string) (*mongo.DeleteResult, error)
}

type bookRepository struct {
	commonRepo *common.CommonRepository
}

func NewBookRepository(collection *mongo.Collection) BookRepository {
	return &bookRepository{commonRepo: common.NewCommonRepository(collection)}
}

func (r *bookRepository) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	err := r.commonRepo.FindAll(ctx, bson.M{}, &books)
	return books, err
}

func (r *bookRepository) GetBookByID(ctx context.Context, id string) (*models.Book, error) {
	objectID, err := r.commonRepo.ConvertID(id)
	if err != nil {
		return nil, err
	}

	var book models.Book
	err = r.commonRepo.FindOne(ctx, bson.M{"_id": objectID}, &book)
	return &book, err
}

func (r *bookRepository) CreateBook(ctx context.Context, book *models.Book) (*mongo.InsertOneResult, error) {
	return r.commonRepo.InsertOne(ctx, book)
}

func (r *bookRepository) UpdateBook(ctx context.Context, id string, updateData map[string]interface{}) (*mongo.UpdateResult, error) {
	objectID, err := r.commonRepo.ConvertID(id)
	if err != nil {
		return nil, err
	}

	return r.commonRepo.UpdateOne(ctx, bson.M{"_id": objectID}, updateData)
}

func (r *bookRepository) DeleteBook(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	objectID, err := r.commonRepo.ConvertID(id)
	if err != nil {
		return nil, err
	}

	return r.commonRepo.DeleteOne(ctx, bson.M{"_id": objectID})
}
