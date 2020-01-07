package main

import (
	"github.com/jfreymuth/ui/impl/gofont"
	"github.com/jfreymuth/ui/impl/sdl"
	"github.com/jfreymuth/ui/toolkit"
)

func main() {
	synth := NewSynth(8)
	synth.SetAttack(.01)
	synth.SetRelease(.02)

	keyboard := NewKeyboard()
	keyboard.NoteOn = synth.NoteOn
	keyboard.NoteOff = synth.NoteOff

	audio := &AudioPlayer{}
	audio.Init(1024)
	go func() {
		for {
			audio.L.Lock()
			for !audio.ready {
				audio.Wait()
			}
			for i := range audio.buf {
				audio.buf[i] = 0
			}
			synth.Process(audio.buf)
			audio.ready = false
			audio.L.Unlock()
		}
	}()

	sdl.Show(sdl.Options{
		Title:      "UI Demo",
		FontLookup: gofont.Lookup,
		Root: &toolkit.Container{Bottom: keyboard, Center: toolkit.NewBar(100,
			NewSlider("Volume", 1, synth.SetVolume),
			NewSlider("Attack", 0, synth.SetAttack),
			NewSlider("Release", .2, synth.SetRelease),
			NewSlider("Modulation", .6, synth.SetModulation),
		)},
	})
}
