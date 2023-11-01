package main

import "fmt"

func chk(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
