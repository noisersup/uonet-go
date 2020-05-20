package utils

/*
	Some utilities like error logging
*/
import "fmt"

func ErrLog(err error) {
	fmt.Print(err.Error)
}
