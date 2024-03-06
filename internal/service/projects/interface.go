package projects

import (
	"context"

	"github.com/KadyrPoyraz/httplayout/internal/domain"
)

type Service interface {
    GetProjects(ctx context.Context) ([]domain.Project, error)
}
