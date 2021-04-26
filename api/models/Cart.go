package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Cart struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Items     []Item    `gorm:"foreignKey:CartID"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Cart) Prepare() {
	c.Items = []Item{}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (c *Cart) SaveCart(db *gorm.DB) (*Cart, error) {
	var err error
	err = db.Debug().Model(&Cart{}).Create(&c).Error
	if err != nil {
		return &Cart{}, err
	}
	return c, nil
}

func (c *Cart) FindCartByID(db *gorm.DB, pid uint64) (*Cart, error) {
	var err error
	err = db.Debug().Model(&Cart{}).Where("id = ?", pid).Take(&c).Error
	if err != nil {
		return &Cart{}, err
	}
	return c, nil
}
