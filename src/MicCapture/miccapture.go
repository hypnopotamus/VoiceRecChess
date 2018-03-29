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
	framesPerBuffer := make([]byte, 64)

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

	go func() {
		var input string
		fmt.Scanln(&input)

		stream.Stop()
		waveWriter.Close()
		stream.Close()
		portaudio.Terminate()
		os.Exit(0)
	}()

	defer func() {
		stream.Stop()
		waveWriter.Close()
		stream.Close()
		portaudio.Terminate()
		os.Exit(0)
	}()

	stream.Start()
	for {
		stream.Read()

		_, err := waveWriter.Write([]byte(framesPerBuffer))

		if err != nil {
			panic(err)
		}

		select {
		case <-sig:
			return
		default:
		}
	}
}
