package medium

func IntToRoman(num int) string {
	var roman string
	if num/1000 > 0 {
		for i := int(num / 1000); i > 0; i-- {
			roman += "M"
		}
		num = num - (int(num/1000) * 1000)
	}

	if num-900 >= 0 {
		roman += "CM"
		num = num - 900
	}

	if num-500 >= 0 {
		roman += "D"
		num = num - 500
	}

	if num-400 >= 0 {
		roman += "CD"
		num = num - 400
	}

	if num/100 > 0 {
		for i := int(num / 100); i > 0; i-- {
			roman += "C"
		}
		num = num - (int(num/100) * 100)
	}

	if num-90 >= 0 {
		roman += "XC"
		num = num - 90
	}

	if num-50 >= 0 {
		roman += "L"
		num = num - 50
	}

	if num-40 >= 0 {
		roman += "XL"
		num = num - 40
	}

	if num/10 > 0 {
		for i := int(num / 10); i > 0; i-- {
			roman += "X"
		}
		num = num - (int(num/10) * 10)
	}

	if num-9 >= 0 {
		roman += "IX"
		num = num - 9
	}

	if num-5 >= 0 {
		roman += "V"
		num = num - 5
	}

	if num-4 >= 0 {
		roman += "IV"
		num = num - 4
	}

	if num/1 > 0 {
		for i := int(num / 1); i > 0; i-- {
			roman += "I"
		}
	}

	return roman
}
