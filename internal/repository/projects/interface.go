package projects

import (
	"context"

	"github.com/KadyrPoyraz/httplayout/internal/domain"
	"github.com/KadyrPoyraz/httplayout/internal/repository/db"
)

type Repo interface {
    GetProjects(ctx context.Context, tx db.Transaction) ([]domain.Project, error)
}
