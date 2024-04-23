package dto

import (
	"errors"

	sf "github.com/wissance/stringFormatter"
)

type StackData struct {
	Id   uint `json:"id"`
	Data int  `json:"data"`
}

func (data *StackData) CheckRequiredFieldNotEmpty() error {
	if data.Data == 0 {
		return ErrEmptyField("StackData", "Data")
	}
	return nil
}

var ErrEmptyField = func(structName, field string) error {
	return errors.New(sf.Format("{0} contains empty field {1}, which is not allowed.", structName, field))
}

func NewData(dto StackData) StackData {
	return StackData{
		Id:   dto.Id,
		Data: dto.Data,
	}
}
