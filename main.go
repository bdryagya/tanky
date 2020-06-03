package main

import (
	"fmt"

	"github.com/hive-repo/01-06-2020/tanky"
)

func main() {
	a := tanky.Tanky{}
	for {
		fmt.Println("what do you want to do?\nfill or drain ")
		var want string
		fmt.Scan(&want)
		if want == "fill" {
			a.Fill()
		} else if want == "drain" {
			a.Drain()
		} else {
			fmt.Println("invalid operation")
		}
	}
}
