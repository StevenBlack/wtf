package wtf

import (
	"fmt"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
)

var Config *config.Config

type TextWidget struct {
	enabled         bool
	focusable       bool
	focusChar       string
	key             string
	name            string
	refreshInterval int

	View *tview.TextView

	CommonSettings *cfg.Common
	Position
}

func NewTextWidget(app *tview.Application, commonSettings *cfg.Common, focusable bool) TextWidget {
	configKey := commonSettings.ConfigKey

	focusChar := string('0' + commonSettings.FocusChar)
	if commonSettings.FocusChar == -1 {
		focusChar = ""
	}

	widget := TextWidget{
		CommonSettings: commonSettings,

		enabled:         commonSettings.Enabled,
		focusable:       focusable,
		focusChar:       focusChar,
		key:             commonSettings.ConfigKey,
		name:            commonSettings.Name,
		refreshInterval: commonSettings.RefreshInterval,
	}

	widget.Position = NewPosition(
		commonSettings.Position.Top,
		commonSettings.Position.Left,
		commonSettings.Position.Width,
		commonSettings.Position.Height,
	)

	widget.addView(app, configKey)

	return widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *TextWidget) BorderColor() string {
	if widget.Focusable() {
		return widget.CommonSettings.Colors.BorderFocusable
	}

	return widget.CommonSettings.Colors.BorderNormal
}

func (widget *TextWidget) ContextualTitle(defaultStr string) string {
	if widget.FocusChar() == "" {
		return fmt.Sprintf(" %s ", defaultStr)
	}

	return fmt.Sprintf(" %s [darkgray::u]%s[::-][green] ", defaultStr, widget.FocusChar())
}

func (widget *TextWidget) Disable() {
	widget.enabled = false
}

func (widget *TextWidget) Disabled() bool {
	return !widget.Enabled()
}

func (widget *TextWidget) Enabled() bool {
	return widget.enabled
}

func (widget *TextWidget) Focusable() bool {
	return widget.enabled && widget.focusable
}

func (widget *TextWidget) FocusChar() string {
	return widget.focusChar
}

// IsPositionable returns TRUE if the widget has valid position parameters, FALSE if it has
// invalid position parameters (ie: cannot be placed onscreen)
func (widget *TextWidget) IsPositionable() bool {
	return widget.Position.IsValid()
}

func (widget *TextWidget) Key() string {
	return widget.key
}

func (widget *TextWidget) Name() string {
	return widget.name
}

func (widget *TextWidget) RefreshInterval() int {
	return widget.refreshInterval
}

func (widget *TextWidget) SetFocusChar(char string) {
	widget.focusChar = char
}

func (widget *TextWidget) TextView() *tview.TextView {
	return widget.View
}

/* -------------------- Unexported Functions -------------------- */

func (widget *TextWidget) addView(app *tview.Application, configKey string) {
	view := tview.NewTextView()

	view.SetBackgroundColor(ColorFor(widget.CommonSettings.Colors.Background))
	view.SetBorderColor(ColorFor(widget.BorderColor()))
	view.SetTextColor(ColorFor(widget.CommonSettings.Colors.Text))
	view.SetTitleColor(ColorFor(widget.CommonSettings.Colors.Title))

	view.SetBorder(true)
	view.SetDynamicColors(true)
	view.SetTitle(widget.ContextualTitle(widget.name))
	view.SetWrap(false)

	view.SetChangedFunc(func() {
		app.Draw()
	})

	widget.View = view
}
