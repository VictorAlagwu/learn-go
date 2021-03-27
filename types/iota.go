package iota

import (
	"fmt"
)

func main () {
	const (
		monday = iota + 1
		tuesday 
		wednesday
		thursday 
	)

	const (
		EST = -(5 + iota)
		_
		MST 
		PST
	)

	fmt.Println(EST, MST, PST)

	fmt.Println(monday, tuesday,wednesday,thursday )
}
