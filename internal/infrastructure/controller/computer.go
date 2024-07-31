package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"screencapturer/internal/domain/computer"
	"screencapturer/internal/infrastructure/repository"
	"screencapturer/pkg/common"
)

type computerController struct {
	ComputerService computer.ComputerService
}

func NewComputerController(db *gorm.DB) *computerController {
	computerRepository := repository.NewComputerRepository(db)
	computerService := computer.NewComputerService(computerRepository)
	return &computerController{ComputerService: *computerService}
}

func (u *computerController) FindAll(c *gin.Context) {
	computers, err := u.ComputerService.FindAll()
	if err != nil {
		common.RaiseHttpError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, computers)
}

func (u *computerController) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	computer, err := u.ComputerService.FindById(id)
	if err != nil {
		common.RaiseHttpError(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, computer)
}

func (u *computerController) Create(c *gin.Context) {
	var createComputerDto computer.CreateComputerDto
	if err := c.BindJSON(&createComputerDto); err != nil {
		common.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	computer, err := u.ComputerService.Create(&createComputerDto)
	if err != nil {
		common.RaiseHttpError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, computer)
}

func (u *computerController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	if _, err := u.ComputerService.FindById(id); err != nil {
		common.RaiseHttpError(c, http.StatusNotFound, err)
		return
	}

	var updateComputerDto computer.UpdateComputerDto
	if err := c.BindJSON(&updateComputerDto); err != nil {
		common.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	computer, err := u.ComputerService.Update(&updateComputerDto, id)
	if err != nil {
		common.RaiseHttpError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, computer)
}

func (u *computerController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	if err := u.ComputerService.Delete(id); err != nil {
		common.RaiseHttpError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
