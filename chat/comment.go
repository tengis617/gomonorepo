package chat

import "github.com/tengis617/gomonorepo/kanban"

type Comment struct {
	TaskID string
	Body   string
}

func NewComment(task *kanban.Task, body string) *Comment {
	return &Comment{TaskID: task.ID, Body: body}
}
