package models

type TodoList []TodoItem

type TodoListInterface interface {
	removeByIndex() TodoList
	addItem() TodoList
}

func (list TodoList) removeByIndex(index int) (resultList TodoList) {
	for i, v := range list {
		if index != i+1 {
			resultList = append(resultList, v)
		}
	}

	return resultList
}

func (list TodoList) addItem(item TodoItem) (resultList TodoList) {
	return append(list, item)
}
