package pkg

import "fmt"

type ContextKey int

const (
	TransactionKey ContextKey = iota
	UnscopedKey
)

type ErrorCodeName string

const (
	ErrUniqueConstraint ErrorCodeName = "ERR_UNIQUE_CONSTRAINT"
)

// Errors
var (
	ErrInvalidId error = fmt.Errorf("invalid id")
)
