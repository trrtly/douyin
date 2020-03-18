package microapp

import (
	"github.com/trrtly/douyin/context"
)

// Microapp struct
type Microapp struct {
	*context.Context
}

// New init
func New(c *context.Context) *Microapp {
	return &Microapp{c}
}
