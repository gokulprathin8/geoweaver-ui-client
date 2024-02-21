package main

import (
	"errors"
	"image/color"

	"geoweaver-gui-client/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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
	window.Resize(fyne.NewSize(800, 800))

	// check for java
	isJavaInstalled := utils.CheckJavaInstallation()
	if !isJavaInstalled {
		errMsg := errors.New("Java is not installed. Please install Java to use to the application. (https://adoptopenjdk.net/releases.html)")
		dialog.ShowError(errMsg, window)
	}

	geoweaverLogo := utils.SetGeoweaverLogo()
	centeredLogo := container.New(layout.NewCenterLayout(), geoweaverLogo)

	startServerBtn := widget.NewButton("Start Geoweaver", nil)
	startServerBtn.Importance = widget.HighImportance
	startServerBtn.OnTapped = func() {
		if startServerBtn.Text == "Start Geoweaver" {
			startServerBtn.SetText("Stop Geoweaver")
			startServerBtn.Importance = widget.DangerImportance
			utils.StartServer()
		} else {
			startServerBtn.SetText("Start Geoweaver")
			startServerBtn.Importance = widget.HighImportance
		}
	}

	// Use a vertical box layout to stack the logo and the button
	content := container.NewVBox(
		centeredLogo,
		startServerBtn,
	)

	window.SetContent(content)

	window.ShowAndRun()
}
