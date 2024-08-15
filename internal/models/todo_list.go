package models

type TodoList []TodoItem

type TodoListInterface interface {
	RemoveByIndex() TodoList
	AddItem() TodoList
	ToggleStatus(int) TodoList
	ToJsonString() string
}

func (list TodoList) RemoveByIndex(index int) (resultList TodoList) {
	for i, v := range list {
		if index != i+1 {
			resultList = append(resultList, v)
		}
	}

	return resultList
}

func (list TodoList) AddItem(item TodoItem) (resultList TodoList) {
	return append(list, item)
}

func (list TodoList) ToggleStatus(index int) (resultList TodoList) {
	for i, v := range list {
		if index == i+1 {
			v.ToggleStatus()
		}
		resultList = append(resultList, v)
	}

	return resultList
}

func (list TodoList) ToJsonString() (listString string) {
	for _, v := range list {
		listString += v.ToJsonString()
	}

	return listString
}
