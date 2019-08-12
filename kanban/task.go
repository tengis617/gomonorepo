package kanban

import (
	"github.com/tengis617/gomonorepo"
)

type Task struct {
	ID      string
	OwnerID string
	Name    string
}

func NewTask(owner *gomonorepo.User, id, name string) *Task {
	return &Task{OwnerID: owner.ID, ID: id, Name: name}
}
