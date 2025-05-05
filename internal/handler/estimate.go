package handler

import (
	"base/internal/service/estimate"
	"base/internal/utils/web"

	"github.com/gin-gonic/gin"
)

type EstimateHandler struct {
	service estimate.Service
	*web.JsonRender
	contextWith web.ContextWith
}

func NewEstimateHandler(
	service estimate.Service,
	contextWith web.ContextWith,
) EstimateHandler {
	return EstimateHandler{
		service:     service,
		contextWith: contextWith,
	}
}

// CreateHandler
// @Tags         Estimate
// @Summary      Create estimate
// @Description  Create estimate
// @ID           estimate-create
// @Accept       json
// @Produce      json
// @Param        data       body      models.Estimate  true   "estimate data"
// @Success      200        {object}  models.Estimate
// @Router       /api/v1/estimates [post]
func (es EstimateHandler) CreateHandler(ctx *gin.Context) {
	var input estimate.Input
	web.AssertNil(ctx.BindJSON(&input))
	getUser, e := es.contextWith.GetUser(ctx)
	web.AssertNil(e)
	f, err := es.service.Create(ctx.Request.Context(), input, getUser.Email)
	web.AssertNil(err)
	es.SendData(ctx, f)
}

// UpdateHandler
// @Tags         Estimate
// @Summary      Update estimate
// @Description  Update estimate
// @ID           estimate-update
// @Accept       json
// @Produce      json
// @Param        id  		header    string      		 false  "estimate id"
// @Param        data       body      models.Estimate   true   "estimate data"
// @Success      200        {object}  models.Estimate
// @Router       /api/v1/estimates [put]
func (es EstimateHandler) UpdateHandler(ctx *gin.Context) {
	var input estimate.Input
	web.AssertNil(ctx.BindJSON(&input))
	id := ctx.Param("id")
	var f, err = es.service.Update(ctx.Request.Context(), id, input)
	web.AssertNil(err)
	es.SendData(ctx, f)
}

// DeleteHandler
// @Tags         Estimate
// @Summary      delete estimate
// @Description  delete estimate
// @ID           estimate-delete
// @Accept       json
// @Produce      json
// @Param        id  		in:query    string       false  "estimate id"
// @Success      200        {object}  null
// @Router       /api/v1/estimates/:id [delete]
func (es EstimateHandler) DeleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := es.service.Delete(ctx.Request.Context(), id)
	web.AssertNil(err)
	es.SendData(ctx, nil)
}

// GetHandler
// @Tags         Estimate
// @Summary      get estimate
// @Description  get estimate
// @ID          estimate-get
// @Accept       json
// @Produce      json
// @Param        id  header    string       false  "estimate id"
// @Param        id       in:query    string  true   "estimate data"
// @Success      200        {object}  models.Estimate
// @Router       /api/v1/estimates/:id [get]
func (es EstimateHandler) GetHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var f, err = es.service.GetByID(ctx.Request.Context(), id)
	web.AssertNil(err)
	es.SendData(ctx, f)
}
