package main

import (
	"fmt"
	"github.com/zatrasz75/just/logger"
	"zatrasz75/SkillsRock/configs"
	"zatrasz75/SkillsRock/internal/app"
)

const logFilePath = "./var/log/main.log"

func main() {
	l, err := logger.NewLogger(logFilePath)
	if err != nil {
		fmt.Println("Ошибка при создании файла логгера:", err)
		return
	}
	defer l.Close()

	// Configuration
	cfg, err := configs.NewConfig()
	if err != nil {
		l.Fatal("ошибка при разборе конфигурационного файла", err)
	}

	app.Run(cfg, l)
}
