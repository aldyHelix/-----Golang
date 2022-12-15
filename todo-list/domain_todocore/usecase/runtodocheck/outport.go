package runtodocheck

import "todo-list/domain_todocore/model/repository"

type Outport interface {
	repository.FindOneTodoByIDRepo
	repository.SaveTodoRepo
}
