package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Hello struct {
	ID   uint   `gorm:"primary_key"`
	Body string `json:"body" gorm:"size:255"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数から接続情報を取得
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", dbHost, dbUser, dbPassword, dbName, dbPort)

	// GORMでデータベースに接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// データベースにテーブルを作成
	db.AutoMigrate(&Hello{})

	//以下API
	e := echo.New()
	e.GET("/",
		func(c echo.Context) error {
			u := &Hello{
				Body: "Hello World",
				ID:   213,
			}
			return c.JSON(404, u)
		})
	e.GET("/hello", func(c echo.Context) error {
		var hellos []Hello
		db.Find(&hellos)
		return c.JSON(200, hellos)
	})

	e.GET("/set-hello", func(c echo.Context) error {
		text := c.QueryParam("hello")
		hello := &Hello{
			Body: text,
		}
		db.Create(&hello)
		return c.JSON(200, hello)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
