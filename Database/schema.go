package Database

import (
	"gorm.io/gorm"
	"time"
)

type baseModel struct {
	// UUID が主キーになり、UUID は Postgres が勝手に生成する
	ID        string `gorm:"primaryKey;size:255;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Hello struct {
	baseModel
	Body string `json:"body" gorm:"size:255"`
}

type Video struct {
	baseModel
	Title       string
	Description string
	ViewCount   uint
	UserID      string `gorm:"primaryKey;size:255;default:uuid_generate_v4()"`
}

type User struct {
	baseModel
	Name  string
	Video []Video
}

// DBはじめ
func SetSchema(db *gorm.DB) {
	// データベースにテーブルを作成
	db.AutoMigrate(&Hello{}) //本番のときはこれ消そう
}
