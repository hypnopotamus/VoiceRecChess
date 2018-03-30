package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gordonklaus/portaudio"
	"github.com/zenwerk/go-wave"
)

func main() {
	var input string

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	for {
		fmt.Println("Press enter to start recording (5 second limit), CTRL+C to quit")
		fmt.Scanln(&input)
		select {
		case <-sig:
			os.Exit(0)
		default:
		}
		recording := true
		go CaptureAudio(&recording)
		time.Sleep(time.Duration(5) * time.Second)
		recording = false
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}

/*
CaptureAudio -- captures audio until the passed in boolean pointer is false
*/
func CaptureAudio(recording *bool) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	fileName := "temp.wav"

	waveFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	inputChannels := 1
	outputChannels := 0
	sampleRate := 16000
	framesPerBuffer := make([]float32, 64)

	portaudio.Initialize()

	stream, err := portaudio.OpenDefaultStream(inputChannels, outputChannels, float64(sampleRate), len(framesPerBuffer), framesPerBuffer)

	param := wave.WriterParam{
		Out:           waveFile,
		Channel:       inputChannels,
		SampleRate:    sampleRate,
		BitsPerSample: 16,
	}

	waveWriter, err := wave.NewWriter(param)

	if err != nil {
		panic(err)
	}

	stream.Start()
	for {
		stream.Read()

		wavBuffer := make([]uint16, len(framesPerBuffer))
		for i, num := range framesPerBuffer {
			val := (num + 1.0) / 2.0 * 65536.0
			if val < 0.0 {
				val = 0.0
			}
			if val > 65535.0 {
				val = 65535.0
			}
			var valInt = uint16(val+0.5) - 32768
			wavBuffer[i] = valInt
		}

		_, err := waveWriter.WriteSample16([]uint16(wavBuffer))

		if err != nil {
			panic(err)
		}

		if !(*recording) {
			stream.Stop()
			waveWriter.Close()
			stream.Close()
			portaudio.Terminate()
			return
		}
	}
}
