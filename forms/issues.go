package forms

import (
	"fyneclient/api"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// ShowRooms creates the user interface for listing theaters
func (f Form) ShowIssues(room api.Room, theater api.Theater) {
	issues, _ := f.apiClient.IssueAPI.GetListOfIssues(room.RoomId)

	issueList := widget.NewList(
		func() int { return len(issues) },
		func() fyne.CanvasObject {
			// Create a container to hold the status and title
			return container.NewHBox(
				widget.NewLabel(""),
				widget.NewLabel(""),
			)
		},
		func(i int, item fyne.CanvasObject) {
			// Cast the item to a container and update the labels
			box := item.(*fyne.Container)
			statusLabel := box.Objects[0].(*widget.Label)
			titleLabel := box.Objects[1].(*widget.Label)
			statusLabel.SetText(issues[i].Status)
			titleLabel.SetText(issues[i].Title)
		},
	)

	issueList.OnSelected = func(id widget.ListItemID) {
		f.showIssue(room, issues[id], theater)
	}

	newProblemButton := widget.NewButton("Nouveau Problème", func() {
		f.showNewIssue(room, theater)
	})

	scroll := container.NewVScroll(issueList)
	scroll.SetMinSize(fyne.NewSize(600, 500))

	homeButton := widget.NewButton("Accueil", func() {
		// Return to the home page
		f.ShowTheaters()
	})

	backButton := widget.NewButton("Retour", func() {
		// Return to the home page
		f.ShowRooms(theater)
	})

	content := container.NewVBox(
		widget.NewLabel("Sélectionnez un problème:"),
		scroll,
		container.NewHBox(backButton, homeButton, newProblemButton),
	)
	f.window.SetContent(container.NewBorder(createHeader("Cinéphoria Maintenance - Liste des problèmes"), nil, nil, nil, content))
}

// create the new issue and display back the list of issues
func (f Form) showNewIssue(room api.Room, theater api.Theater) {
	titleEntry := widget.NewEntry()
	descriptionEntry := widget.NewEntry()
	statusSelect := widget.NewSelect([]string{"Nouveau", "Ouvert", "Terminé"}, func(selected string) {})
	statusSelect.SetSelected("Nouveau")

	form := widget.NewForm(
		widget.NewFormItem("Titre", titleEntry),
		widget.NewFormItem("Description", descriptionEntry),
		widget.NewFormItem("Statut", statusSelect),
	)
	saveButton := widget.NewButton("Enregistrer", func() {
		newIssue := api.Issue{
			RoomId:      room.RoomId,
			Title:       titleEntry.Text,
			Description: descriptionEntry.Text,
			Status:      statusSelect.Selected,
		}
		status, err := f.apiClient.IssueAPI.CreateNewIssue(newIssue)
		if err != nil {
			dialog.ShowError(err, f.window)
			return
		}
		if status == 201 {
			f.ShowIssues(room, theater)
		} else {
			dialog.ShowInformation("Problème de création", "Les informations saisies semblent incorrectes.", f.window)
		}
	})
	cancelButton := widget.NewButton("Annuler", func() {
		f.ShowIssues(room, theater)
	})

	content := container.NewVBox(
		form,
		container.NewHBox(saveButton, cancelButton),
	)
	f.window.SetContent(container.NewBorder(createHeader("Cinéphoria Maintenance - Créer un nouveau problème"), nil, nil, nil, content))
}

// Allow to view an Issue and to update it
func (f Form) showIssue(room api.Room, issue api.Issue, theater api.Theater) {
	titleEntry := widget.NewEntry()
	titleEntry.SetText(issue.Title)
	descriptionEntry := widget.NewEntry()
	descriptionEntry.SetText(issue.Description)
	statusSelect := widget.NewSelect([]string{"Nouveau", "Ouvert", "Terminé"}, func(selected string) {
		// This function will be called when an option is selected
	})
	statusSelect.SetSelected(issue.Status) // Set the current status

	form := widget.NewForm(
		widget.NewFormItem("Titre", titleEntry),
		widget.NewFormItem("Description", descriptionEntry),
		widget.NewFormItem("Statut", statusSelect),
	)

	saveButton := widget.NewButton("Enregistrer", func() {
		issue.Title = titleEntry.Text
		issue.Description = descriptionEntry.Text
		issue.Status = statusSelect.Selected
		// You would typically update the problem in the data source here
		status, err := f.apiClient.IssueAPI.UpdateIssue(issue)
		if err != nil {
			dialog.ShowError(err, f.window)
			return
		}
		if status == 200 {
			f.ShowIssues(room, theater)
		} else {
			dialog.ShowInformation("Problème de mise à jour", "Les informations saisies semblent incorrectes.", f.window)
		}
		f.ShowIssues(room, theater)
	})

	cancelButton := widget.NewButton("Annuler", func() {
		// Return to the problem list without making changes
		f.ShowIssues(room, theater)
	})

	content := container.NewVBox(
		form,
		container.NewHBox(saveButton, cancelButton),
	)
	f.window.SetContent(container.NewBorder(createHeader("Cinéphoria Maintenance - Modifier un problème"), nil, nil, nil, content))
}
