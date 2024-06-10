package util

func CompareSliceString(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func Compare2DSliceString(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if !CompareSliceString(v, b[i]) {
			return false
		}
	}

	return true
}

func CompareSliceInt(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func Compare2DSliceInt(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if !CompareSliceInt(v, b[i]) {
			return false
		}
	}

	return true
}

func CompareSliceInt64(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func Compare2DSliceInt64(a, b [][]int64) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if !CompareSliceInt64(v, b[i]) {
			return false
		}
	}

	return true
}

func CheckStringInSlice(str string, arr []string) bool {
	for _, s := range arr {
		if str == s {
			return true
		}
	}

	return false
}
