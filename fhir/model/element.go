package model

import "time"

type URI string

type Element struct {
	ID        string
	Extension Extension
}

type Extension[T any] struct {
	URL   URI
	Value T
}

type Meta struct {
	VersionID   string
	LastUpdated time.Time
	Source      URI
}