package main

import (
	"time"
)

type store struct {
	ID        int64      `dbq:"id"`
	Product   string     `dbq:"product"`
	Price     float64    `dbq:"price"`
	Quantity  int64      `dbq:"quantity"`
	Available *bool      `dbq:"available"`
	Timing    *time.Time `dbq:"timing"`
}

type store2 struct {
	ID        int64   `dbq:"id"`
	Product   string  `dbq:"product"`
	Price     float64 `dbq:"price"`
	Quantity  int64   `dbq:"quantity"`
	Available *bool   `dbq:"available"`
	Timing    *string `dbq:"timing"`
}

type bnchmark struct {
	ID        int64      `dbq:"id"`
	Name      string     `dbq:"name"`
	Timestamp *time.Time `dbq:"timestamp"`
}

type bnchmark2 struct {
	ID        int64   `dbq:"id"`
	Name      string  `dbq:"name"`
	Timestamp *string `dbq:"timestamp"`
}
