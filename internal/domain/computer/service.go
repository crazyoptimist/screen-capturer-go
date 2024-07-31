package computer

import (
	"screencapturer/internal/domain/model"
)

type ComputerRepository interface {
	FindAll() ([]*model.Computer, error)
	FindById(id int) (*model.Computer, error)
	Create(computer model.Computer) (*model.Computer, error)
	Update(computer model.Computer) (*model.Computer, error)
	Delete(computer model.Computer) error
}

type ComputerService struct {
	ComputerRepository ComputerRepository
}

func NewComputerService(computerRepository ComputerRepository) *ComputerService {
	return &ComputerService{ComputerRepository: computerRepository}
}

func (u *ComputerService) FindAll() ([]*model.Computer, error) {
	return u.ComputerRepository.FindAll()
}

func (u *ComputerService) FindById(id int) (*model.Computer, error) {
	return u.ComputerRepository.FindById(id)
}

func (u *ComputerService) Create(createComputerDto *CreateComputerDto) (*model.Computer, error) {
	return u.ComputerRepository.Create(model.Computer{Name: createComputerDto.Name, IsActive: false})
}

func (u *ComputerService) Update(updateComputerDto *UpdateComputerDto, id int) (*model.Computer, error) {
	return u.ComputerRepository.Update(MapUpdateComputerDto(updateComputerDto, id))
}

func (u *ComputerService) Delete(id int) error {
	return u.ComputerRepository.Delete(model.Computer{Common: model.Common{ID: id}})
}
