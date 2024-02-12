package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

var cfg config

func main() {
	// create a fyne application
	a := app.New()

	// create a window for the app
	win := a.NewWindow("Go Markdown")

	// get the user interface
	edit, preview := cfg.makeUI()
	cfg.createMenuItems(win)

	// set the content of the window
	win.SetContent(container.NewHSplit(edit, preview))

	// show window and run app
	win.Resize(fyne.Size{Width: 800, Height: 500})
	win.CenterOnScreen()
	win.ShowAndRun()
}

func (app *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	app.EditWidget = edit
	app.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}

func (app *config) createMenuItems(win fyne.Window) {

	openMenuItem := fyne.NewMenuItem("Öffnen...", func() {

	})

	saveMenuItem := fyne.NewMenuItem("Speichern", func() {

	})

	saveAsMenuItem := fyne.NewMenuItem("Speicher als...", func() {

	})

	quit := fyne.NewMenuItem("Schließen", nil)
	quit.IsQuit = true

	fileMenu := fyne.NewMenu("Datei", openMenuItem, saveMenuItem, saveAsMenuItem, quit)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)
}
