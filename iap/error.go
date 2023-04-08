package iap

import "fmt"

type ConnectionError struct {
	Err string
}

func (e *ConnectionError) Error() string {
	return fmt.Sprintf("connection error: %v", e.Err)
}

type ProtocolError struct {
	Err string
}

func (e *ProtocolError) Error() string {
	return fmt.Sprintf("protocol error: %v", e.Err)
}
