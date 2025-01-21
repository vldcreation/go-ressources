// @Author Vicktor Lambok Desrony
package podify

import (
	"strconv"
)

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
	if len(bank.number) >= 12 && len(bank.number) <= 19 {
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

// Utillity
func ToInt(in string) int {
	var out int
	out, _ = strconv.Atoi(in)
	return out
}

func isPrefDiscover(bank string) bool {
	switch {
	case string(bank[0:2]) == "65", string(bank[0:4]) == "6011":
		return true
	case ToInt(string(bank[0:6])) >= 622126 && ToInt(string(bank[0:6])) <= 622925:
		return true
	case ToInt(string(bank[0:3])) >= 644 && ToInt(string(bank[0:3])) <= 649:
		return true
	}

	return false
}

func isPrefMaestro(pref string) bool {
	switch {
	case pref == "50":
		return true
	case ToInt(pref) >= 56 && ToInt(pref) <= 59:
		return true
	}

	return false
}

func Case1(input string) string {
	bank := NewBanker(input)
	return bank.Network(Bank{number: input})
}
