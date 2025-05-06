package filter

import (
	"base/internal/base/mgo/filter"
	"base/internal/data/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeatureListParams struct {
	filter.PaginationFilter
	FeatureName string `json:"feature_name" form:"feature_name"`
	Status      string `json:"status" form:"status"`
}

func NewFeatureListPrams() *FeatureListParams {
	return &FeatureListParams{
		PaginationFilter: *filter.NewPaginationFilter(),
	}
}

func (f *FeatureListParams) GetWhere() filter.Where {
	if f.FeatureName != "" {
		f.AddWhere("feature_name", "name", primitive.Regex{Pattern: f.FeatureName, Options: ""})
	}
	if f.Status != "" || enums.StringToStatusType(f.Status) != enums.StatustypeUndefined {
		f.AddWhere("status", "status", f.Status)
	}
	f.AddWhere("dtime", "dtime", 0)
	return f.BasicFilter.GetWhere()
}
