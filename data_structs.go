package main

import (
	"context"
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

func (s *store) PostUnmarshal(ctx context.Context, row, count int) error {

	loc, err := time.LoadLocation("Europe/Budapest")
	if err != nil {
		panic(err)
	}
	newTimeZone := s.Timing.In(loc)
	s.Timing = &newTimeZone

	return nil
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
