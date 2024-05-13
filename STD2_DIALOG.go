package main

import (
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) OpenFileDialog() string {
	result, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "Question",
		DefaultDirectory: "\\\\ME-Datenbank-1\\Projektdaten 2024",
	})
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (a *App) OpenMultipleFilesDialog() []string {
	result, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "Choose File",
		DefaultDirectory: "\\\\ME-Datenbank-1\\Projektdaten 2024",
	})
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (a *App) OpenDirectoryDialog() string {
	result, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "Question",
		DefaultDirectory: "\\\\ME-Datenbank-1\\Database\\Schnittstelle",
	})
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (a *App) SaveFileDialog() string {
	result, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title:            "Question",
		DefaultDirectory: "\\\\ME-Datenbank-1\\Database\\Schnittstelle",
	})
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (a *App) MessageDialog() string {
	result, _ := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         "Question",
		Message:       "Do you want to continue?",
		DefaultButton: "No",
	})
	return result
}
