package model

type Book struct {
	Name string `json:"name", gorm:"size:150;unique;not_null"`

	Common
}
