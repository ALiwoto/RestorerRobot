package stringUtils

import "strconv"

// MakeSureNum will make sure that when you convert `i`
// to string, its length be the exact same as `count`.
// it will append 0 to the left side of the number to do so.
// for example:
// MakeSureNum(5, 8) will return "00000005"
func MakeSureNum(i, count int) string {
	return MakeSureNumCustom(i, count, "0")
}

// MakeSureNumCustom will make sure that when you convert `i`
// to string, its length be the exact same as `count`.
// it will append 0 to the left side of the number to do so.
// for example:
// MakeSureNum(5, 8) will return "00000005"
func MakeSureNumCustom(i, count int, holder string) string {
	s := strconv.Itoa(i)
	final := count - len(s)
	for ; final > 0; final-- {
		s = holder + s
	}

	return s
}
