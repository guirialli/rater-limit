// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/guirialli/rater_limit/internals/infra/webserver/controller"
	"github.com/guirialli/rater_limit/internals/usecases"
)

// Injectors from wire.go:

func NewAuthorController(db *sql.DB) *controller.Author {
	author := usecases.NewAuthor()
	book := usecases.NewBook()
	controllerAuthor := controller.NewAuthor(db, author, book)
	return controllerAuthor
}

func NewBookController(db *sql.DB) *controller.Book {
	book := usecases.NewBook()
	author := usecases.NewAuthor()
	controllerBook := controller.NewBook(db, book, author)
	return controllerBook
}

func NewAuthController(db *sql.DB, user usecases.IUser) *controller.Auth {
	auth := controller.NewAuth(db, user)
	return auth
}

// wire.go:

var setAuthorUseCaseDependency = wire.NewSet(usecases.NewAuthor, wire.Bind(new(usecases.IAuthor), new(*usecases.Author)))

var setBookUseCaseDependency = wire.NewSet(usecases.NewBook, wire.Bind(new(usecases.IBook), new(*usecases.Book)))
