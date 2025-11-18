package repository

import (
	"fmt"
	"loyalty-api/internal/models"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ClientsRepository interface {
	Save(clients models.Clients) models.Clients
	FindAll() []models.Clients
	GetClientById(id int) models.Clients
	DeleteClientsById(id int) error
	UpdateClient(id int, clients models.Clients) models.Clients
	SearchClient(query string) (models.Clients, error)
}
type clientRepository struct {
	DB *gorm.DB
}

// SearchClient implements ClientsRepository.
func (c *clientRepository) SearchClient(query string) (models.Clients, error) {
	var client models.Clients
	resultat := c.DB.Where("email ILIKE ? OR mobile ILIKE ?", "%"+query+"%", "%"+query+"%").First(&client)

	if resultat.Error != nil {
		return client, resultat.Error
	}
	return client, nil
}

// DeleteClientsById implements ClientsRepository.
func (c *clientRepository) DeleteClientsById(id int) error {
	if result := c.DB.Delete(&models.Clients{}, id).Error; result != nil {
		return result
	}
	return nil
}

// FindAll implements ClientsRepository.
func (c *clientRepository) FindAll() []models.Clients {
	var clients []models.Clients
	c.DB.Find(&clients)
	return clients
}

// GetClientById implements ClientsRepository.
func (c *clientRepository) GetClientById(id int) models.Clients {
	var clients models.Clients
	c.DB.Find(&clients, id)
	return clients
}

// Save implements ClientsRepository.
func (c *clientRepository) Save(clients models.Clients) models.Clients {
	id := uuid.New()
	clients.UserRef = id.String()
	if result := c.DB.Save(&clients).Error; result != nil {
		return clients
	}
	//On renseigne la table users
	// Génère le hash du mot de passe avant de sauvegarder
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(clients.Password), bcrypt.DefaultCost)
	if err != nil {
		// tu peux gérer l’erreur comme tu veux, par exemple log.Fatal ou retourner une erreur
		panic("Erreur lors du hachage du mot de passe : " + err.Error())
	}
	fmt.Print("password saisie", clients.Password)
	var user models.User
	user.CreatedAt = time.Now()
	user.Role = "client"
	user.Email = clients.Email
	user.Name = clients.Name
	user.Password = string(hashedPassword)
	user.Ref = id.String()
	c.DB.Save(&user)
	return clients
}

// UpdateClient implements ClientsRepository.
func (c *clientRepository) UpdateClient(id int, clients models.Clients) models.Clients {
	var clientsExist models.Clients
	c.DB.Find(&clientsExist)
	if clients.Email == "" {
		clients.Email = clientsExist.Email
	}
	if clients.Mobile == "" {
		clients.Mobile = clientsExist.Mobile
	}
	if clients.Name == "" {
		clients.Name = clientsExist.Name
	}
	clients.ID = clientsExist.ID
	return clients
}

// Constructeur
func NewClientsRepository(repo *gorm.DB) ClientsRepository {
	return &clientRepository{
		DB: repo,
	}
}
