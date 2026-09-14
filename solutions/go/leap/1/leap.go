// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// package leap Show current year is leap or not
package leap

// IsLeapYear shows what kind of year is input
func IsLeapYear(year int) bool {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	return (year % 400 == 0) || (year % 4 == 0 && year % 100 != 0)
}
