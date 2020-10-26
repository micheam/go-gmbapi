package pointer

import "time"

// StringPtrDeref は、string prt が nil で無い場合はその値を返却し、
// nil の場合は string default_ を返却する。
func StringPtrDeref(ptr *string, defaultVal string) string {
	if ptr != nil {
		return *ptr
	}
	return defaultVal
}

// IntPtrDeref は、int prt が nil で無い場合はその値を返却し、
// nil の場合は int default_ を返却する。
func IntPtrDeref(ptr *int, defaultVal int) int {
	if ptr != nil {
		return *ptr
	}
	return defaultVal
}

// StringPrtTimeParse は、 string ptr が nil で無い場合にはその値を
// time.Parse する。nil の場合は time.Time _default を返却する。
func StringPrtTimeParse(ptr *string, format string, defaultVal time.Time) (time.Time, error) {
	if ptr != nil {
		return time.Parse(format, *ptr)
	}
	return defaultVal, nil
}
