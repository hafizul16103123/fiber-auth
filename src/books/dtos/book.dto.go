package dtos

// CreateDTO represents the structure for creating a new book.
type CreateDTO struct {
	Title  string `json:"title" bson:"title"`
	Author string `json:"author" bson:"author"`
	Year   string `json:"year" bson:"year"`
}

// UpdateDTO represents the structure for updating an existing book.
type UpdateDTO struct {
	Title  string `json:"title,omitempty" bson:"title,omitempty"`
	Author string `json:"author,omitempty" bson:"author,omitempty"`
	Year   string `json:"year,omitempty" bson:"year,omitempty"`
}
