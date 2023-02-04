package music

import (
	"bytes"
	_ "embed"
	"io"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/sirupsen/logrus"
)

//go:embed intro.mp3
var introMp3 []byte

//go:embed suspense.mp3
var suspenseMp3 []byte

var Songs = map[string][]byte{
	"intro":    introMp3,
	"suspense": suspenseMp3,
}

func Init() {
	logrus.Debugf("Started - MusicInit()")

	for {
		for song := range Songs {
			logrus.Infof("Playing '%v'", song)
			Play(song)
			time.Sleep(60 * time.Second)
		}
	}
}

func Play(name string) {
	f := io.NopCloser(bytes.NewReader(Songs[name]))

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		logrus.Errorf("Error: %v", err)
		return
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}
