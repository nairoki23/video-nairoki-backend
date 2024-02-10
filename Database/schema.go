package Database

import (
	"gorm.io/gorm"
)

type Hello struct {
	gorm.Model
	Body string `json:"body" gorm:"size:255"`
}
type Konnitiha struct {
	gorm.Model
	Body string `json:"body" gorm:"size:255"`
}

// DBはじめ
func SetSchema(db *gorm.DB) {
	// データベースにテーブルを作成
	db.AutoMigrate(&Hello{}, &Konnitiha{}) //本番のときはこれ消そう
}
