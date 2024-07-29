package models

import (
	"encoding/json"
	"fmt"
	"time"
)

func (t TodoItem) String() string {
	item_str := fmt.Sprintf("id: %d\ttitle: %s\tis_done: %t\tcreated_at: %s\n", t.Id, t.Title, t.Is_done, t.Created_at)
	return item_str
}

func (t TodoItem) JsonString() string {
	tJson, _ := json.Marshal(t)
	return string(tJson) + "\n"
}

type TodoItem struct {
	Created_at time.Time `json:"created_at"`
	Title      string    `json:"title"`
	Priority   string    `json:"priority"`
	Is_done    bool      `json:"is_done"`
	Id         int       `json:"id"`
}
