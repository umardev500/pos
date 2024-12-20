package pkg

type ContextKey int

const (
	TransactionKey ContextKey = iota
	UnscopedKey
)

type ErrorCodeName string

const (
	ErrUniqueConstraint ErrorCodeName = "ERR_UNIQUE_CONSTRAINT"
)
