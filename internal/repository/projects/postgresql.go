package projects

import (
	"context"

	"github.com/KadyrPoyraz/httplayout/internal/domain"
	"github.com/KadyrPoyraz/httplayout/internal/repository/db"
)

type repo struct {
	databse *db.PostgresqlDB
}

func NewPostgresqlRepo(database *db.PostgresqlDB) Repo {
	return &repo{
		databse: database,
	}
}

func (r *repo) GetProjects(ctx context.Context, tx db.Transaction) ([]domain.Project, error) {
	query := `
        select id, name from projects
    `
    rows, err := tx.Tx().QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }

    var projects []domain.Project
    for rows.Next() {
        var project domain.Project
        err = rows.Scan(&project.Id, &project.Name)
        if err != nil {
            return nil, err
        }

        projects = append(projects, project)
    }

    return projects, nil
}
