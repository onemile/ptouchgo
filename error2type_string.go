// Code generated by "stringer -linecomment -type Error2Type"; DO NOT EDIT.

package ptouchgo

import "strconv"

const (
	_Error2Type_name_0 = "Invalid media"
	_Error2Type_name_1 = "Cover open"
	_Error2Type_name_2 = "Too hot"
)

func (i Error2Type) String() string {
	switch {
	case i == 1:
		return _Error2Type_name_0
	case i == 16:
		return _Error2Type_name_1
	case i == 32:
		return _Error2Type_name_2
	default:
		return "Error2Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
