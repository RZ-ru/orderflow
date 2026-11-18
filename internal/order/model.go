package order

import (
	"slices"
	"time"
)

type Status string

const (
	StatusCreated   Status = "created"
	StatusPaid      Status = "paid"
	StatusShipped   Status = "shipped"
	StatusCancelled Status = "cancelled"
)

type Product struct {
	ID         int64
	Name       string
	PriceCents int64
	Stock      int
}

type Item struct {
	ProductID  int64
	Qty        int
	PriceCents int64 //цена за штуку на момент заказа
}

type Order struct {
	ID         int64
	UserID     int64
	Items      []Item
	Status     Status
	TotalCents int64
	CreatedAt  time.Time
}

var transitions = map[Status][]Status{
	StatusCreated: {StatusPaid, StatusCancelled},
	StatusPaid:    {StatusShipped, StatusCancelled},
}

func CanTransition(from, to Status) bool {
	//Для отсутствующего ключа map вернет nil-слайс
	return slices.Contains(transitions[from], to)
}

func ValidStatus(s Status) bool {
	switch s {
	case StatusCreated, StatusPaid, StatusCancelled, StatusShipped:
		return true
	}
	return false
}
