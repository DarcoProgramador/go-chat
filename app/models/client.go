package models

import (
	"sync"
)

// Add more data to this type if needed
type Client struct {
	IsClosing bool
	Mu        sync.Mutex
}
