package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gordonklaus/portaudio"
	"github.com/zenwerk/go-wave"
)

func main() {
	fmt.Println("Recording to test.wav -- Ctrl+C to stop recording")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	fileName := "test.wav"

	waveFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	inputChannels := 1
	outputChannels := 0
	sampleRate := 44100
	framesPerBuffer := make([]float32, 64)

	portaudio.Initialize()

	stream, err := portaudio.OpenDefaultStream(inputChannels, outputChannels, float64(sampleRate), len(framesPerBuffer), framesPerBuffer)

	param := wave.WriterParam{
		Out:           waveFile,
		Channel:       inputChannels,
		SampleRate:    sampleRate,
		BitsPerSample: 8,
	}

	waveWriter, err := wave.NewWriter(param)

	if err != nil {
		panic(err)
	}

	stream.Start()
	for {
		stream.Read()

		wavBuffer := make([]byte, len(framesPerBuffer))
		for i, num := range framesPerBuffer {
			val := num
			if val < -1.0 {
				val = -1.0
			}
			if val > 1.0 {
				val = 1.0
			}
			var valInt = byte(((val + 1.0) * 127))
			wavBuffer[i] = valInt
		}

		_, err := waveWriter.Write([]byte(wavBuffer))

		if err != nil {
			panic(err)
		}

		select {
		case <-sig:
			stream.Stop()
			waveWriter.Close()
			stream.Close()
			portaudio.Terminate()
			os.Exit(0)
		default:
		}
	}
}
