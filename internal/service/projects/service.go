package projects

import (
	"context"

	"github.com/KadyrPoyraz/httplayout/internal/domain"
	"github.com/KadyrPoyraz/httplayout/internal/repository/db"
	"github.com/KadyrPoyraz/httplayout/internal/repository/projects"
)

type service struct {
	postgresqlDB db.DB
    projectsRepo projects.Repo
}

func New(
    db db.DB,
	projectsRepo projects.Repo,
) Service {
	return &service{
		projectsRepo: projectsRepo,
        postgresqlDB: db,
	}
}

func (s *service) GetProjects(ctx context.Context) ([]domain.Project, error) {
    tx, err := s.postgresqlDB.StartTransaction(ctx)
    if err != nil {
        return nil, err
    }

    projects, err := s.projectsRepo.GetProjects(ctx, tx)
    if err != nil {
        return nil, err
    }

    return projects, nil
}

func (s *service) SayHello(ctx context.Context) ([]string, error) {
	return nil, nil
}
