//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hours, minutes, seconds int64
}
type TimeParseError struct {
	e     string
	input string
}

func main() {
	t, err := timeparse("23:61:23")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
}
func (t Time) String() string {
	return fmt.Sprintf("Hours : %v\nMinutes : %v\nSeconds %v", t.hours, t.minutes, t.seconds)
}
func (t *TimeParseError) Error() string {
	return "TimeParseError : " + t.e + "\ninput : " + t.input + "\n"
}
func timeparse(s string) (Time, error) {
	parts := strings.Split(s, ":")
	integerParts := []int64{}
	for _, v := range parts {
		val, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return Time{}, err
		}
		integerParts = append(integerParts, val)
	}
	if len(integerParts) > 3 {
		return Time{}, &TimeParseError{fmt.Sprintf("Too many time stamps\nwant (3)\n got (%v)\n", len(parts)), s}
	} else if len(integerParts) < 3 {
		return Time{}, &TimeParseError{fmt.Sprintf("Too few time stamps\nwant (3)\n got (%v)\n", len(parts)), s}
	} else if len(integerParts) == 3 {
		for i, v := range integerParts {
			f := ""
			switch i {
			case 0:
				f = "hours"
				if v > 23 || v < 0 {
					return Time{}, &TimeParseError{fmt.Sprintf("Invalid value for %v", f), s}
				}
			case 1:
				f = "minutes"
				if v > 60 || v < 0 {
					return Time{}, &TimeParseError{fmt.Sprintf("Invalid value for %v", f), s}
				}
			case 2:
				f = "seconds"
				if v > 60 || v < 0 {
					return Time{}, &TimeParseError{fmt.Sprintf("Invalid value for %v", f), s}
				}
			}
		}
		return Time{integerParts[0], integerParts[1], integerParts[2]}, nil
	} else {
		return Time{}, errors.New("Uknowkn Error")
	}
}
