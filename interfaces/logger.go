package interfaces

import "fmt"

type Logger struct{}

func (logger *Logger) Log(message string) error {
	fmt.Println(message)
	return nil
}
