// Code generated by "goenum -type TapeWidth"; DO NOT EDIT.

package ptouchgo

func (i TapeWidth) Valid() bool {
	switch {
	case i == 0:
		return true
	case i == 4:
		return true
	case i == 6:
		return true
	case i == 9:
		return true
	case i == 12:
		return true
	case i == 18:
		return true
	case i == 24:
		return true
	case i == 62:
		return true
	default:
		return false
	}
}
