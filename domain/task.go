package domain

import (
	"fmt"
	"time"
)

const POSTPONE_MAX_COUNT = 5

type Task struct {
	id            string
	status        Status
	name          string
	dueDate       time.Time
	priority      Priority
	postponeCount int64
}

func NewTask(name string, dueDate time.Time, priority Priority) (*Task, error) {
	if name == "" || dueDate.IsZero() {
		return nil, fmt.Errorf("DomainError: %s", "No required fields filled in")
	}

	return &Task{
		id:            "",
		status:        Doing,
		name:          name,
		dueDate:       dueDate,
		priority:      priority,
		postponeCount: 0,
	}, nil
}

func (t *Task) Postpone() (*Task, error) {
	if !t.CanPostpone() {
		return nil, fmt.Errorf("DomainError: %s", "no more postponement")
	}
	t.dueDate.Add(24 * time.Hour)
	t.postponeCount++
	return t, nil
}

func (t *Task) CanPostpone() bool {
	return t.postponeCount < POSTPONE_MAX_COUNT
}

func (t *Task) Done() {
	t.status = Done
}

func (t *Task) IsDoing() bool {
	return t.status == Doing
}

func (t *Task) GetID() string {
	return t.id
}
func (t *Task) GetName() string {
	return t.name
}
func (t *Task) GetDueDate() time.Time {
	return t.dueDate
}
func (t *Task) GetPriority() Priority {
	return t.priority
}
