package Database

import (
	"gorm.io/gorm"
	"time"
)

type Hello struct {
	gorm.Model
	Body string `json:"body" gorm:"size:255"`
}

type Video struct {
	ID        uint      // 主キーの標準フィールド
	Title     string    // 通常の文字列フィールド
	CreatedAt time.Time // GORMによって自動的に管理される作成時間
	UpdatedAt time.Time // GORMによって自動的に管理される更新時間
}

// DBはじめ
func SetSchema(db *gorm.DB) {
	// データベースにテーブルを作成
	db.AutoMigrate(&Hello{}) //本番のときはこれ消そう
}
