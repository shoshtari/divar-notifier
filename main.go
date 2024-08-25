package main

import (
	"fmt"

	"github.com/shoshtari/divar-notifier/internal/configs"
	"github.com/shoshtari/divar-notifier/internal/divar"
	"github.com/shoshtari/divar-notifier/internal/jarchi"
	"github.com/shoshtari/divar-notifier/internal/notify"
)

func main() {
	serviceConfig, err := configs.GetConfig()
	if err != nil {
		panic(err)
	}

	divarClient := divar.NewDivarClient()
	notifier := notify.NewNotifier(serviceConfig.Notifier)

	jarchiService := jarchi.New(divarClient, notifier)
	if err := jarchiService.Start(); err != nil {
		panic(err)
	}
	fmt.Println("go")
}
