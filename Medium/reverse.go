package main

import "fmt"
import "strconv"
import "strings"

func reverse(x int) int {
	var xrev []string
	var negative bool
	if x < 0 {
		negative = true
		x = -x
	} else {
		negative = false
	}
	xstring := strconv.Itoa(x)
	for i := len(xstring) - 1; i >= 0; i-- {
		xrev = append(xrev, string(xstring[i]))
	}
	strxrev := strings.Join(xrev, "")
	xrevint, err := strconv.Atoi(strxrev)
	if err != nil {
		panic(err)
	}
	if negative {
		xrevint = -xrevint
	}
    if xrevint < -2147483648 || xrevint > 2147483648 {
        return 0
    }
	return xrevint
}
