package getalltodo

import "todo-list/domain_todocore/model/repository"

type Outport interface {
	repository.FindAllTodoRepo
}
