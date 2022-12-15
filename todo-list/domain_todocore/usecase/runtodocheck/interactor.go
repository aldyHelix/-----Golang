package runtodocheck

import (
	"context"
	"fmt"
)

type runTodoCheckInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runTodoCheckInteractor{
		outport: outputPort,
	}
}

func (r *runTodoCheckInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...

	//ambil data dengan id tertentu
	todoObj, err := r.outport.FindOneTodoByID(ctx, req.TodoID)
	if err != nil {
		return nil, err
	}
	if todoObj == nil {
		return nil, fmt.Errorf("object not found")
	}
	//ubah state checked = true
	err = todoObj.Check()
	if err != nil {
		return nil, err
	}
	//simpan kembali ke persistent storage
	r.outport.SaveTodo(ctx, todoObj)
	if err != nil {
		return nil, err
	}

	res.Todo = todoObj

	return res, nil
}
