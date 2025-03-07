package repoTasks

import (
	"context"
	"fmt"
	"github.com/zatrasz75/tools_postgres/psql"
	"time"
	"zatrasz75/SkillsRock/internal/models"
)

type Storage struct {
	*psql.Postgres
}

func New(pg *psql.Postgres) (*Storage, error) {
	var s = &Storage{pg}

	return s, nil
}

// CreateTask создает новую задачу в базе данных.
func (s *Storage) CreateTask(task *models.Task) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("не удалось запустить транзакцию: %w", err)
	}
	defer tx.Rollback(ctx)

	query := `
		INSERT INTO tasks (title, description, status)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var id int
	err = tx.QueryRow(ctx, query, task.Title, task.Description, task.Status).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("ошибка при создании задачи: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, fmt.Errorf("не удалось зафиксировать транзакцию: %w", err)
	}

	return id, nil
}

// GetAllTasks возвращает все задачи из базы данных.
func (s *Storage) GetAllTasks() ([]models.Task, error) {
	query := `
		SELECT id, title, description, status, created_at, updated_at
		FROM tasks
		LIMIT 100;
	`

	rows, err := s.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении задач: %w", err)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("ошибка при сканировании задачи: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при обработке результатов: %w", err)
	}

	return tasks, nil
}

// UpdateTask обновляет задачу в базе данных.
func (s *Storage) UpdateTask(task *models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("не удалось запустить транзакцию: %w", err)
	}
	defer tx.Rollback(ctx)

	query := `
		UPDATE tasks
		SET title = $1, description = $2, status = $3, updated_at = now()
		WHERE id = $4
	`

	_, err = tx.Exec(context.Background(), query, task.Title, task.Description, task.Status, task.ID)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении задачи: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("не удалось зафиксировать транзакцию: %w", err)
	}

	return nil
}

// DeleteTask удаляет задачу из базы данных по её ID.
func (s *Storage) DeleteTask(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("не удалось запустить транзакцию: %w", err)
	}
	defer tx.Rollback(ctx)

	query := `
		DELETE FROM tasks
		WHERE id = $1
	`

	_, err = tx.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("ошибка при удалении задачи: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("не удалось зафиксировать транзакцию: %w", err)
	}

	return nil
}
