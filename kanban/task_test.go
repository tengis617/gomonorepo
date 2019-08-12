package kanban

import (
	"testing"

	"github.com/tengis617/gomonorepo"
)

func TestNewTask(t *testing.T) {
	u := gomonorepo.NewUser("user1", "bob", 50)
	task := NewTask(u, "do some work", "task1")

	if task.OwnerID != u.ID {
		t.Errorf("expected %s, got %s", u.ID, task.OwnerID)
	}
}
