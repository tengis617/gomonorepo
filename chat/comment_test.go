package chat

import (
	"testing"

	"github.com/tengis617/gomonorepo"
	"github.com/tengis617/gomonorepo/kanban"
)

func TestNewComment(t *testing.T) {
	u := gomonorepo.NewUser("user1", "bob", 50)
	task := kanban.NewTask(u, "do some stuff", "task1")
	c := NewComment(task, "should i ?")

	if c.Body != "should i ?" {
		t.Errorf("expected %s, got %s", "should i ?", c.Body)
	}
}
