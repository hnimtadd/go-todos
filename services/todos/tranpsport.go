package todos

type TodoTransport interface {
	GetAll()
	GetUserTodos()
	AddTodo()
}
