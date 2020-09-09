package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with the given error message.
		return -1, errors.New("can't work with 42")
	}
	// A nil value in the error position indicates that there was no error.

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		//In this case we use &argError syntax to build a new struct, supplying values for the two fields arg and prob.

		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil { //if there is error
			fmt.Println("f1 failed:", e)
		} else { // elde there is no error
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} { // working with custom error
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42) // above things can be written in this way as well
	ae, notOk := e.(*argError)
	if notOk {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
