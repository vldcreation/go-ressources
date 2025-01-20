package podify

import "fmt"

type Bank struct {
	// properties
}

type TypeBank interface {
	Type(input string) string // out: type bank name
}

// // Bank{} // mastercard

// return mastercard.Type()

func In(in string, arr []string) bool {
	for _, str := range arr {
		if str == in {
			return true
		}
	}

	return false
}

func InIntRange(in, start, end int) bool {
	if in >= start && in <= end {
		return true
	}

	return false
}

func genSlice(start, end int) []string {
	arr := make([]string, 0)
	for start < end {
		arr = append(arr, fmt.Sprintf("%d", start))
		start++
	}

	return arr
}

func isPrefDiscover(bank string) bool {
	bankPref := []string{}
	bankPref = append(bankPref, "6011")
	bankPref = append(bankPref, genSlice(622126, 622925)...)
	bankPref = append(bankPref, genSlice(644, 649)...)
	bankPref = append(bankPref, "65")
	switch {
	case string(bank[0:2]) == "65", string(bank[0:4]) == "6011", In(string(bank[0:3]), bankPref), In(string(bank[0:6]), bankPref):
		return true
	}

	return false
}

func isPrefMaestro(pref string) bool {
	bankPref := []string{}
	bankPref = append(bankPref, "50")
	bankPref = append(bankPref, genSlice(56, 59)...)
	return In(pref, bankPref)
}

func Case1(input string) string {
	pref := input[0:2]
	if len(input) == 15 {
		if string(pref) == "34" || string(pref) == "37" {
			return "American Express"
		}
	} else if len(input) == 14 {
		if string(pref) == "38" || string(pref) == "39" {
			return "Diners Club"
		}
	} else if len(input) == 16 {
		// part of visa
		if string(pref[0]) == "4" {
			return "Visa"
		}

		// MC
		if pref == "51" || pref == "52" || pref == "53" || pref == "54" || pref == "55" {
			return "MasterCard"
		}

		// Discover
		if isPrefDiscover(input) {
			return "Discover"
		}

		if isPrefMaestro(pref) {
			return "Maestro"
		}

	} else if len(input) == 13 || len(input) == 16 || len(input) == 19 {
		if string(pref[0]) == "4" {
			return "Visa"
		}

		if len(input) == 19 && isPrefDiscover(input) {
			return "Discover"
		}

		if len(input) == 19 && isPrefMaestro(pref) {
			return "Maestro"
		}
	} else if InIntRange(len(input), 12, 19) {
		// Maestro
		if isPrefMaestro(pref) {
			return "Maestro"
		}
	}

	return "unknown"
}
