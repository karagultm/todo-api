// todo.go
package main

type Todo struct {
	Id      int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Content string `gorm:"not null" json:"content" `
	Done    bool   `json:"done"`
}
