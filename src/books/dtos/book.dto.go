package dtos
import "github.com/go-playground/validator/v10"
// CreateDTO represents the structure for creating a new book.
type CreateDTO struct {
	Title  string `json:"title" bson:"title" validate:"required,min=3,max=100"`
	Author string `json:"author" bson:"author" validate:"required,min=3,max=100"`
	Year   string `json:"year" bson:"year" validate:"required,len=4"` // Year should be exactly 4 digits.
}

// UpdateDTO represents the structure for updating an existing book.
type UpdateDTO struct {
	Title  string `json:"title,omitempty" bson:"title,omitempty" validate:"omitempty,min=3,max=100"`
	Author string `json:"author,omitempty" bson:"author,omitempty" validate:"omitempty,min=3,max=100"`
	Year   string `json:"year,omitempty" bson:"year,omitempty" validate:"omitempty,len=4"` // Year should be exactly 4 digits.
}

// Validate method to validate CreateDTO and UpdateDTO.
func (dto *CreateDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(dto)
}

func (dto *UpdateDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(dto)
}