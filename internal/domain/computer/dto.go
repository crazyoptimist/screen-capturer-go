package computer

import (
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
		Common: model.Common{ID: id},
		Name:   dto.Name,
	}
}
