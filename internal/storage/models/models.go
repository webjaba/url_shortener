package models

type Url struct {
	Alias string `gorm:"size:10;primaryKey;not null"`
	Url   string `gorm:"not null"`
}
