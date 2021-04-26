package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Item struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	CartID    uint32 `json:"cart_id"`
	Product   string
	Quantity  int
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (i *Item) Prepare() {
	i.ID = 0
	i.CartID = 0
	i.Product = html.EscapeString(strings.TrimSpace(i.Product))
	i.Quantity = 0
	i.CreatedAt = time.Now()
	i.UpdatedAt = time.Now()
}

func (i *Item) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	switch strings.ToLower(action) {
	default:
		if i.Product == "" {
			err = errors.New("Required Product")
			errorMessages["Required_product"] = err.Error()
		}
		if i.Quantity < 0 {
			err = errors.New("Quantity must be positive")
			errorMessages["Positive_quantity"] = err.Error()
		}
	}
	return errorMessages
}

func (i *Item) SaveItem(db *gorm.DB) (*Item, error) {
	err := db.Debug().Create(&i).Error
	if err != nil {
		return &Item{}, err
	}
	if i.ID != 0 {
		err = db.Debug().Model(&Cart{}).Where("id = ?", i.CartID).Error
		if err != nil {
			return &Item{}, err
		}
	}
	return i, nil
}
