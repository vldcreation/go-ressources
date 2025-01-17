package conditional

import "fmt"

/*
// Conditional statements in Go.
// ref: https://go.dev/doc/effective_go#if
*/

/*
// IF ELSE
*/

func IfElse() {
	// If
	start := 0
	end := 10
	if start < end {
		fmt.Printf("start: %d is less than end: %d\n", start, end)
	}

	// increement start ny 1
	start++
	for start <= end {
		if start < end {
			fmt.Printf("increement start: %d by 1 is less than end: %d\n", start, end)
		} else {
			fmt.Printf("start: %d is equal to end: %d\n", start, end)
		}
		start++
	}

}

/*
// FOR
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }
*/
