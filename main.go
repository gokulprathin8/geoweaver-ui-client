package main

import (
	"errors"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"

	"geoweaver-gui-client/utils"
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
	window.Canvas().Content().MinSize().Min(fyne.NewSize(800, 800))

	// check for java
	isJavaInstalled := utils.CheckJavaInstallation()
	if isJavaInstalled {
		errMsg := errors.New("Java is not installed. Please install Java to use to the application. (https://adoptopenjdk.net/releases.html)")
		dialog.ShowError(errMsg, window)
	}

	text1 := utils.SetGeoweaverLogo()
	text2 := canvas.NewText("2", color.Black)
	text3 := canvas.NewText("3", color.Black)
	grid := container.New(layout.NewGridLayoutWithColumns(2), text1, text2, text3)
	window.SetContent(grid)

	window.ShowAndRun()

}
