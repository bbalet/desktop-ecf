package forms

import (
	"fyneclient/api"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ShowRooms creates the user interface for listing theaters
func (f Form) ShowRooms(theater api.Theater) {
	rooms, _ := f.apiClient.RoomAPI.GetListOfRooms(theater.TheaterID)

	roomList := widget.NewList(
		func() int { return len(rooms) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i int, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(rooms[i].Number)
		},
	)

	roomList.OnSelected = func(id widget.ListItemID) {
		f.ShowIssues(rooms[id], theater)
	}

	scroll := container.NewVScroll(roomList)
	scroll.SetMinSize(fyne.NewSize(600, 500))

	homeButton := widget.NewButton("Retour", func() {
		// Return to the home page
		f.ShowTheaters()
	})

	content := container.NewVBox(
		widget.NewLabel("Sélectionnez une salle:"),
		scroll,
		container.NewHBox(homeButton),
	)
	f.window.SetContent(container.NewBorder(createHeader("Cinéphoria Maintenance - Liste des salles"), nil, nil, nil, content))
}
