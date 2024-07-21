package forms

import (
	"fyneclient/api"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

// Form is an interface that defines a method to create a user interface
type Form struct {
	application fyne.App
	window      fyne.Window
	apiClient   *api.APIClient
}

// CustomTheme implements fyne.Theme interface
type CustomTheme struct{}

// Return a predefined color palette
func (m CustomTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground:
		return color.RGBA{R: 255, G: 248, B: 231, A: 255}
	case theme.ColorNamePrimary:
		return color.RGBA{R: 34, G: 139, B: 34, A: 255}
	case theme.ColorNameWarning:
		return color.RGBA{R: 255, G: 99, B: 71, A: 255}
	case theme.ColorNameButton:
		return color.RGBA{R: 255, G: 215, B: 0, A: 255}
	default:
		return theme.DefaultTheme().Color(n, v)
	}
}

// Return the font for the theme
func (m CustomTheme) Font(s fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(s)
}

// Return the size for the theme
func (m CustomTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}

// Return the icon for the theme
func (m CustomTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (f Form) Run() {
	f.application.Run()
}

func createHeader(pageTitle string) *fyne.Container {
	currentPageName := canvas.NewText(pageTitle, color.White)
	currentPageName.TextSize = 24
	logo := canvas.NewImageFromResource(resourceIconPng)
	logo.SetMinSize(fyne.NewSize(50, 50))

	headerBackground := canvas.NewRectangle(color.RGBA{R: 96, G: 180, B: 90, A: 255})

	header := container.NewHBox(
		logo,
		currentPageName,
		layout.NewSpacer(), // Spacer to push everything to the left
	)
	return container.NewMax(headerBackground, header)
}

// NewApplication creates a new form struct from baseUrl (the API URL)
func NewApplication(baseUrl string) *Form {
	a := app.NewWithID("balet.benjamin.cinephoriadesk")
	//a.Settings().SetTheme(&CustomTheme{})
	w := a.NewWindow("Cinéphoria Maintenance")
	w.Resize(fyne.NewSize(900, 600))
	iconPng := fyne.NewStaticResource("icon.png", resourceIconPng.StaticContent)
	a.SetIcon(iconPng)
	w.SetIcon(iconPng)
	w.Show()
	api := api.NewAPIClient(baseUrl)
	forms := &Form{
		application: a,
		window:      w,
		apiClient:   api,
	}
	forms.ShowLogin()
	return forms
}
