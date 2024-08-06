// @Author: Vicktor Lambok Desrony
// Thanks to golang.org, github.com and AI for the inspiration and idea.
// This is the implementation of the money package.
// That can be used to format money and spell out money in different languages.
package money

import (
	"fmt"
	"strings"

	"github.com/leekchan/accounting"
	"github.com/shopspring/decimal"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

type spellingFunc func(int64) string
type internalSpellingsWrapper struct {
	version Version
	fn      spellingFunc
}

type internalSpellings map[Version]internalSpellingsWrapper

func wrappedSpelling(v Version, f spellingFunc) internalSpellingsWrapper {
	return internalSpellingsWrapper{
		version: v,
		fn:      f,
	}
}

func wrappendSpellingMap(v Version, f spellingFunc) internalSpellings {
	return internalSpellings{
		v: wrappedSpelling(v, f),
	}
}

type Money struct {
	amount    decimal.Decimal
	currency  string
	spellings map[language.Tag]internalSpellings
}

func NewMoney(amount decimal.Decimal, currency string) *Money {
	m := &Money{
		amount:    amount,
		currency:  currency,
		spellings: defaultSpellings(),
	}

	return m
}

func (m *Money) RegisterSpellings(v Version, lang language.Tag, fn spellingFunc) error {
	if lang == language.Und {
		return fmt.Errorf("language cannot be empty")
	}

	if fn == nil {
		return fmt.Errorf("function cannot be empty")
	}

	if fn, ok := m.spellings[lang]; ok {
		if spell := fn[v]; spell.version != "" {
			return fmt.Errorf("spelling in %s with version %v already registered", lang, v)
		}
	}

	m.spellings[lang] = wrappendSpellingMap(v, fn)

	return nil
}

func (m *Money) Amount() decimal.Decimal {
	return m.amount
}

func (m *Money) Currency() string {
	return m.currency
}

func (m *Money) Format(lang language.Tag) string {
	return formatMoney(m.amount, m.currency, lang)
}

func (m *Money) Add(money *Money) *Money {
	if m.currency != money.currency {
		return nil
	}

	return NewMoney(m.amount.Add(money.amount), m.currency)
}

func (m *Money) Subtract(money *Money) *Money {
	if m.currency != money.currency {
		return nil
	}

	return NewMoney(m.amount.Sub(money.amount), m.currency)
}

func (m *Money) SpellOut(v Version, lang language.Tag) string {
	cur, err := currency.ParseISO(m.currency)
	if err != nil {
		return fmt.Sprintf("%s%s", m.currency, accounting.FormatNumberDecimal(m.amount, 2, ",", "."))
	}

	p := message.NewPrinter(lang)
	if fn, ok := m.spellings[lang]; ok {
		if spell := fn[v]; spell.version != "" {
			return spell.fn(m.amount.BigInt().Int64())
		} else {
			return fn[V1].fn(m.amount.BigInt().Int64())
		}
	}

	return p.Sprintf(SpellOutEnglishV1(m.amount.BigInt().Int64())) + " " + cur.String()
}

// FormatMoney format money by currency code (ISO 4217)
func formatMoney(value decimal.Decimal, currencyCode string, lang language.Tag) string {
	cur, err := currency.ParseISO(currencyCode)
	if err != nil {
		return fmt.Sprintf("%s%s", currencyCode, accounting.FormatNumberDecimal(value, 2, ",", "."))
	}

	scale, _ := currency.Cash.Rounding(cur) // fractional digits
	unit, _ := value.Float64()
	dec := number.Decimal(unit, number.Scale(scale))

	return message.NewPrinter(lang).Sprintf("%v%v", currency.Symbol(cur), dec)
}

var indonesianUnits = []string{"", "Satu", "Dua", "Tiga", "Empat", "Lima", "Enam", "Tujuh", "Delapan", "Sembilan", "Sepuluh", "Sebelas"}

func SpellOutIndonesianV1(n int64) string {
	if n == 0 {
		return ""
	}

	if n < 0 {
		return "Minus " + SpellOutIndonesianV1(-n)
	}

	if n < 12 {
		return indonesianUnits[n]
	} else if n < 20 {
		return indonesianUnits[n-10] + " Belas"
	} else if n < 100 {
		return indonesianUnits[n/10] + " Puluh " + SpellOutIndonesianV1(n%10)
	} else if n < 200 {
		return "Seratus " + SpellOutIndonesianV1(n%100)
	} else if n < 1000 {
		return indonesianUnits[n/100] + " Ratus " + SpellOutIndonesianV1(n%100)
	} else if n < 2000 {
		return "Seribu " + SpellOutIndonesianV1(n%1000)
	} else if n < 1000000 {
		return SpellOutIndonesianV1(n/1000) + " Ribu " + SpellOutIndonesianV1(n%1000)
	} else if n < 1000000000 {
		return SpellOutIndonesianV1(n/1000000) + " Juta " + SpellOutIndonesianV1(n%1000000)
	} else if n < 1000000000000 {
		return SpellOutIndonesianV1(n/1000000000) + " Miliar " + SpellOutIndonesianV1(n%1000000000)
	} else if n < 1000000000000000 {
		return SpellOutIndonesianV1(n/1000000000000) + " Triliun " + SpellOutIndonesianV1(n%1000000000000)
	}
	return "Angka terlalu besar"
}

var (
	ones      = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	tens      = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	teens     = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	thousands = []string{"", "Thousand", "Million", "Billion", "Trillion"}
)

func SpellOutEnglishV1(amount int64) string {
	if amount == 0 {
		return "Zero"
	}

	words := []string{}
	for i := 0; amount > 0; i++ {
		if amount%1000 != 0 {
			words = append([]string{spellOutThreeDigits(amount%1000) + " " + thousands[i]}, words...)
		}
		amount /= 1000
	}

	return strings.TrimSpace(strings.Join(words, " "))
}

func spellOutThreeDigits(n int64) string {
	if n == 0 {
		return ""
	}

	words := []string{}

	if n >= 100 {
		words = append(words, ones[n/100]+" Hundred")
		n %= 100
	}

	if n >= 20 {
		words = append(words, tens[n/10])
		n %= 10
	} else if n >= 10 {
		words = append(words, teens[n-10])
		return strings.Join(words, " ")
	}

	if n > 0 {
		words = append(words, ones[n])
	}

	return strings.TrimSpace(strings.Join(words, " "))
}

func defaultSpellings() map[language.Tag]internalSpellings {
	bakedSpellings := make(map[language.Tag]internalSpellings, 0)
	bakedSpellings[language.Indonesian] = wrappendSpellingMap(V1, func(i int64) string {
		s := strings.Join(strings.Fields(SpellOutIndonesianV1(i)), " ")
		return s + " Rupiah"
	})
	bakedSpellings[language.English] = wrappendSpellingMap(V1, SpellOutEnglishV1)

	return bakedSpellings
}
