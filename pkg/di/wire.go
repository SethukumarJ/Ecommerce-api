	//go:build wireinject
	// +build wireinject

	package di

	import (
		"github.com/google/wire"
		http "ecommerce/pkg/api"
		handler "ecommerce/pkg/api/handler"
		config "ecommerce/pkg/config"
		db "ecommerce/pkg/db"
		repository "ecommerce/pkg/repository"
		usecase "ecommerce/pkg/usecase"
		middleware "ecommerce/pkg/api/middleware"
	)

	func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
		wire.Build(
			db.ConnectDatabase,
			config.NewMailConfig,
		    repository.NewUserMongoRepository,
			repository.NewAdminMongoRepository, 
			handler.NewAdminHandler,
			handler.NewUserHandler,
			handler.NewAuthHandler,
			usecase.NewJWTUsecase,
			usecase.NewUserUseCase,
			usecase.NewAdminUseCase,
			usecase.NewAuthUsecase,
			middleware.NewMiddlewareUser,
			http.NewServerHTTP,
		)
		
		return &http.ServerHTTP{}, nil
	}
