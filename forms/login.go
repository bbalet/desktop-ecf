package forms

import (
	"log"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// ShowLogin creates the user interface for the login form
func (f Form) ShowLogin() {
	username := widget.NewEntry()
	username.SetPlaceHolder("Email")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Mot de passe")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Email", Widget: username},
			{Text: "Mot de passe", Widget: password},
		},
		SubmitText: "Se connecter",
		OnSubmit: func() {
			status, err := f.apiClient.Login(username.Text, password.Text)
			if err != nil {
				dialog.ShowError(err, f.window)
				return
			}
			if status == 200 {
				log.Println("Connecté avec succès")
				user, err := f.apiClient.WhoAmI()
				if err != nil {
					dialog.ShowError(err, f.window)
					return
				} else {
					if user.Role != "ROLE_EMPLOYEE" {
						log.Println("Problème de permission : " + user.Role)
						dialog.ShowInformation("Problème de connexion", "Vous n'avez pas le droit d'utiliser cette application", f.window)
					} else {
						f.ShowTheaters()
					}
				}
			} else {
				dialog.ShowInformation("Problème de connexion", "Vos identifiants semblent incorrects.", f.window)
			}
		},
	}

	content := container.NewVBox(
		form,
	)
	f.window.SetContent(container.NewBorder(createHeader("Cinéphoria Maintenance - Connexion à l'application"), nil, nil, nil, content))
}
