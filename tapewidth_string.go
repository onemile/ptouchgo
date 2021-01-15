// Code generated by "stringer -linecomment -type TapeWidth"; DO NOT EDIT.

package ptouchgo

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[tapeWidthNone-0]
	_ = x[tapeWidth3_5-4]
	_ = x[tapeWidth6-6]
	_ = x[tapeWidth9-9]
	_ = x[tapeWidth12-12]
	_ = x[tapeWidth18-18]
	_ = x[tapeWidth24-24]
	_ = x[tapeWidth62-62]
}

const (
	_TapeWidth_name_0 = "No tape"
	_TapeWidth_name_1 = "3.5mm"
	_TapeWidth_name_2 = "6mm"
	_TapeWidth_name_3 = "9mm"
	_TapeWidth_name_4 = "12mm"
	_TapeWidth_name_5 = "18mm"
	_TapeWidth_name_6 = "24mm"
	_TapeWidth_name_7 = "62mm"
)

func (i TapeWidth) String() string {
	switch {
	case i == 0:
		return _TapeWidth_name_0
	case i == 4:
		return _TapeWidth_name_1
	case i == 6:
		return _TapeWidth_name_2
	case i == 9:
		return _TapeWidth_name_3
	case i == 12:
		return _TapeWidth_name_4
	case i == 18:
		return _TapeWidth_name_5
	case i == 24:
		return _TapeWidth_name_6
	case i == 62:
		return _TapeWidth_name_7
	default:
		return "TapeWidth(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
