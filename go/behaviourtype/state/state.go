package state

import "fmt"

type State interface {
	play()
	pause()
	stop()
}

// State 的基父类
type StateBase struct {
	Context Context
}

// changeState 切换上下文的状态
func (s *StateBase) changeState(state State) {
	s.Context.State = state
}

type PlayState struct {
	StateBase
}

func (playState *PlayState) play() {
	fmt.Println("已经在播放了")
}

func (playState *PlayState) pause() {
	fmt.Println("视频切换到暂停状态")
	playState.StateBase.changeState(&PauseState{})
}

func (playState *PlayState) stop() {
	fmt.Println("视频切换到停止状态")
	playState.StateBase.changeState(&StopState{})
}

type PauseState struct {
	StateBase
}

func (pauseState *PauseState) play() {
	fmt.Println("视频切换到播放状态")
	pauseState.StateBase.changeState(&PlayState{})
}

func (pauseState *PauseState) pause() {
	fmt.Println("视频已经暂停了")
}

func (pauseState *PauseState) stop() {
	fmt.Println("视频切换到停止状态")
	pauseState.StateBase.changeState(&StopState{})
}

type StopState struct {
	StateBase
}

func (stopState *StopState) play() {
	fmt.Println("视频切换到播放状态")
	stopState.StateBase.changeState(&PlayState{})
}

func (stopState *StopState) pause() {
	fmt.Println("视频切换到暂停状态")
	stopState.StateBase.changeState(&PauseState{})
}

func (paustopStateseState *StopState) stop() {
	fmt.Println("视频已经停止了")
}
