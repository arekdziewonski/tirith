package ui

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func NewMainWindow() *gtk.Window {
	wnd := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	wnd.SetTitle("tirith")
	wnd.SetBorderWidth(3)
	wnd.SetSizeRequest(400, 647)
	wnd.SetPosition(gtk.WIN_POS_CENTER)
	wnd.SetIconName("gtk-dialog-info")
	wnd.Connect("destroy", func(ctx *glib.CallbackContext) {
		println("got destroy!", ctx.Data().(string))
		gtk.MainQuit()
	}, "main window")

	quit := gtk.NewButton()
	quit.SetLabel("Quit")
	quit.Clicked(func() {
		gtk.MainQuit()
	})

	vbox := gtk.NewVBox(false, 3)
	vbox.Add(gtk.NewLabel("Hello, World!"))
	vbox.Add(quit)

	wnd.Add(vbox)

	return wnd
}
