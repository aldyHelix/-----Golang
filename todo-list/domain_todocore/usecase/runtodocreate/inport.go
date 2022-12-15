package runtodocreate

import (
	"todo-list/domain_todocore/model/entity"
	"todo-list/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.TodoCreateRequest
}

type InportResponse struct {
	Todo *entity.Todo
}
