// Code generated by goderive DO NOT EDIT.

package compare

import (
	"strings"
)

func deriveCompare(this, that *MyStruct) int {
	if this == nil {
		if that == nil {
			return 0
		}
		return -1
	}
	if that == nil {
		return 1
	}
	if c := deriveCompare_(this.Int64, that.Int64); c != 0 {
		return c
	}
	if c := deriveCompare_1(this.StringPtr, that.StringPtr); c != 0 {
		return c
	}
	return 0
}

func deriveCompare_(this, that int64) int {
	if this != that {
		if this < that {
			return -1
		} else {
			return 1
		}
	}
	return 0
}

func deriveCompare_1(this, that *string) int {
	if this == nil {
		if that == nil {
			return 0
		}
		return -1
	}
	if that == nil {
		return 1
	}
	return deriveCompare_s(*this, *that)
}

func deriveCompare_s(this, that string) int {
	return strings.Compare(this, that)
}