package computer

import (
	"gorm.io/gorm"

	"screencapturer/internal/domain/model"
)

type CreateComputerDto struct {
	Name string `json:"name"`
}

type UpdateComputerDto struct {
	Name string `json:"name"`
}

func MapUpdateComputerDto(dto *UpdateComputerDto, id int) model.Computer {
	return model.Computer{
		Model: gorm.Model{ID: uint(id)},
		Name:  dto.Name,
	}
}
