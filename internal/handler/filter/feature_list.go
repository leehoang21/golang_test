package filter

import (
	"base/internal/base/mgo/filter"
	"base/internal/data/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeatureListParams struct {
	filter.PaginationFilter
	FeatureName string           `json:"feature_name" form:"feature_name"`
	Status      enums.StatusType `json:"status" form:"status"`
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
	if f.Status.String() != "" || f.Status != enums.StatusType(2) {
		f.AddWhere("status", "status", f.Status)
	}
	return f.BasicFilter.GetWhere()
}
