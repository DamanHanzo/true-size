package models

import (
	u "github.com/damanhanzo/true-size/utils"
	"github.com/jinzhu/gorm"
	"fmt"
)

type Sneaker struct {
	gorm.Model
	Brand string `json: "brand"`
	Sneaker_Model string `json: "sneaker_model"`
	True_Size float32 `json: "true_size"`
}

func(sneaker *Sneaker) Validate(map[string] interface{}, bool) {
	//...
}

func (sneaker *Sneaker) Create() (map[string] interface{}) {

	// if resp, ok := sneaker.Validate(); !ok {
	// 	return resp
	// }

	GetDB().Create(sneaker)

	resp := u.Message(true, "success")
	resp["sneaker"] = sneaker
	return resp
}

func GetSneaker(id uint) (*Sneaker) {

	sneaker := &Sneaker{}
	err := GetDB().Table("sneakers").Where("id = ?", id).First(sneaker).Error
	if err != nil {
		return nil
	}
	return sneaker
}

func GetSneakers() ([]*Sneaker) {

	sneakers := make([]*Sneaker, 0)
	err := GetDB().Table("sneakers").Find(&sneakers).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return sneakers
}

func GetTrueSize(brand string, model string) (map[string] interface{}) {
	sneakers := make([]*Sneaker, 0)
	err := GetDB().Table("sneakers").Where("brand = ?", brand).Where("sneaker_model = ?", model).Find(&sneakers).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//iterate over all sneaker and calculate an average
	var totalOfSizes float32
	for _, val := range sneakers {
		if val.True_Size != 0.0 {
			totalOfSizes += val.True_Size
		}
	}
	resp := u.Message(true, "success")
	resp["true_size"] = totalOfSizes
	return resp
}
