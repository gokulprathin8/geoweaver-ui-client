package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var LogoWidth int = 110
var LogoHeight int = 110

func SetGeoweaverLogo() *fyne.Container {

	geoweaverLogo := canvas.NewImageFromFile("assets/logo.png")
	logoContainer := container.New(layout.NewGridWrapLayout(fyne.NewSize(float32(LogoWidth), float32(LogoHeight))), geoweaverLogo)
	return logoContainer

}
