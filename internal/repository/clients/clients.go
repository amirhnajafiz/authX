package clients

import (
	"github.com/amirhnajafiz/authX/internal/model"

	"gorm.io/gorm"
)

// Clients manages the client model.
type Clients interface {
	Create(client *model.Client) error
	GetSingle(clientID string) (*model.Client, error)
}

// New generates a new client repository.
func New(db *gorm.DB) Clients {
	return &clients{
		db: db,
	}
}

// clients manages the functions of repository.
type clients struct {
	db *gorm.DB
}

// Create a new client.
func (a *clients) Create(client *model.Client) error {
	return a.db.Create(client).Error
}

// GetSingle client.
func (a *clients) GetSingle(clientID string) (*model.Client, error) {
	client := new(model.Client)

	if err := a.db.First(&client).Where("client_id = ?", clientID).Error; err != nil {
		return nil, err
	}

	return client, nil
}
