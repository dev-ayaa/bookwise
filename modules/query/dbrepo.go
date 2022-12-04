package query

import (
	"github.com/yusuf/p-catalogue/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CatalogueRepo interface {
	SendAvailableBooks() ([]model.Library, error)
	CreateUserAccount(user model.User) (int, primitive.ObjectID, error)
	VerifyUser(email, password, hashedPassword string) (bool, error)
	UpdateUserDetails(userID primitive.ObjectID, token, renewToken string) error
	UpdateUserBook(userID primitive.ObjectID, book model.Library) error
	GetUserBooks(userID primitive.ObjectID) ([]model.UserLibrary, error)
	FindBook(userID primitive.ObjectID, title string) (model.UserLibrary, error)
	DeleteBook(bookID, userID primitive.ObjectID) error

	CheckLibrary(title string, bookData model.Book) (int64, primitive.ObjectID, error)
	GetBook(bookID primitive.ObjectID) (primitive.M, error)
}
