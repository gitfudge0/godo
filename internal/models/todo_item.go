package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type TodoItem struct {
	Created_at time.Time `json:"created_at"`
	Title      string    `json:"title"`
	Priority   string    `json:"priority"`
	Is_done    bool      `json:"is_done"`
	Id         int       `json:"id"`
}

type TodoItemInterface interface {
	ToString() string
	ToJsonString() string
	ToggleStatus()
}

func (item *TodoItem) toggleStatus() {
	item.Is_done = !item.Is_done
}

func (item *TodoItem) ToJsonString() string {
	tJson, _ := json.Marshal(item)
	return string(tJson) + "\n"
}

func (item *TodoItem) ToString() string {
	item_str := fmt.Sprintf("id: %d\ttitle: %s\tis_done: %t\tcreated_at: %s\n", item.Id, item.Title, item.Is_done, item.Created_at)
	return item_str
}
