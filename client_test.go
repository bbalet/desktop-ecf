package main

import (
	"fyneclient/forms"
	"testing"
)

func TestAdd(t *testing.T) {
	forms := forms.NewApplication("https://cinephoria.jorani.org")
	forms.Run()

	//assert.Equal(t, "2", calc.output.Text)
}
