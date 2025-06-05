package usecaseinterface

import "Kuraz-Tech-Challenge/domain/entity"

type TaskUsecaseInterface interface {
	CreateTask(task *entity.Task) (*entity.Task, error)
	GetTasks() ([]*entity.Task, error)
	UpdateTask(task *entity.Task) (*entity.Task, error)
	DeleteTask(id uint) error
}