package uibuttons

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dragon162/go-get-games/games/common/events"
	"image/color"
)

type renderer struct {
	fyne.WidgetRenderer
	icon  *canvas.Image
	label *canvas.Text

	objects []fyne.CanvasObject
	button  *Button
}

const bugSize = 18

// MinSize calculates the minimum size of a bug button. A fixed amount.
func (b *renderer) MinSize() fyne.Size {
	return fyne.NewSize(bugSize+theme.Padding()*2, bugSize+theme.Padding()*2)
}

// Layout the components of the widget
func (b *renderer) Layout(size fyne.Size) {
	inner := size.Subtract(fyne.NewSize(theme.Padding()*2, theme.Padding()*2))
	b.icon.Resize(inner)
	b.icon.Move(fyne.NewPos(theme.Padding(), theme.Padding()))

	textSize := size.Height * .67
	textMin := fyne.MeasureText(b.label.Text, textSize, fyne.TextStyle{Bold: true})

	b.label.TextSize = textSize
	b.label.Resize(fyne.NewSize(size.Width, textMin.Height))
	b.label.Move(fyne.NewPos(0, (size.Height-textMin.Height)/2))
}

// ApplyTheme is called when the Button may need to update its look
func (b *renderer) ApplyTheme() {
	b.label.Color = theme.ForegroundColor()
	b.Refresh()
}

func (b *renderer) BackgroundColor() color.Color {
	return theme.ButtonColor()
}

func (b *renderer) Refresh() {
	b.label.Text = b.button.text

	b.icon.Hidden = b.button.icon == nil
	if b.button.icon != nil {
		b.icon.Resource = b.button.icon
	}

	b.Layout(b.button.Size())
	canvas.Refresh(b.button)
}

func (b *renderer) Objects() []fyne.CanvasObject {
	return b.objects
}

func (b *renderer) Destroy() {
}

// Button widget is a scalable button that has a text label and icon and triggers an event func when clicked
type Button struct {
	widget.BaseWidget
	text string
	icon fyne.Resource

	tap *events.Feed[bool]
}

// Tapped is called when a regular tap is reported
func (b *Button) Tapped(ev *fyne.PointEvent) {
	b.tap.Send(true)
}

// TappedSecondary is called when an alternative tap is reported
func (b *Button) TappedSecondary(ev *fyne.PointEvent) {
	b.tap.Send(false)
}

func (b *Button) CreateRenderer() fyne.WidgetRenderer {
	text := canvas.NewText(b.text, theme.ForegroundColor())
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle.Bold = true

	icon := canvas.NewImageFromResource(b.icon)
	icon.FillMode = canvas.ImageFillContain

	objects := []fyne.CanvasObject{
		text,
		icon,
	}

	return &renderer{
		icon:    icon,
		label:   text,
		objects: objects,
		button:  b,
	}
}

// SetText allows the button label to be changed
func (b *Button) SetText(text string) {
	b.text = text
	b.Refresh()
}

// SetIcon updates the icon on a label - pass nil to hide an icon
func (b *Button) SetIcon(icon fyne.Resource) {
	b.icon = icon
	b.Refresh()
}

func (b *Button) GetClickEvent() *events.Feed[bool] {
	return b.tap
}

// NewButton creates a new button widget with the specified label, themed icon and tap handler
func NewButton(label string, icon fyne.Resource, tap func(bool)) *Button {
	button := &Button{text: label, icon: icon, tap: events.Make[bool]()}
	button.tap.Subscribe(tap)
	button.ExtendBaseWidget(button)
	return button
}
