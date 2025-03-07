package app

import (
	"github.com/gofiber/swagger"
	"github.com/zatrasz75/just/logger"
	"github.com/zatrasz75/tools_postgres/psql"
	"os"
	"os/signal"
	"syscall"
	"zatrasz75/SkillsRock/configs"
	"zatrasz75/SkillsRock/internal/handlers"
	"zatrasz75/SkillsRock/internal/httpServer"
	"zatrasz75/SkillsRock/internal/repository"
)

func Run(cfg *configs.Config, l logger.LoggersInterface) {
	pg, err := psql.New(cfg.ConnStr, psql.OptionSet(cfg.PoolMax, cfg.ConnAttempts, cfg.ConnTimeout))
	if err != nil {
		l.Fatal("ошибка запуска - postgres.New:", err)
	}
	defer pg.Close()

	err = pg.Migrate()
	if err != nil {
		l.Fatal("Ошибка выполнения миграций:", err)
	}

	repo, err := repository.NewRepository(pg)
	if err != nil {
		l.Error("ошибка repository:", err)
		return
	}

	router := handlers.NewRouter()
	router.Get("/swagger/*", swagger.HandlerDefault)

	api := handlers.New(l, repo)

	handlers.RegisterHandlers(router, api)

	srv := httpServer.New(router, httpServer.OptionSet(cfg.AddrHost, cfg.AddrPort, cfg.ShutdownTime))
	go func() {
		err = srv.Start()
		if err != nil {
			l.Warn("Остановка сервера:", err)
		}
	}()

	l.Info("Запуск сервера на http://" + cfg.AddrHost + ":" + cfg.AddrPort)
	l.Info("Документация Swagger API: http://" + cfg.AddrHost + ":" + cfg.AddrPort + "/swagger/index.html")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Trace("принят сигнал прерывания прерывание", s.String())
	case err = <-srv.Notify():
		l.Error("получена ошибка сигнала прерывания сервера", err)
	}

	err = srv.Shutdown()
	if err != nil {
		l.Error("не удалось завершить работу сервера", err)
	}
}
