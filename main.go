package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/guntoroyk/cake-store-api/config"
	httpHandler "github.com/guntoroyk/cake-store-api/handler/http"
	"github.com/guntoroyk/cake-store-api/lib/validator"
	dbRepo "github.com/guntoroyk/cake-store-api/repository/db"
	"github.com/guntoroyk/cake-store-api/storage"
	"github.com/guntoroyk/cake-store-api/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.GetConfig()
	e := echo.New()

	validator := validator.GetValidator()
	db := storage.NewDB(&config.DB)
	cakeRepo := dbRepo.NewCakeRepo(db)
	cakeUsecase := usecase.NewCakeUsecase(cakeRepo, validator)
	handler := httpHandler.NewHandler(cakeUsecase)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/cakes", handler.GetCakes)
	e.POST("/cakes", handler.CreateCake)
	e.GET("/cakes/:id", handler.GetCake)
	e.PATCH("/cakes/:id", handler.UpdateCake)
	e.DELETE("/cakes/:id", handler.DeleteCake)

	e.Logger.Fatal(e.Start(config.App.Host + ":" + config.App.Port))
}
