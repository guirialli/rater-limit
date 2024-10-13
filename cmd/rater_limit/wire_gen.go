// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/guirialli/rater_limit/config"
	"github.com/guirialli/rater_limit/internals/infra/database"
	"github.com/guirialli/rater_limit/internals/infra/webserver/controller"
	"github.com/guirialli/rater_limit/internals/infra/webserver/middleware"
	"github.com/guirialli/rater_limit/internals/infra/webserver/router"
	"github.com/guirialli/rater_limit/internals/usecases"
)

// Injectors from wire.go:

// DI
func NewAuthorController(db *sql.DB) *controller.Author {
	author := usecases.NewAuthor()
	book := usecases.NewBook()
	utils := controller.NewUtils()
	controllerAuthor := controller.NewAuthor(db, author, book, utils)
	return controllerAuthor
}

func NewBookController(db *sql.DB) *controller.Book {
	book := usecases.NewBook()
	author := usecases.NewAuthor()
	utils := controller.NewUtils()
	controllerBook := controller.NewBook(db, book, author, utils)
	return controllerBook
}

func NewAuthController(db *sql.DB) *controller.Auth {
	user := newUser()
	utils := controller.NewUtils()
	auth := controller.NewAuth(db, user, utils)
	return auth
}

func NewAuthorRouter(db *sql.DB) *router.Author {
	author := usecases.NewAuthor()
	book := usecases.NewBook()
	utils := controller.NewUtils()
	controllerAuthor := controller.NewAuthor(db, author, book, utils)
	user := newUser()
	routerAuthor := router.NewAuthor(controllerAuthor, user)
	return routerAuthor
}

func NewBookRouter(db *sql.DB) *router.Book {
	book := usecases.NewBook()
	author := usecases.NewAuthor()
	utils := controller.NewUtils()
	controllerBook := controller.NewBook(db, book, author, utils)
	user := newUser()
	routerBook := router.NewBook(controllerBook, user)
	return routerBook
}

func NewAuthRouter(db *sql.DB) *router.Auth {
	user := newUser()
	utils := controller.NewUtils()
	auth := controller.NewAuth(db, user, utils)
	routerAuth := router.NewAuth(auth)
	return routerAuth
}

func NewRaterLimitMiddleware() *middleware.RaterLimit {
	user := newUser()
	raterLimit := newRateLimitUseCase(user)
	middlewareRaterLimit := middleware.NewRaterLimit(raterLimit)
	return middlewareRaterLimit
}

// wire.go:

// Use Cases Dependency
var setAuthorUseCaseDependency = wire.NewSet(usecases.NewAuthor, wire.Bind(new(usecases.IAuthor), new(*usecases.Author)))

var setBookUseCaseDependency = wire.NewSet(usecases.NewBook, wire.Bind(new(usecases.IBook), new(*usecases.Book)))

var setUserUseCaseDependency = wire.NewSet(
	newUser, wire.Bind(new(usecases.IUser), new(*usecases.User)),
)

var setRaterLimitUseCaseDependency = wire.NewSet(
	newRateLimitUseCase,
	setUserUseCaseDependency, wire.Bind(new(usecases.IRaterLimit), new(*usecases.RaterLimit)),
)

var setHttpHandlerErrorDependency = wire.NewSet(controller.NewUtils, wire.Bind(new(controller.IHttpHandlerError), new(*controller.Utils)))

// controller dependency
var setAuthorControllerDependency = wire.NewSet(controller.NewAuthor, setBookUseCaseDependency,
	setAuthorUseCaseDependency,
	setHttpHandlerErrorDependency, wire.Bind(new(controller.IAuthor), new(*controller.Author)),
)

var setBookControllerDependency = wire.NewSet(controller.NewBook, setBookUseCaseDependency,
	setAuthorUseCaseDependency,
	setHttpHandlerErrorDependency, wire.Bind(new(controller.IBooks), new(*controller.Book)),
)

var setAuthControllerDependency = wire.NewSet(controller.NewAuth, setUserUseCaseDependency, wire.Bind(new(controller.IAuth), new(*controller.Auth)))

// router dependency
var setAuthMiddlewareDependency = wire.NewSet(
	newUser, wire.Bind(new(router.IAuthToken), new(*usecases.User)),
)

// utils constructors
// This function create a user use case without errors
func newUser() *usecases.User {
	jwtConfig := config.LoadJwtConfig()
	user, err := usecases.NewUser(jwtConfig.Secret, jwtConfig.ExpireIn, jwtConfig.UnitTime)
	if err != nil {
		panic(err)
	}
	return user
}

func newRateLimitUseCase(user usecases.IUser) *usecases.RaterLimit {
	cfg, _ := config.LoadRaterLimitConfig()
	rdb := database.NewRedisClient()
	raterLimit, err := usecases.NewRaterLimit(user, *cfg, rdb)
	if err != nil {
		panic(err)
	}
	return raterLimit
}
