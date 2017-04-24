package pool

//Интерфейс задания. В пуле вызывается Execute().
type Task interface {
	Execute()
}
