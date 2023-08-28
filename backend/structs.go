package main

import (
	"time"
)

type List struct {
	ListID       int       `json:"listID"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Status       int       `json:"status"`
	CreateDate   time.Time `json:"created"`
	CompleteDate time.Time `json:"completed"`
}

type Item struct {
	ItemID      int    `json:"itemID"`
	ListID      int    `json:"listID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}
