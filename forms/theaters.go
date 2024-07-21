package forms

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ShowTheaters creates the user interface for listing theaters
func (f Form) ShowTheaters() {
	theaters, _ := f.apiClient.TheaterAPI.GetListOfTheaters()

	theaterList := widget.NewList(
		func() int { return len(theaters) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i int, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(theaters[i].City)
		},
	)

	theaterList.OnSelected = func(id widget.ListItemID) {
		f.ShowRooms(theaters[id])
	}

	scroll := container.NewVScroll(theaterList)
	scroll.SetMinSize(fyne.NewSize(600, 500))

	closeButton := widget.NewButton("Fermer", func() {
		f.window.Close()
	})

	logoutButton := widget.NewButton("Se déconnecter", func() {
		f.apiClient.Logout()
		f.ShowLogin()
	})

	content := container.NewVBox(
		widget.NewLabel("Sélectionnez un cinéma:"),
		scroll,
		container.NewHBox(closeButton, logoutButton),
	)
	f.window.SetContent(container.NewBorder(createHeader("Cinéphoria Maintenance - Liste des cinémas"), nil, nil, nil, content))
}
