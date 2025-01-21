package podify

import "fmt"

type Bank struct {
	number string
}

type Banker interface {
	Network(bank Bank) string // out: type bank name
}

var ListBank = []Banker{
	&AmericanExpress{},
	&DinersClub{},
	&Visa{},
	&MasterCard{},
	&Discover{},
	&Maestro{},
	&Unknown{},
}

func NewBanker(number string) Banker {
	bank := Bank{
		number: number,
	}

	for _, b := range ListBank {
		if b.Network(bank) != "unknown" {
			return b
		}
	}

	return &Unknown{}
}

// American Expres Typer
type AmericanExpress struct{}

func (a *AmericanExpress) Network(bank Bank) string {
	if len(bank.number) == 15 {
		if bank.number[0:2] == "34" || bank.number[0:2] == "37" {
			return "American Express"
		}
	}

	return "unknown"
}

// Diners Club Typer
type DinersClub struct{}

func (d *DinersClub) Network(bank Bank) string {
	if len(bank.number) == 14 {
		if bank.number[0:2] == "38" || bank.number[0:2] == "39" {
			return "Diners Club"
		}
	}

	return "unknown"
}

// Visa Typer
type Visa struct{}

func (v *Visa) Network(bank Bank) string {
	if len(bank.number) == 13 || len(bank.number) == 16 || len(bank.number) == 19 {
		if bank.number[0:1] == "4" {
			return "Visa"
		}
	}

	return "unknown"
}

// MasterCard Typer
type MasterCard struct{}

func (m *MasterCard) Network(bank Bank) string {
	if len(bank.number) == 16 {
		if bank.number[0:2] == "51" || bank.number[0:2] == "52" ||
			bank.number[0:2] == "53" || bank.number[0:2] == "54" ||
			bank.number[0:2] == "55" {
			return "MasterCard"
		}
	}

	return "unknown"
}

// Discover Typer
type Discover struct{}

func (d *Discover) Network(bank Bank) string {
	if len(bank.number) == 16 || len(bank.number) == 19 {
		if isPrefDiscover(bank.number) {
			return "Discover"
		}
	}

	return "unknown"
}

// Maestro Typer
type Maestro struct{}

func (m *Maestro) Network(bank Bank) string {
	if InIntRange(len(bank.number), 12, 19) {
		if isPrefMaestro(bank.number[0:2]) {
			return "Maestro"
		}
	}

	return "unknown"
}

// Unknown Typer
type Unknown struct{}

func (u *Unknown) Network(bank Bank) string {
	return "unknown"
}

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
	bank := NewBanker(input)
	return bank.Network(Bank{number: input})
}
