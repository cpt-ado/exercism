// Package leap checks whether a given uear is leap or not
package leap

// IsLeapYear returns true of year is leap, false otherwise
func IsLeapYear(year int) bool {
	return ((year%4 == 0) && (year%100 != 0)) || (year%400 == 0)
}
