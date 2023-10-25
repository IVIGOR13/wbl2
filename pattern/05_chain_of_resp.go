package pattern

import "fmt"

/*
Паттерн Chain относится к поведенческим паттернам уровня объекта.
Паттерн Chain позволяет избежать привязки объекта-отправителя запроса к объекту-получателю запроса, при этом давая шанс обработать этот запрос нескольким объектам.
Получатели связываются в цепочку, и запрос передается по цепочке, пока не будет обработан каким-то объектом.
По сути это цепочка обработчиков, которые по очереди получают запрос, а затем решают, обрабатывать его или нет.
Если запрос не обработан, то он передается дальше по цепочке. Если же он обработан, то паттерн сам решает передавать его дальше или нет.
Если запрос не обработан ни одним обработчиком, то он просто теряется.
*/

type Middleware interface {
	execute(*Handler)
	setNext(Middleware)
}

type Handler struct {
	username string
	path     string
	args     []string
}

type AuthMiddleware struct {
	next Middleware
}

func (this AuthMiddleware) execute(handler *Handler) {
	if handler.username != "" {
		fmt.Println("auth is ok")
		this.next.execute(handler)
		return
	}
	fmt.Println("auth is not ok")
}

func (this AuthMiddleware) setNext(next Middleware) {
	this.next = next
}

type AdminMiddleware struct {
	next Middleware
}

func (this AdminMiddleware) execute(handler *Handler) {
	if handler.username == "admin" {
		fmt.Println("user is admin")
		this.next.execute(handler)
		return
	}
	fmt.Println("user is not admin")
}

func (this AdminMiddleware) setNext(next Middleware) {
	this.next = next
}

func main() {

	auth := AuthMiddleware{}
	admin := AdminMiddleware{}

	auth.setNext(admin)

	auth.execute(&Handler{
		username: "user",
		path:     "/auth",
		args:     []string{},
	})
}
