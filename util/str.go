package util

func CompareSlice(a, b []string) bool {
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

func Compare2DSlice(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if !CompareSlice(v, b[i]) {
			return false
		}
	}

	return true
}
