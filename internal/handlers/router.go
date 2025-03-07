package handlers

import (
	"github.com/gofiber/fiber/v2"
	_ "zatrasz75/SkillsRock/docs"
)

//	@title			Тестовое задание SkillsRock
//	@version		1.0

// @contact.name Михаил Токмачев
// @contact.url https://t.me/Zatrasz
// @contact.email zatrasz@ya.ru

// @host						localhost:3000
// @BasePath					/

// NewRouter -.
func NewRouter() *fiber.App {
	r := fiber.New()

	return r
}

func RegisterHandlers(r *fiber.App, s *Service) {
	g := r.Group("/tasks")

	g.Post("/", s.CreateTask)
	g.Get("/", s.GetAllTasks)
	g.Put("/:id", s.UpdateTask)
	g.Delete("/:id", s.DeleteTask)

	r.Static("/swagger", "./docs")
}
