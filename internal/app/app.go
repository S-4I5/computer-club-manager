package app

import (
	"bufio"
	"computer-club-manager/internal/config"
	"computer-club-manager/internal/manager"
)

type App struct {
	in      *bufio.Reader
	out     *bufio.Writer
	config  config.Config
	manager manager.Manager
}

func NewApp(
	in *bufio.Reader,
	out *bufio.Writer,
	config config.Config,
) *App {

	serviceProvider := newServiceProvider(config, in, out)

	return &App{
		in:      in,
		out:     out,
		config:  config,
		manager: serviceProvider.Manager(),
	}
}

func (a *App) Start() error {
	return a.manager.Start()
}
