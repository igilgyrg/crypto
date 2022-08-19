package error

type InternalCode string

const (
	ItemNotFound InternalCode = "I404"
	SystemError               = "S500"
)
