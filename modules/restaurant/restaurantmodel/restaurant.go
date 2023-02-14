package restaurantmodel

import (
	"awesomeProject/common"
	"errors"
	"strings"
)

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"string" gorm:"column:string;"`
	Addr            string        `json:"address" gorm:"column:addr;"`
	Logo            *common.Image `json:"logo" gorm:"column:logo;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name *string       `json:"string" gorm:"column:string;"`
	Addr *string       `json:"address" gorm:"column:addr;"`
	Logo *common.Image `json:"logo" gorm:"column:logo;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	Name string        `json:"string" gorm:"column:string;"`
	Addr string        `json:"address" gorm:"column:addr;"`
	Logo *common.Image `json:"logo" gorm:"column:logo;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("restaurant name cannot be blank")
	}
	return nil
}
