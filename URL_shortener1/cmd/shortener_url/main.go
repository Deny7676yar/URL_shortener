package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Deny7676yar/URL_shortener/URL_shortener1/internal/infrastructure/api/handler"
	"github.com/Deny7676yar/URL_shortener/URL_shortener1/internal/infrastructure/api/routergin"
	"github.com/Deny7676yar/URL_shortener/URL_shortener1/internal/infrastructure/db/pgstore"
	"github.com/Deny7676yar/URL_shortener/URL_shortener1/internal/infrastructure/server"
	"github.com/Deny7676yar/URL_shortener/URL_shortener1/internal/usecase/app/repo"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	//ust := memory.NewLinks()
	//ust, err := userfilemanager.NewUsers("./data.json", "mem://userRefreshTopic")
	lst, err := pgstore.NewLinks(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	us := repo.NewLinks(lst)
	hs := handler.NewHandlers(us)
	// h := defmux.NewRouter(hs)
	h := routergin.NewRouterGin(hs)
	//h := routeropenapi.NewRouterOpenAPI(hs)
	srv := server.NewServer(":8000", h)

	srv.Start(us)
	log.WithFields(log.Fields{
		"Start": time.Now(),
	}).Info()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)
	for {
		select {
		case <-ctx.Done():
			return
		case <-sigCh:
			log.WithFields(log.Fields{
				"SIGINT": <-sigCh,
			}).Info("cencel context")
			srv.Stop()
			cancel() //Если пришёл сигнал SigInt - завершаем контекст
			lst.Close()
		}
	}
}
