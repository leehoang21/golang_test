package filter

import (
	"base/internal/base/mgo/filter"
	"base/internal/data/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GroupRoleListParams struct {
	filter.PaginationFilter
	Name   string           `json:"name" form:"name"`
	Status enums.StatusType `json:"status" form:"status"`
}

func NewGroupRoleListPrams() *GroupRoleListParams {
	return &GroupRoleListParams{
		PaginationFilter: *filter.NewPaginationFilter(),
	}
}

func (f *GroupRoleListParams) GetWhere() filter.Where {
	if f.Name != "" {
		f.AddWhere("name", "name", primitive.Regex{Pattern: f.Name, Options: ""})

	}
	if f.Status.String() != "" || f.Status != enums.StatusType(2) {
		f.AddWhere("status", "status", f.Status)
	}
	return f.BasicFilter.GetWhere()
}
