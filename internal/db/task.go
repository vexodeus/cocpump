package db

type Task struct {
	ID     int
	Type   TaskType
	Status TaskStatus
	Data   any
}

type TaskType int

const (
	TASK_TYPE_NULL TaskType = iota
	TASK_TYPE_INSERT
	TASK_TYPE_DELETE
	TASK_TYPE_GET
	TASK_TYPE_UPDATE
)

type TaskStatus int

const (
	TASK_STATUS_NEW TaskStatus = iota
	TASK_STATUS_ERROR
	TASK_STATUS_DONE
)
