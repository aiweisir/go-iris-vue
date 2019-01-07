package mappers

const (
	// ReadOnlyMode will RLock(read) the data .
	ReadOnlyMode = iota
	// ReadWriteMode will Lock(read/write) the data.
	ReadWriteMode
)

// Query represents the visitor and action queries.
type Query func(interface{}) bool

type Mapper interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)
}
