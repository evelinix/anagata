//go:build windows

package errors

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// ShowDialog displays an error dialog on Windows.
// Returns true if user clicked Retry.
func ShowDialog(err *AppError) bool {
	var dlg *walk.Dialog
	retry := false

	buttons := []Widget{
		PushButton{
			Text: "Close",
			OnClicked: func() {
				dlg.Close(0)
			},
		},
	}

	if err.Retriable {
		buttons = append([]Widget{
			PushButton{
				Text: "Retry",
				OnClicked: func() {
					retry = true
					dlg.Close(0)
				},
			},
		}, buttons...)
	}

	Dialog{
		AssignTo: &dlg,
		Title:    "AnagataSentinel",
		MinSize:  Size{Width: 400, Height: 200},
		MaxSize:  Size{Width: 400, Height: 200},
		Layout:   VBox{},
		Children: []Widget{
			VSpacer{Size: 10},
			Label{
				Text: err.Error(),
				Font: Font{Family: "Segoe UI Variable", PointSize: 10},
			},
			VSpacer{Size: 5},
			Label{
				Text:      "Stage: " + err.Stage,
				Font:      Font{Family: "Segoe UI Variable", PointSize: 8},
				TextColor: walk.RGB(120, 120, 120),
			},
			VSpacer{Size: 15},
			Composite{
				Layout: HBox{},
				Children: buttons,
			},
			VSpacer{Size: 10},
		},
	}.Create(dlg)

	dlg.Run()
	return retry
}
