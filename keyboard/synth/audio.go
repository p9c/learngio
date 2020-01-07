package main

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

// Use SDL for audio output, because we already depend on it.

// typedef unsigned char Uint8;
// void AudioCallback(void *userdata, Uint8 *stream, int len);
import "C"

//export AudioCallback
func AudioCallback(userdata unsafe.Pointer, stream *C.Uint8, length C.int) {
	n := int(length)
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(stream)), Len: n / 4, Cap: n / 4}
	buf := *(*[]float32)(unsafe.Pointer(&hdr))

	player.L.Lock()
	copy(buf, player.buf)
	player.ready = true
	player.L.Unlock()
	player.Signal()
}

type AudioPlayer struct {
	dev   sdl.AudioDeviceID
	buf   []float32
	ready bool
	sync.Cond
}

var player *AudioPlayer

func (a *AudioPlayer) Init(size int) {
	a.L = &sync.Mutex{}
	player = a
	sdl.Init(sdl.INIT_AUDIO)
	a.open(size)
}

func (a *AudioPlayer) ChangeBufferSize(size int) {
	sdl.PauseAudioDevice(a.dev, true)
	sdl.CloseAudioDevice(a.dev)
	a.open(size)
}

func (a *AudioPlayer) open(samples int) {
	spec := &sdl.AudioSpec{
		Freq:     SampleRate,
		Format:   sdl.AUDIO_F32,
		Channels: 1,
		Samples:  uint16(samples),
		Callback: sdl.AudioCallback(C.AudioCallback),
	}
	id, err := sdl.OpenAudioDevice("", false, spec, spec, 0)
	if err != nil {
		fmt.Println(err)
	}
	a.dev = id
	a.buf = make([]float32, spec.Samples)
	sdl.PauseAudioDevice(id, false)
}
