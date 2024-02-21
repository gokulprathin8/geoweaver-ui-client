package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

type forcedVariant struct {
	fyne.Theme

	variant fyne.ThemeVariant
}

func (f *forcedVariant) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return f.Theme.Color(name, f.variant)
}

func LightTheme(fallback fyne.Theme) fyne.Theme {
	return &forcedVariant{Theme: fallback, variant: 1}
}

func main() {
	app := app.New()
	app.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantLight})
	window := app.NewWindow("Geoweaver")

	logoWidth := 200
	logoHeight := 200
	geoweaverLogo := canvas.NewImageFromFile("assets/logo.png")
	fixedSizeContainer := container.New(layout.NewFixedGridLayout(fyne.NewSize(float32(logoWidth), float32(logoHeight))), geoweaverLogo)

	text2 := canvas.NewText("2", color.Black)
	text3 := canvas.NewText("3", color.Black)
	grid := container.New(layout.NewGridLayoutWithColumns(2), text1, text2, text3)
	window.SetContent(grid)

	window.ShowAndRun()

}
