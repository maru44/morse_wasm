package main

import (
	"github.com/maru44/morse/config"
	"github.com/maru44/morse/pkg/domain"
	"github.com/maru44/morse/pkg/execute"
	"github.com/maru44/morse/pkg/usecase"
)

func main() {
	// window := js.Global()

	// normalButton := window.Get("document").Call("getElementById", "normalButton")
	// cancelButton := window.Get("document").Call("getElementById", "cancelButton")

	ch := make(chan string)
	ret := ""

	settings := domain.MorseSetting{
		Input:  domain.InputMode(config.INPUT_MODE_BROWSER_BUTTON),
		Output: domain.OutputMode(config.OUTPUT_MODE_PLAIN_TEXT),
	}
	morse := domain.InitMorse(settings)
	mi := usecase.MorseInteractor(
		&execute.MorseRepository{
			Morse: morse,
		},
	)

	mi.Ignition()

}
