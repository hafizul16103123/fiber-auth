package bookService

import (
	"context"

	"fiber-app/src/books/dtos"
	"fiber-app/src/books/repository"
	"fiber-app/src/common"
	"fiber-app/src/models"
	"fiber-app/src/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type BookService struct {
	repo repository.BookRepository
}

// NewBookService initializes the repository and returns a new BookService instance.
func NewBookService() *BookService {
	// Initialize the database collection
	dbCollection := common.GetDBCollection("books")

	// Initialize the repository with the collection
	repo := repository.NewBookRepository(dbCollection)

	// Return the service with the repository
	return &BookService{repo: repo}
}

func (s *BookService) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	return s.repo.GetAllBooks(ctx)
}

func (s *BookService) GetBookByID(ctx context.Context, id string) (*models.Book, error) {
	return s.repo.GetBookByID(ctx, id)
}

func (s *BookService) CreateBook(ctx context.Context, dto *dtos.CreateDTO) (*models.Book, error) {
	// Validate the DTO
	if err := dto.Validate(); err != nil {
		// Format validation error
		errMsg := utils.FormatValidationError(err)
		return nil, &utils.Error{Msg: errMsg} // Return a custom error
	}
	book := models.Book{
		Title:  dto.Title,
		Author: dto.Author,
		Year:   dto.Year,
	}
	_, err := s.repo.CreateBook(ctx, &book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *BookService) UpdateBook(ctx context.Context, id string, dto *dtos.UpdateDTO) (*models.Book, error) {
	updateData := map[string]interface{}{}
	if dto.Title != "" {
		updateData["title"] = dto.Title
	}
	if dto.Author != "" {
		updateData["author"] = dto.Author
	}
	if dto.Year != "" {
		updateData["year"] = dto.Year
	}

	_, err := s.repo.UpdateBook(ctx, id, updateData)
	if err != nil {
		return nil, err
	}

	return s.repo.GetBookByID(ctx, id)
}

func (s *BookService) DeleteBook(ctx context.Context, id string) (*mongo.DeleteResult , error) {
	res, err := s.repo.DeleteBook(ctx, id)
	return res,err
}







// package bookService

// import (
// 	"context"

// 	"fiber-app/src/books/dtos"
// 	"fiber-app/src/common"
// 	"fiber-app/src/models"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type BookService struct {
// 	dbCollection *mongo.Collection
// }

// func NewBookService() *BookService {
// 	return &BookService{
// 		dbCollection: common.GetDBCollection("books"),
// 	}
// }

// func (s *BookService) GetAllBooks(ctx context.Context) ([]models.Book, error) {
// 	books := make([]models.Book, 0)
// 	cursor, err := s.dbCollection.Find(ctx, bson.M{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Iterate over the cursor
// 	for cursor.Next(ctx) {
// 		book := models.Book{}
// 		err := cursor.Decode(&book)
// 		if err != nil {
// 			return nil, err
// 		}
// 		books = append(books, book)
// 	}

// 	return books, nil
// }

// func (s *BookService) GetBookByID(ctx context.Context, id string) (*models.Book, error) {
// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	book := models.Book{}
// 	err = s.dbCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&book)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &book, nil
// }

// func (s *BookService) CreateBook(ctx context.Context, dto *dtos.CreateDTO) (*mongo.InsertOneResult, error) {
// 	book := models.Book{
// 		Title:  dto.Title,
// 		Author: dto.Author,
// 		Year:   dto.Year,
// 	}
// 	result, err := s.dbCollection.InsertOne(ctx, book)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func (s *BookService) UpdateBook(ctx context.Context, id string, dto *dtos.UpdateDTO) (*mongo.UpdateResult, error) {
// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	updateData := bson.M{}
// 	if dto.Title != "" {
// 		updateData["title"] = dto.Title
// 	}
// 	if dto.Author != "" {
// 		updateData["author"] = dto.Author
// 	}
// 	if dto.Year != "" {
// 		updateData["year"] = dto.Year
// 	}

// 	result, err := s.dbCollection.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": updateData})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

// func (s *BookService) DeleteBook(ctx context.Context, id string) (*mongo.DeleteResult, error) {
// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	result, err := s.dbCollection.DeleteOne(ctx, bson.M{"_id": objectId})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }