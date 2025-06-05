package main

import (
	"Kuraz-Tech-Challenge/config"
	"Kuraz-Tech-Challenge/handlers"
	database "Kuraz-Tech-Challenge/infrastructure"
	"Kuraz-Tech-Challenge/repository"
	"Kuraz-Tech-Challenge/router"
	"Kuraz-Tech-Challenge/usecase"

	"github.com/gin-gonic/gin"
)
func main() {
	cfg := config.LoadConfig()
	database.ConnectDB(cfg)
	database.MigrateSchema()

	//setup task api
	taskrepo := repository.NewTaskRepository(database.DB)
	taskusecase := usecase.NewTaskUsecase(taskrepo)
	taskhandler := handlers.NewTaskHanlder(taskusecase)

	route := gin.Default()
	api := route.Group("/api")

	router.SetUpRouter(api, taskhandler)
	route.Run()
}