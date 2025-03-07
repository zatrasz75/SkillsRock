package repository

import (
	"github.com/zatrasz75/tools_postgres/psql"
	"zatrasz75/SkillsRock/internal/repoTasks"
)

type Repository struct {
	Tasks repoTasks.TasksInterface
}

func NewRepository(pg *psql.Postgres) (*Repository, error) {
	tasks, err := repoTasks.New(pg)
	if err != nil {
		return nil, err
	}

	return &Repository{
		Tasks: tasks,
	}, nil
}
