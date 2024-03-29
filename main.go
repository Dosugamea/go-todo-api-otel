package main

import (
	"os"

	_ "github.com/Dosugamea/go-todo-api-otel/docs"
	"github.com/Dosugamea/go-todo-api-otel/internal/bootstrap"
	"github.com/Dosugamea/go-todo-api-otel/internal/infrastructure/database"
	"github.com/Dosugamea/go-todo-api-otel/internal/infrastructure/persistence"
	"github.com/Dosugamea/go-todo-api-otel/internal/infrastructure/router"
	"github.com/Dosugamea/go-todo-api-otel/internal/interface/handler"
	"github.com/Dosugamea/go-todo-api-otel/internal/interface/routing"
	"github.com/Dosugamea/go-todo-api-otel/internal/model"
	"github.com/Dosugamea/go-todo-api-otel/internal/usecase"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Todo tasks API
// @version 1.0
// @description This is a simple todo tasks API
// @title Todo tasks API

// @host 127.0.0.1:8080
// @BasePath /api/v1

// @schemes http https
// @produce	application/json
// @consumes application/json
func main() {
	// OpentelemetryのTracerを初期化
	// APMにJaegerを使う
	jaegerEndpoint := os.Getenv("EXPORTER_ENDPOINT")
	envName := os.Getenv("DEPLOY_ENV_NAME")
	if envName == "" {
		envName = "development"
	}
	bootstrap.InitTracer(jaegerEndpoint, envName)

	// Echoを初期化
	r := router.New()
	r.GET("/swagger/*", echoSwagger.WrapHandler)

	// データベースを初期化
	db := database.GetConnection()
	if err := db.AutoMigrate(&model.Task{}); err != nil {
		panic(err)
	}

	// 依存関係を初期化
	taskPersistence := persistence.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskPersistence)

	// ハンドラ(echo インターフェイス)を初期化
	h := handler.NewTaskHandler(taskUsecase)

	// ルーティングを初期化
	v1 := r.Group("/api/v1")
	routing.RegisterTaskRoutings(v1, h)

	// サーバーを起動
	r.Logger.Fatal(r.Start(":8080"))
}
