package main

import (
	"github.com/arekdziewonski/tirith/internal/ui"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(nil)
	mainWnd := ui.NewMainWindow()
	mainWnd.ShowAll()
	gtk.Main()
}
