package handler

import (
	"base/internal/handler/filter"
	"base/internal/service/feature"
	"base/internal/utils/web"

	"github.com/gin-gonic/gin"
)

type FeatureHandler struct {
	service feature.Service
	*web.JsonRender
	contextWith web.ContextWith
}

func NewFeatureHandler(
	service feature.Service,
	contextWith web.ContextWith,
) FeatureHandler {
	return FeatureHandler{
		service:     service,
		contextWith: contextWith,
	}
}

// CreateHandler
// @Tags         Feature
// @Summary      Create feature
// @Description  Create feature
// @ID           feature-create
// @Accept       json
// @Produce      json
// @Param        data       body      models.Feature  true   "feature data"
// @Success      200        {object}  models.Feature
// @Router       /api/v1/features [post]
func (gr FeatureHandler) CreateHandler(ctx *gin.Context) {
	var input feature.Input
	web.AssertNil(ctx.BindJSON(&input))
	f, err := gr.service.Create(ctx.Request.Context(), input)
	web.AssertNil(err)
	gr.SendData(ctx, f)
}

// UpdateHandler
// @Tags         Feature
// @Summary      Update Feature
// @Description  Update Feature
// @ID           feature-update
// @Accept       json
// @Produce      json
// @Param        id  		header    string      		 false  "feature id"
// @Param        data       body      models.Feature   true   "feature data"
// @Success      200        {object}  models.Feature
// @Router       /api/v1/features [put]
func (gr FeatureHandler) UpdateHandler(ctx *gin.Context) {
	var input feature.Input
	web.AssertNil(ctx.BindJSON(&input))
	id := ctx.Param("id")
	var f, err = gr.service.Update(ctx.Request.Context(), id, input)
	web.AssertNil(err)
	gr.SendData(ctx, f)
}

// DeleteHandler
// @Tags         Feature
// @Summary      delete Feature
// @Description  delete Feature
// @ID           feature-delete
// @Accept       json
// @Produce      json
// @Param        id  		in:query    string       false  "feature id"
// @Success      200        {object}  null
// @Router       /api/v1/features/:id [delete]
func (gr FeatureHandler) DeleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := gr.service.Delete(ctx.Request.Context(), id)
	web.AssertNil(err)
	gr.SendData(ctx, nil)
}

// GetListHandler
// @Tags         Feature
// @Summary      get list feature
// @Description  get list feature
// @ID           feature-get-list
// @Accept       json
// @Produce      json
// @Param        data       body      models.Feature  true   "Feature list data"
// @Success      200        {object}  []models.Feature
// @Router       /api/v1/features [get]
func (gr FeatureHandler) GetListHandler(ctx *gin.Context) {
	var fil = filter.NewFeatureListPrams()
	web.AssertNil(ctx.BindQuery(fil))
	var f, total, err = gr.service.List(ctx.Request.Context(), fil)
	web.AssertNil(err)
	gr.SendData(ctx, responseList{
		Data:  f,
		Total: total,
	})
}

// GetHandler
// @Tags         Feature
// @Summary      get feature
// @Description  get feature
// @ID          feature-get
// @Accept       json
// @Produce      json
// @Param        id  header    string       false  "feature id"
// @Param        id       in:query    string  true   "feature data"
// @Success      200        {object}  models.Feature
// @Router       /api/v1/features/:id [get]
func (gr FeatureHandler) GetHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var f, err = gr.service.GetByID(ctx.Request.Context(), id)
	web.AssertNil(err)
	gr.SendData(ctx, f)
}
