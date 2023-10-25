package pattern

import "fmt"

/*
Паттерн Command относится к поведенческим паттернам уровня объекта.
Он позволяет представить запрос в виде объекта. Из этого следует, что команда - это объект.
Такие запросы, например, можно ставить в очередь, отменять или возобновлять.
*/

type Command interface {
	execute()
}

type Editor interface {
	Save()
	Close()
	Undo()
}

type VSCode struct{}

func (vs *VSCode) Save() {
	fmt.Println("Save")
}

func (vs *VSCode) Close() {
	fmt.Println("Close")
}

func (vs *VSCode) Undo() {
	fmt.Println("Undo")
}

type SaveCommand struct {
	editor Editor
}

func (this *SaveCommand) execute() {
	this.editor.Save()
}

type CloseCommand struct {
	editor Editor
}

func (this *CloseCommand) execute() {
	this.editor.Close()
}

type UndoCommand struct {
	editor Editor
}

func (this *UndoCommand) execute() {
	this.editor.Undo()
}

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}
