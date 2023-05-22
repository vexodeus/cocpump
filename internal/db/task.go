package db

type Task struct {
	ID   int
	Type int
	Data int
}
type TaskType int

const (
	TASK_TYPE_NULL TaskType = iota
	TASK_TYPE_INSERT
	TASK_TYPE_DELETE
	TASK_TYPE_GET
	TASK_TYPE_UPDATE
)
