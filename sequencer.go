package fluidsynth2

// #cgo pkg-config: fluidsynth
// #include <fluidsynth.h>
// #include <stdlib.h>
/*


 */
import "C"

type Sequencer struct {
	ptr       *C.fluid_sequencer_t
	synthPtr  C.fluid_seq_id_t
	clientPtr C.fluid_seq_id_t
}

func NewSequencer() *Sequencer {
	return &Sequencer{ptr: C.new_fluid_sequencer2(0)}
}

// TODO unregister
// see https://www.fluidsynth.org/api/group__sequencer.html#gaae466e7d4341c9e17bad52c636f1a46d
func (s *Sequencer) RegisterSynth(synth Synth) {
	s.synthPtr = C.fluid_sequencer_register_fluidsynth(s.ptr, synth.ptr)
	// s.clientPtr = C.fluid_sequencer_register_client(s.ptr, C.CString("sequencer"), C.closure(C.CallMyFunction), C.NULL)
}
