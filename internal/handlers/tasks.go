package handlers

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"zatrasz75/SkillsRock/internal/models"
)

// CreateTask godoc
//
// @Summary Создайте нового задачи
// @Tags		Tasks
// @Description Принимает обязательные поля title и description. Поле status устанавливается автоматически в "new".
// @Accept  json
// @Produce  json
// @Param   task body models.Task true "Данные задачи"
// @Success 200 {object} models.Task "Успешно созданная задача"
// @Failure 400 {string} string "Ошибка парсинга тела запроса"
// @Failure 500 {string} string "Ошибка создания задачи"
// @Router /tasks [post]
func (s *Service) CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		s.l.Error("Ошибка парсинга тела запроса:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный формат данных",
		})
	}
	if task.Title == "" || task.Description == "" {
		s.l.Warn("Поле 'title' и 'description' обязательно для заполнения")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Поле 'title' и 'description' обязательно для заполнения",
		})
	}
	task.Status = "new"

	id, err := s.repo.Tasks.CreateTask(&task)
	if err != nil {
		s.l.Error("Ошибка создания задачи:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось создать задачу",
		})
	}

	s.l.Info("Добавлена задача с id -", id)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}

// GetAllTasks godoc
//
// @Summary Получение всех задач
// @Tags		Tasks
// @Description Возвращает список всех задач с лимитом 100 записей
// @Success 200 {object} models.Task "Успешно созданная задача"
// @Failure 500 {string} string "Ошибка получения задач"
// @Router /tasks [get]
func (s *Service) GetAllTasks(c *fiber.Ctx) error {
	tasks, err := s.repo.Tasks.GetAllTasks()
	if err != nil {
		s.l.Error("Ошибка получения задач:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось получить задачи",
		})
	}
	if tasks == nil {
		s.l.Error("Нет сохраненных задач:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Нет сохраненных задач",
		})
	}

	s.l.Info("Получение всех задач")

	return c.JSON(tasks)
}

// UpdateTask godoc
//
// @Summary Обновление задачи
// @Tags		Tasks
// @Description Обновляет задачу по её ID. Принимает обязательные поля title, description и status.
// @Accept  json
// @Produce  json
// @Param   id path int true "ID задачи"
// @Param   task body models.Task true "Данные для обновления задачи"
// @Success 200 {object} fiber.Map "Задача успешно обновлена"
// @Failure 400 {string} string "Ошибка парсинга тела запроса или неверный ID"
// @Failure 500 {string} string "Ошибка обновления задачи"
// @Router /tasks/{id} [put]
func (s *Service) UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID задачи обязателен",
		})
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "не удалось отформатировать id в int ",
		})
	}

	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		s.l.Error("Ошибка парсинга тела запроса:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный формат данных",
		})
	}
	if task.Title == "" || task.Description == "" || task.Status == "" {
		s.l.Warn("Поле 'title' , 'description' и 'status' обязательно для заполнения")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Поле 'title' , 'description' и 'status' обязательно для заполнения",
		})
	}
	task.ID = idInt

	err = s.repo.Tasks.UpdateTask(&task)
	if err != nil {
		s.l.Error("Ошибка обновления задачи:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось обновить задачу",
		})
	}

	s.l.Info("Обновлена задача c id", task.ID)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Задача успешно обновлена",
	})
}

// DeleteTask godoc
//
// @Summary Удаление задачи
// @Tags		Tasks
// @Description Удаляет задачу по её ID.
// @Produce  json
// @Param   id path int true "ID задачи"
// @Success 200 {object} fiber.Map "Задача успешно удалена"
// @Failure 400 {string} string "Неверный ID"
// @Failure 500 {string} string "Ошибка удаления задачи"
// @Router /tasks/{id} [delete]
func (s *Service) DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID задачи обязателен",
		})
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "не удалось отформатировать id в int ",
		})
	}

	err = s.repo.Tasks.DeleteTask(idInt)
	if err != nil {
		s.l.Error("Ошибка удаления задачи:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось удалить задачу",
		})
	}

	s.l.Info("Удалена задача с id", idInt)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Задача успешно удалена",
	})
}
