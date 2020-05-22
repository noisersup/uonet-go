package utils

/*
	Some utilities like error logging
*/
import "fmt"

func ErrLog(err error) {
	fmt.Println("ERROR=======================")
	fmt.Println(err.Error())
	fmt.Println("END=========================")
}
