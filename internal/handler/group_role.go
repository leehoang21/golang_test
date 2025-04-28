package handler

import (
	"base/internal/handler/filter"
	"base/internal/service/group_role"
	"base/internal/utils/web"

	"github.com/gin-gonic/gin"
)

type GroupRoleHandler struct {
	service group_role.Service
	*web.JsonRender
	contextWith web.ContextWith
}

func NewGroupRoleHandler(
	service group_role.Service,
	contextWith web.ContextWith,
) GroupRoleHandler {
	return GroupRoleHandler{
		service:     service,
		contextWith: contextWith,
	}
}

// CreateHandler
// @Tags         GroupRole
// @Summary      Create group role
// @Description  Create group role
// @ID           group-role-create
// @Accept       json
// @Produce      json
// @Param        data       body      models.GroupRole  true   "group role data"
// @Success      200        {object}  models.GroupRole
// @Router       /api/v1/group-roles [post]
func (gr GroupRoleHandler) CreateHandler(ctx *gin.Context) {
	var input group_role.GroupRoleInput
	web.AssertNil(ctx.BindJSON(&input))
	var gRole, err = gr.service.Create(ctx.Request.Context(), input)
	web.AssertNil(err)
	gr.SendData(ctx, gRole)
}

// UpdateHandler
// @Tags         GroupRole
// @Summary      Update GroupRole
// @Description  Update GroupRole
// @ID           group-role-update
// @Accept       json
// @Produce      json
// @Param        id  		header    string      		 false  "group role id"
// @Param        data       body      models.GroupRole   true   "group role data"
// @Success      200        {object}  models.GroupRole
// @Router       /api/v1/group-roles [put]
func (gr GroupRoleHandler) UpdateHandler(ctx *gin.Context) {
	var input group_role.GroupRoleInput
	web.AssertNil(ctx.BindJSON(&input))
	id := ctx.Param("id")
	var gRole, err = gr.service.Update(ctx.Request.Context(), id, input)
	web.AssertNil(err)
	gr.SendData(ctx, gRole)
}

// DeleteHandler
// @Tags         GroupRole
// @Summary      delete GroupRole
// @Description  delete GroupRole
// @ID           group-role-delete
// @Accept       json
// @Produce      json
// @Param        id  		in:query    string       false  "group role id"
// @Success      200        {object}  null
// @Router       /api/v1/GroupRoles/:id [delete]
func (gr GroupRoleHandler) DeleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := gr.service.Delete(ctx.Request.Context(), id)
	web.AssertNil(err)
	gr.SendData(ctx, nil)
}

// GetListHandler
// @Tags         GroupRole
// @Summary      get list group role
// @Description  get list group role
// @ID           group-role-get-list
// @Accept       json
// @Produce      json
// @Param        data       body      models.GroupRole  true   "GroupRole data"
// @Success      200        {object}  []models.GroupRole
// @Router       /api/v1/group-roles [get]
func (gr GroupRoleHandler) GetListHandler(ctx *gin.Context) {
	var f = filter.NewGroupRoleListPrams()
	web.AssertNil(ctx.BindQuery(f))
	var gRole, total, err = gr.service.List(ctx.Request.Context(), f)
	web.AssertNil(err)
	gr.SendData(ctx, responseList{
		Data:  gRole,
		Total: total,
	})
}

// GetHandler
// @Tags         GroupRole
// @Summary      get  group role
// @Description  get  group role
// @ID           group-role-get
// @Accept       json
// @Produce      json
// @Param        id  header    string       false  "group role id"
// @Param        id       in:query    string  true   "group role data"
// @Success      200        {object}  models.GroupRole
// @Router       /api/v1/group-roles/:id [get]
func (gr GroupRoleHandler) GetHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var gRole, err = gr.service.GetByID(ctx.Request.Context(), id)
	web.AssertNil(err)
	gr.SendData(ctx, gRole)
}
