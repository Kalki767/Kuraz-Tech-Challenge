package usecase

import (
	repositoryinterface "Kuraz-Tech-Challenge/domain/contracts/repository_interface"
	usecaseinterface "Kuraz-Tech-Challenge/domain/contracts/usecase_interface"
	"Kuraz-Tech-Challenge/domain/entity"
)

type TaskUsecase struct {
	repo repositoryinterface.TaskRepositoryInterface
}

func NewTaskUsecase(repo repositoryinterface.TaskRepositoryInterface) usecaseinterface.TaskUsecaseInterface{
	return &TaskUsecase{repo: repo}
}

func (usecase *TaskUsecase) CreateTask(task *entity.Task) (*entity.Task, error) {
	return usecase.repo.CreateTask(task)
}
func (usecase *TaskUsecase) GetTasks() ([]*entity.Task, error) {
	return usecase.repo.GetTasks()
}
func (usecase *TaskUsecase) UpdateTask(task *entity.Task) (*entity.Task, error) {
	return usecase.repo.UpdateTask(task)
}
func (usecase *TaskUsecase) DeleteTask(id uint) error {
	return usecase.repo.DeleteTask(id)
}