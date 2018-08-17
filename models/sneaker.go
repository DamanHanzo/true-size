package models

import (
	u "github.com/damanhanzo/true-size/utils"
	"github.com/jinzhu/gorm"
	"fmt"
)

type Sneaker struct {
	gorm.Model
	Brand string `json: "brand"`
	Model string `json: "model"`
	TrueSize float32 `json: "trueSize"`
}

func(sneaker *Sneaker) Validate(map[string] interface{}, bool) {
	//...
}

func (sneaker *Sneaker) Create() (map[string] interface{}) {

	if resp, ok := sneaker.Validate(); !ok {
		return resp
	}

	GetDB().Create(sneaker)

	resp := u.Message(true, "success")
	resp["sneaker"] = contact
	return resp
}

func GetSneaker() (*Sneaker) {
	
	sneaker := &Sneaker{}
	err := GetDB().Table("sneakers").Where("id = ?", id).First(sneaker).Error
	if err != nil {
		return nil
	}
	return sneaker
}

func GetSneakers(sneaker unit) ([]*Sneaker) {

	sneakers := make([]*Sneaker, 0)
	err := GetDB().Table("sneakers").Find(&sneakers).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return sneakers

}