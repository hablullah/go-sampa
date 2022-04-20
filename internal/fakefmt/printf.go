package fakefmt

import "fmt"

func Printf(f string, args ...interface{}) {
	// fmt.Printf(f, args...)
}

func Errorf(f string, args ...interface{}) error {
	return fmt.Errorf(f, args...)
}
