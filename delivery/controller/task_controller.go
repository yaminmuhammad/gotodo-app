package controller

import (
	"gotodo-app/config"
	"gotodo-app/model"
	"gotodo-app/shared/common"
	"gotodo-app/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUC usecase.TaskUseCase
	rg     *gin.RouterGroup
}

func (t *TaskController) createHandler(ctx *gin.Context) {
	var payload model.Task
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	task, err := t.taskUC.RegisterNewTask(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, task, "Created")
}

func (t *TaskController) listHandler(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))
	tasks, paging, err := t.taskUC.FindAllTask(page, size)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	var response []interface{}
	for _, v := range tasks {
		response = append(response, v)
	}
	common.SendPagedResponse(ctx, response, paging, "Ok")
}

func (t *TaskController) getByAuthorHandler(ctx *gin.Context) {
	author := ctx.Param("author")
	tasks, err := t.taskUC.FindTaskByAuthor(author)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, "task with author ID "+author+" not found")
		return
	}
	common.SendSingleResponse(ctx, tasks, "Ok")
}

func (t *TaskController) Route() {
	t.rg.POST(config.TaskPost, t.createHandler)
	t.rg.GET(config.TaskGetList, t.listHandler)
	t.rg.GET(config.TaskGetByAuthor, t.getByAuthorHandler)
}

func NewTaskController(taskUC usecase.TaskUseCase, rg *gin.RouterGroup) *TaskController {
	return &TaskController{
		taskUC: taskUC,
		rg:     rg,
	}
}
