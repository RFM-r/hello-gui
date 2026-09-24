package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Создаём приложение и окно
	a := app.New()
	w := a.NewWindow("Приветствие")

	// Приветственная надпись
	greeting := widget.NewLabel("Привет, мир! 👋")
	greeting.TextStyle = fyne.TextStyle{Bold: true}
	greeting.Alignment = fyne.TextAlignCenter

	// Кнопка «Закрыть»
	closeBtn := widget.NewButton("Закрыть", func() {
		w.Close()
	})

	// Компоновка: надпись сверху, кнопка снизу по центру
	content := container.NewVBox(
		container.NewPadded(greeting),
		container.NewCenter(closeBtn),
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(320, 200))
	w.CenterOnScreen()
	w.ShowAndRun()
}
