package application

import (
	"context"
	"time"

	"github.com/MatsuoTakuro/ddd-article/domain"
	"github.com/MatsuoTakuro/ddd-article/repository"
)

type Task struct {
	taskRepo repository.Task
}

func (t *Task) Create(ctx context.Context, name string, dueDate time.Time, priority domain.Priority) error {
	task, err := domain.NewTask(name, dueDate, priority)
	if err != nil {
		return err
	}
	if err := t.taskRepo.Save(ctx, *task); err != nil {
		return err
	}
	return nil
}

func (t *Task) Postpone(ctx context.Context, taskID string) error {
	task, err := t.taskRepo.GetByID(ctx, taskID)
	if err != nil {
		return nil
	}
	postponeTask, err := task.Postpone()
	if err != nil {
		return err
	}
	if err := t.taskRepo.Save(ctx, *postponeTask); err != nil {
		return err
	}
	return nil
}
