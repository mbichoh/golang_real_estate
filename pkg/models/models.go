package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Agent struct {
	ID        int
	Name      string
	Email     string
	Phone     string
	CreatedAt time.Time
}

type Estate struct {
	ID        int
	AgentID   int
	Name      string
	Address   string
	ShortDesc string
	LongDesc  string
	County    string
	Bedroom   int
	Washroom  int
	SpaceArea int
	Packing   int
	Price     bool
	CreatedAt time.Time
}
