package filter

import (
	"eclectric/internal/base/mgo/filter"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserListParams struct {
	filter.PaginationFilter
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func NewUserListPrams() *UserListParams {
	return &UserListParams{
		PaginationFilter: *filter.NewPaginationFilter(),
	}
}

func (f *UserListParams) GetWhere() filter.Where {
	if f.Name != "" {
		f.AddWhere("name", "fullname", primitive.Regex{Pattern: f.Name, Options: ""})
	}
	if f.Email != "" {
		f.AddWhere("email", "email", primitive.Regex{Pattern: f.Email, Options: ""})
	}
	if f.Phone != "" {
		f.AddWhere("phone", "phone", primitive.Regex{Pattern: f.Phone, Options: ""})
	}
	return f.BasicFilter.GetWhere()
}
