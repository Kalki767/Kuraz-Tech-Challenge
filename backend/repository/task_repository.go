package repository

import (
	repositoryinterface "Kuraz-Tech-Challenge/domain/contracts/repository_interface"
	"Kuraz-Tech-Challenge/domain/entity"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) repositoryinterface.TaskRepositoryInterface {
	return &TaskRepository{db: db}
}

func (taskrepo *TaskRepository) CreateTask(task *entity.Task) (*entity.Task, error) {
	err := taskrepo.db.Create(task).Error
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (taskrepo *TaskRepository) GetTasks() ([]*entity.Task, error) {
	var tasks []*entity.Task
	err := taskrepo.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (taskrepo *TaskRepository) UpdateTask(task *entity.Task) (*entity.Task, error){
	err := taskrepo.db.Model(&entity.Task{}).Where("id = ?", task.ID).Updates(task).Error

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (taskrepo *TaskRepository) DeleteTask(id uint) error{
	err := taskrepo.db.Delete(&entity.Task{}, id).Error
	if err != nil {
		return err
	}
	return nil
}