package handlers

import (
	usecaseinterface "Kuraz-Tech-Challenge/domain/contracts/usecase_interface"
	"Kuraz-Tech-Challenge/domain/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	usecase usecaseinterface.TaskUsecaseInterface
}

func NewTaskHanlder(usecase usecaseinterface.TaskUsecaseInterface) *TaskHandler {
	return &TaskHandler{usecase: usecase}
}

func (taskhandler *TaskHandler) CreateTask(ctx *gin.Context) {
	var task entity.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task format", "error": err})
		return
	}

	created, err := taskhandler.usecase.CreateTask(&task)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Something Went wrong", "error": err})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Successfully Created", "data": created})
}

func (taskhandler *TaskHandler) GetTasks(ctx *gin.Context) {
	
	tasks, err := taskhandler.usecase.GetTasks()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Something Went wrong", "error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully Retrieved", "data": tasks})
}

func (taskhandler *TaskHandler) UpdateTask(ctx *gin.Context) {
	var task entity.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task format", "error": err})
		return
	}

	updated, err := taskhandler.usecase.UpdateTask(&task)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Something Went wrong", "error": err})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Successfully Created", "data": updated})
}

func (taskhandler *TaskHandler) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	if id != "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Id is not provided"})
		return
	}

	task_id, err := strconv.Atoi(id)
	if err != nil || task_id <= 0{
		ctx.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"Message": "Invalid id format", "error": err})
	}

	err = taskhandler.usecase.DeleteTask(uint(task_id))
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Something Went wrong", "error": err})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Successfully Deleted"})
}