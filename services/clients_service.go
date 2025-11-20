package service

import (
	"loyalty-api/internal/models"
	"loyalty-api/repository"
	"time"
)

type ClientsService interface {
	Save(clients models.Clients) error
	FindAll() []models.Clients
	GetClientById(id int) models.Clients
	DeleteClientsById(id int) error
	UpdateClient(id int, clients models.Clients) models.Clients
	SearchClient(query string) (models.Clients, error)
}
type clientsService struct {
	service repository.ClientsRepository
}

// SearchClient implements ClientsService.
func (c *clientsService) SearchClient(query string) (models.Clients, error) {
	var client models.Clients
	client, err := c.service.SearchClient(query)
	if err != nil {
		return client, err
	}
	return client, nil
}

// DeleteClientsById implements ClientsService.
func (c *clientsService) DeleteClientsById(id int) error {
	if err := c.service.DeleteClientsById(id); err != nil {
		return err
	}
	return nil
}

// FindAll implements ClientsService.
func (c *clientsService) FindAll() []models.Clients {
	return c.service.FindAll()
}

// GetClientById implements ClientsService.
func (c *clientsService) GetClientById(id int) models.Clients {
	return c.service.GetClientById(id)
}

// Save implements ClientsService.
func (c *clientsService) Save(clients models.Clients) error {
	clients.CreatedAt = time.Now().UTC()
	return c.service.Save(clients)
}

// UpdateClient implements ClientsService.
func (c *clientsService) UpdateClient(id int, clients models.Clients) models.Clients {
	return c.service.UpdateClient(id, clients)
}

// Constructeur
func NewClientsService(repo repository.ClientsRepository) ClientsService {
	return &clientsService{
		service: repo,
	}
}
