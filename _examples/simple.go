package main

import (
	"github.com/coral/fluidsynth2"
)

func main() {

	s := fluidsynth2.NewSettings()
	synth := fluidsynth2.NewSynth(s)
	i := synth.SFLoad("soundfont.sf2", false)

	player := fluidsynth2.NewPlayer(synth)
	player.Add("song.mid")

	// Example of how to play from memory
	// dat, err := ioutil.ReadFile("midifile.mid")
	// if err != nil {
	// 	panic(err)
	// }

	// player.AddMem(dat)

	// Easy way to set audio backend
	//s.SetString("audio.driver", "coreaudio")

	fluidsynth2.NewAudioDriver(s, synth)

	player.Play()
	player.Join()

}
