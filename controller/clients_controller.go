package controller

import (
	"loyalty-api/internal/models"
	service "loyalty-api/services"

	"github.com/gin-gonic/gin"
)

type ClientsController interface {
	Save(ctx *gin.Context) models.Clients
	FindAll() []models.Clients
	GetClientById(id int) models.Clients
	DeleteClientsById(id int) error
	UpdateClient(ctx *gin.Context, id int) models.Clients
	SearchClient(query string) (models.Clients, error)
}
type clientsController struct {
	controller service.ClientsService
}

// SearchClient implements ClientsController.
func (c *clientsController) SearchClient(query string) (models.Clients, error) {
	var client models.Clients
	client, err := c.controller.SearchClient(query)
	if err != nil {
		return client, err
	}
	return client, nil
}

// DeleteClientsById implements ClientsController.
func (c *clientsController) DeleteClientsById(id int) error {
	if err := c.controller.DeleteClientsById(id); err != nil {
		return err
	}
	return nil
}

// FindAll implements ClientsController.
func (c *clientsController) FindAll() []models.Clients {
	return c.controller.FindAll()
}

// GetClientById implements ClientsController.
func (c *clientsController) GetClientById(id int) models.Clients {
	return c.controller.GetClientById(id)
}

// Save implements ClientsController.
func (c *clientsController) Save(ctx *gin.Context) models.Clients {
	var clients models.Clients
	ctx.BindJSON(&clients)
	c.controller.Save(clients)
	return clients
}

// UpdateClient implements ClientsController.
func (c *clientsController) UpdateClient(ctx *gin.Context, id int) models.Clients {
	var clients models.Clients
	ctx.BindJSON(&clients)
	c.controller.UpdateClient(id, clients)
	return clients
}

// Constructeur
func NewClientsController(service service.ClientsService) ClientsController {
	return &clientsController{
		controller: service,
	}
}
