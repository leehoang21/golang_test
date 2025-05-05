package handler

import (
	"base/internal/handler/filter"
	"base/internal/service/task"
	"base/internal/utils/web"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service task.Service
	*web.JsonRender
	contextWith web.ContextWith
}

func NewTaskHandler(
	service task.Service,
	contextWith web.ContextWith,
) TaskHandler {
	return TaskHandler{
		service:     service,
		contextWith: contextWith,
	}
}

// CreateHandler
// @Tags         Task
// @Summary      Create task
// @Description  Create task
// @ID           task-create
// @Accept       json
// @Produce      json
// @Param        data       body      models.Task  true   "task data"
// @Success      200        {object}  models.Task
// @Router       /api/v1/tasks [post]
func (es TaskHandler) CreateHandler(ctx *gin.Context) {
	var input task.Input
	web.AssertNil(ctx.BindJSON(&input))
	getUser, e := es.contextWith.GetUser(ctx)
	web.AssertNil(e)
	f, err := es.service.Create(ctx.Request.Context(), input, getUser.Email)
	web.AssertNil(err)
	es.SendData(ctx, f)
}

// UpdateHandler
// @Tags         Task
// @Summary      Update task
// @Description  Update task
// @ID           task-update
// @Accept       json
// @Produce      json
// @Param        id  		header    string      		 false  "task id"
// @Param        data       body      models.Task   true   "task data"
// @Success      200        {object}  models.Task
// @Router       /api/v1/tasks [put]
func (es TaskHandler) UpdateHandler(ctx *gin.Context) {
	var input task.Input
	web.AssertNil(ctx.BindJSON(&input))
	id := ctx.Param("id")
	var f, err = es.service.Update(ctx.Request.Context(), id, input)
	web.AssertNil(err)
	es.SendData(ctx, f)
}

// DeleteHandler
// @Tags         Task
// @Summary      delete task
// @Description  delete task
// @ID           task-delete
// @Accept       json
// @Produce      json
// @Param        id  		in:query    string       false  "task id"
// @Success      200        {object}  null
// @Router       /api/v1/tasks/:id [delete]
func (es TaskHandler) DeleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := es.service.Delete(ctx.Request.Context(), id)
	web.AssertNil(err)
	es.SendData(ctx, nil)
}

// GetListHandler
// @Tags         Task
// @Summary      get list task
// @Description  get list task
// @ID           task-get-list
// @Accept       json
// @Produce      json
// @Param        data       body      models.Task  true   "task list data"
// @Success      200        {object}  []models.Task
// @Router       /api/v1/tasks [get]
func (es TaskHandler) GetListHandler(ctx *gin.Context) {
	var fil = filter.NewTaskListPrams()
	web.AssertNil(ctx.BindQuery(fil))
	var f, total, err = es.service.List(ctx.Request.Context(), fil)
	web.AssertNil(err)
	es.SendData(ctx, responseList{
		Data:  f,
		Total: total,
	})
}

// GetHandler
// @Tags         Task
// @Summary      get task
// @Description  get task
// @ID          task-get
// @Accept       json
// @Produce      json
// @Param        id  header    string       false  "task id"
// @Param        id       in:query    string  true   "task data"
// @Success      200        {object}  models.Task
// @Router       /api/v1/tasks/:id [get]
func (es TaskHandler) GetHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var f, err = es.service.GetByID(ctx.Request.Context(), id)
	web.AssertNil(err)
	es.SendData(ctx, f)
}
