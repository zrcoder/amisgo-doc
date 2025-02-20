package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/conf"
	"github.com/zrcoder/amisgo/theme"
)

func main() {
	app := amisgo.New(
		conf.WithThemes(
			theme.Theme{Value: theme.Cxd, Label: "Light"},
			theme.Theme{Value: theme.Dark, Label: "Dark"},
		),
	)
	app.Mount("/", app.Page().Body(
		app.ThemeButtonGroupSelect(),
		"Hello, World!",
	))
	app.Run(":8888")
}
