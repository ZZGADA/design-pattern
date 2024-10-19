package main

import "state"

func main() {
	context := state.Context{}
	playState := state.PlayState{state.StateBase{
		Context: context,
	}}
	context.State = &playState

	context.Play()
	context.Play()
	context.Pause()

	context.Pause()
	context.Pause()
	context.Stop()

	context.Stop()
	context.Stop()

}
