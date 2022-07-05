package pkg

import "fmt"

const (
	NotFound = 404
	Duplicate= 409
	Unknown = 503
	)


type InmemoryNotFound struct {
	Code  int
	Description  string
}

func (e InmemoryNotFound) Error() string {
	return fmt.Sprintf("%s", e.Description)
}


type InmemoryDuplicationKey struct {
	Code int
	Description  string
}

func (e InmemoryDuplicationKey) Error() string {
	return fmt.Sprintf("%s", e.Description)
}