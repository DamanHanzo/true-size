
package models

import (
	u "github.com/damanhanzo/true-size/utils"
	"github.com/jinzhu/gorm"
	"fmt"
	"strings"
)

type Sneaker struct {
	gorm.Model
	Brand string `json: "brand"`
	Sneaker_Model string `json: "sneaker_model"`
	True_Size float32 `json: "true_size"`
}

func (sneaker *Sneaker) Validate() (map[string] interface{}, bool) {
	if sneaker.True_Size < 1.0 && sneaker.True_Size > 5.0 {
		return u.Message(false, "true_size has to be between 1 and 5"), false
	}
	if sneaker.Brand == "" {
		return u.Message(false, "Brand cannot be null"), false
	}
	if sneaker.Sneaker_Model == "" {
		return u.Message(false, "Sneaker Model is cannot be null"), false
	}
	return u.Message(true, "success"), true
}

func (sneaker *Sneaker) Create() (map[string] interface{}) {

	if resp, ok := sneaker.Validate(); !ok {
		return resp
	}
	sneaker.Brand = strings.ToLower(sneaker.Brand)
	sneaker.Sneaker_Model = strings.ToLower(sneaker.Sneaker_Model)
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
	err := GetDB().Table("sneakers").Where("brand = ?", strings.ToLower(brand)).Where("sneaker_model = ?", strings.ToLower(model)).Find(&sneakers).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//iterate over all sneakers and calculate an average
	var totalOfSizes float32
	for _, val := range sneakers {
		totalOfSizes += val.True_Size
	}
	trueSize := totalOfSizes/float32(len(sneakers))
	resp := u.Message(true, "success")
	resp["true_size"] = trueSize
	return resp
}
