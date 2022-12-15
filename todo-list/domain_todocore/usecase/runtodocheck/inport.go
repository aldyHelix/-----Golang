package runtodocheck

import (
	"todo-list/domain_todocore/model/entity"
	"todo-list/domain_todocore/model/vo"
	"todo-list/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	TodoID vo.TodoID
}

type InportResponse struct {
	Todo *entity.Todo
}
