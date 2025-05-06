package enums

type TaskStatusType int

const (
	TaskStatusTypePending TaskStatusType = iota
	TaskStatusTypeInProgress
	TaskStatusTypeCompleted
	TaskStatusTypeUndefined
)

var taskStatusType = map[string]TaskStatusType{
	TaskStatusTypePending.String():    TaskStatusTypePending,
	TaskStatusTypeInProgress.String(): TaskStatusTypeInProgress,
	TaskStatusTypeCompleted.String():  TaskStatusTypeCompleted,
	TaskStatusTypeUndefined.String():  TaskStatusTypeUndefined,
}

func (p TaskStatusType) String() string {
	return []string{
		"status_type_pending",
		"status_type_in_progress",
		"status_type_completed",
		"status_type_undefined",
	}[p]
}

func (p TaskStatusType) Val() string {
	return []string{
		"pending",
		"progress",
		"completed",
		"undefined",
	}[p]
}

func StringToTaskStatusType(s string) TaskStatusType {
	r, ok := taskStatusType[s]
	if ok {
		return r
	}
	return TaskStatusTypeUndefined
}
