package filter

import (
	"base/internal/base/mgo/filter"
	"base/internal/data/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskListParams struct {
	filter.PaginationFilter
	TaskTitle string `json:"task_title" form:"task_title"`
	Status    string `json:"status" form:"status"`
}

func NewTaskListPrams() *TaskListParams {
	return &TaskListParams{
		PaginationFilter: *filter.NewPaginationFilter(),
	}
}

func (f *TaskListParams) GetWhere() filter.Where {
	if f.TaskTitle != "" {
		f.AddWhere("task_title", "title", primitive.Regex{Pattern: f.TaskTitle, Options: ""})
	}
	if f.Status != "" || enums.StringToTaskStatusType(f.Status) != enums.TaskStatusTypeUndefined {
		f.AddWhere("status", "status", f.Status)
	}
	f.AddWhere("dtime", "dtime", 0)
	return f.BasicFilter.GetWhere()
}
