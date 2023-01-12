package internal

import (
	"fmt"

	"github.com/bspaans/bleep/audio"
	"github.com/hajimehoshi/oto/v2"
)

type OtoBuffer struct {
	samples []int
}

func (ob OtoBuffer) Read(b []byte) (n int, err error) {
	for _, v := range ob.samples {
		b = append(b, byte(v))
	}
	return len(ob.samples), nil
}

var CurrentOtoSink *OtoSink

type OtoSink struct {
	ctx        *oto.Context
	Config     *audio.AudioConfig
	BufferSize int
	GetSamples func(cfg *audio.AudioConfig, n int) []int
	open       bool
}

func NewOtoSink(cfg *audio.AudioConfig) (*OtoSink, error) {
	s := &OtoSink{
		Config:     cfg,
		BufferSize: 100,
	}
	CurrentOtoSink = s
	otoCtx, readyChan, err := oto.NewContext(s.Config.SampleRate, s.Config.GetNumberOfChannels(), s.Config.BitDepth/8)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	<-readyChan
	s.ctx = otoCtx
	return s, nil
}

func (s *OtoSink) Start(f func(cfg *audio.AudioConfig, n int) []int) error {
	s.open = true
	go func() {
		for {
			samples := f(s.Config, s.BufferSize)
			buffer := OtoBuffer{
				samples: samples,
			}
			player := s.ctx.NewPlayer(buffer)

			player.Play()

			// for player.IsPlaying() {
			// 	time.Sleep(time.Millisecond)
			// }
			fmt.Println(samples)
			err := player.Close()
			if err != nil {
				panic("player.Close failed: " + err.Error())
			}
			if !s.open {
				break
			}
		}
	}()
	return nil
}

func (s *OtoSink) Close(cfg *audio.AudioConfig) error {
	s.open = false
	return nil
}
