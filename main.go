package main

import (
	"fyneclient/forms"
)

func main() {
	forms := forms.NewApplication("https://cinephoria.jorani.org")
	forms.Run()
}
