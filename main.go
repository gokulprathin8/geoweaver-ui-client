package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type forcedVariant struct {
	fyne.Theme

	variant fyne.ThemeVariant
}

func (f *forcedVariant) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return f.Theme.Color(name, f.variant)
}

func LightTheme(fallback fyne.Theme) fyne.Theme {
	return &forcedVariant{Theme: fallback, variant: 1} // avoid import loops
}

func main() {
	app := app.New()
	// app.Settings().SetTheme()
	window := app.NewWindow("Geoweaver")

	text1 := canvas.NewText("1", color.White)
	text2 := canvas.NewText("2", color.White)
	text3 := canvas.NewText("3", color.White)
	grid := container.New(layout.NewGridLayoutWithColumns(2), text1, text2, text3)
	window.SetContent(grid)
	window.ShowAndRun()

}
