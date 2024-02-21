package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func SetGeoweaverLogo() *fyne.Container {

	LogoWidth := 200
	LogoHeight := 200

	geoweaverLogo := canvas.NewImageFromFile("../assets/logo.png")
	logoContainer := container.New(layout.NewGridWrapLayout(fyne.NewSize(float32(LogoWidth), float32(LogoHeight))), geoweaverLogo)
	return logoContainer

}
