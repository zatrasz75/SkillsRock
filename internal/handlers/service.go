package handlers

import (
	"github.com/zatrasz75/just/logger"
	"zatrasz75/SkillsRock/internal/repository"
)

type Service struct {
	l    logger.LoggersInterface
	repo *repository.Repository
}

func New(l logger.LoggersInterface, repo *repository.Repository) *Service {
	return &Service{
		l:    l,
		repo: repo,
	}
}
