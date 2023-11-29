package services

import (
	"cobaecho/go-restapi/domain/helpers"
	"cobaecho/go-restapi/domain/item/models"
	"cobaecho/go-restapi/domain/item/repositories"
	"fmt"

	"gorm.io/gorm"
)

type itemServices struct {
	itemRepo repositories.ItemRepository
}

// Create implements ItemService.
func (service *itemServices) Create(item models.Item) helpers.Response {
	var response helpers.Response
	if err := service.itemRepo.Create(item); err != nil {
		response.Status = 500
		response.Messages = "Failed to create a new item"
	} else {
		response.Status = 200
		response.Messages = "Success to create a new item"
	}
	return response
}

// Delete implements ItemService.
func (service *itemServices) Delete(idItem int) helpers.Response {
	var response helpers.Response
	if err := service.itemRepo.Delete(idItem); err != nil {
		response.Status = 500
		response.Messages = fmt.Sprint("Failed to delete item", idItem)
	} else {
		response.Status = 200
		response.Messages = "Success to create a new item"
	}
	return response
}

// GetAll implements ItemService.
func (service *itemServices) GetAll() helpers.Response {
	var response helpers.Response
	data, err := service.itemRepo.GetAll()
	if err != nil {
		response.Status = 500
		response.Messages = "Failed to get items"
	} else {
		response.Status = 200
		response.Messages = "Success to get items"
		response.Data = data
	}
	return response
}

// GetById implements ItemService.
func (service *itemServices) GetById(idItem int) helpers.Response {
	var response helpers.Response
	data, err := service.itemRepo.GetById(idItem)
	if err != nil {
		response.Status = 500
		response.Messages = fmt.Sprint("Failed to get items", idItem)
	} else {
		response.Status = 200
		response.Messages = "Success to get items"
		response.Data = data
	}
	return response
}

// Update implements ItemService.
func (service *itemServices) Update(idItem int, item models.Item) helpers.Response {
	var response helpers.Response
	if err := service.itemRepo.Update(idItem, item); err != nil {
		response.Status = 500
		response.Messages = fmt.Sprint("Failed to update item", idItem)
	} else {
		response.Status = 200
		response.Messages = "Success update item"
	}
	return response
}

type ItemService interface {
	Create(item models.Item) helpers.Response
	Update(idItem int, item models.Item) helpers.Response
	Delete(idItem int) helpers.Response
	GetById(idItem int) helpers.Response
	GetAll() helpers.Response
}

func NewItemService(db *gorm.DB) ItemService {
	return &itemServices{itemRepo: repositories.NewItemRepository(db)}
}
