package internal

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/bid"
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/doc"
	"github.com/oreshkindev/profilss.ru-backend/internal/post"
	"github.com/oreshkindev/profilss.ru-backend/internal/product"
	"github.com/oreshkindev/profilss.ru-backend/internal/user"
)

type Manager struct {
	Doc     *doc.Manager
	Bid     *bid.Manager
	Post    *post.Manager
	User    *user.Manager
	Product *product.Manager
}

func NewManager(database *database.Database) *Manager {
	return &Manager{
		Doc:     doc.NewManager(),
		Bid:     bid.NewManager(database),
		Post:    post.NewManager(database),
		User:    user.NewManager(database),
		Product: product.NewManager(database),
	}
}
