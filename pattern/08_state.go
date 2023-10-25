package pattern

import (
	"errors"
)

/*
Паттерн State относится к поведенческим паттернам уровня объекта.
Он позволяет объекту изменять свое поведение в зависимости от внутреннего состояния и является объектно-ориентированной реализацией конечного автомата.
Поведение объекта изменяется настолько, что создается впечатление, будто изменился класс объекта.
*/

type State interface {
	doTask() error
}

type Worker struct {
	resting State
	busy    State
	free    State

	currentState State
}

func (w *Worker) doTask() error {
	return w.currentState.doTask()
}

func (w *Worker) SetState(state State) {
	w.currentState = state
}

type RestingState struct {
	w *Worker
}

func (this *RestingState) doTask() error {
	return errors.New("Not now")
}

type BusyState struct {
	w *Worker
}

func (this *BusyState) doTask() error {
	return errors.New("Already busy")
}

type FreeState struct {
	w *Worker
}

func (this *FreeState) doTask() error {
	this.w.SetState(this.w.busy)
	return nil
}
