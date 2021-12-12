package timer

import (
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/lthh91/WAVPlayer"
	"github.com/repassyl/tawa/settings"
)

const (
	buttonCancel = "Cancel"
	buttonStart = "Start"
)

var zeroTime = time.Unix(0, 0).UTC()

type State struct {
	Length time.Duration
	Started time.Time
	Display *widget.Label
	Ticker *time.Ticker
	Button *widget.Button
}

type Timers map[string]State

var timers = make(Timers)

func (s State) update() {
	elapsed := time.Since(s.Started)
	formatted := zeroTime.Format("15:04:05")
	if s.Ticker != nil {
		formatted = zeroTime.Add(elapsed).Format("15:04:05")
	}
	s.Display.SetText(formatted)
	if s.Length >= elapsed {
		s.stop()
		Sound()
	}
}

func (s State) stop() {
	s.Ticker.Stop()
	s.Button.SetText(buttonStart)
}

func Init(label string) *fyne.Container {
	s := State{
		Length: settings.GetMinutes(label),
		Display: widget.NewLabel(""),
	}
	timers[label] = s
	s.Button = widget.NewButton(buttonStart, func() {
		if s.Button.Text == buttonStart {
			go func() {
				s.Started = time.Now()
				s.Button.SetText(buttonCancel)
				s.Ticker = time.NewTicker(500 * time.Millisecond) 
				for range s.Ticker.C {
					s.update()
				}
			}()
		} else if s.Button.Text == buttonCancel {
			s.stop()
		}
	})
	s.update()
	return container.New(layout.NewVBoxLayout(), s.Display, s.Button)
}

func Sound() {
	swoosh, _ := wavPlayer.NewWAV("tibetan-bell-sound.wav")
	swoosh.Play()
}