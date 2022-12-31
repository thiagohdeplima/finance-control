package parser

import (
	"time"
)

// Entry represents a financial
// entry extracted from the account
type Entry struct {
	ID     string
	Date   time.Time
	Desc   string
	Amount float32
}
