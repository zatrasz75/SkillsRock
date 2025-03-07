package repoTasks

import "zatrasz75/SkillsRock/internal/models"

type TasksInterface interface {
	CreateTask(task *models.Task) (int, error)
	GetAllTasks() ([]models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id int) error
}
