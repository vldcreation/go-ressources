package money

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestNewMoney(t *testing.T) {
	amount := decimal.NewFromFloat(1000)
	m := NewMoney(amount, "IDR")

	if !m.Amount().Equal(amount) {
		t.Errorf("Expected amount to be %s, got %s", amount, m.Amount())
	}

	if m.Currency() != "IDR" {
		t.Errorf("Expected currency to be IDR, got %s", m.Currency())
	}
}

func TestMoneyAdd(t *testing.T) {
	amount1 := decimal.NewFromFloat(1000)
	amount2 := decimal.NewFromFloat(500)
	m1 := NewMoney(amount1, "IDR")
	m2 := NewMoney(amount2, "IDR")

	m3 := m1.Add(m2)

	if !m3.Amount().Equal(decimal.NewFromFloat(1500)) {
		t.Errorf("Expected amount to be 1500, got %s", m3.Amount())
	}

	if m3.Currency() != "IDR" {
		t.Errorf("Expected currency to be IDR, got %s", m3.Currency())
	}
}

func TestMoneySubtract(t *testing.T) {
	amount1 := decimal.NewFromFloat(1000)
	amount2 := decimal.NewFromFloat(500)
	m1 := NewMoney(amount1, "IDR")
	m2 := NewMoney(amount2, "IDR")

	m3 := m1.Subtract(m2)

	if !m3.Amount().Equal(decimal.NewFromFloat(500)) {
		t.Errorf("Expected amount to be 500, got %s", m3.Amount())
	}

	if m3.Currency() != "IDR" {
		t.Errorf("Expected currency to be IDR, got %s", m3.Currency())
	}
}

func TestMoneyFormat(t *testing.T) {
	amount := decimal.NewFromFloat(1000)
	m := NewMoney(amount, "IDR")

	lang := language.AmericanEnglish

	if m.Format(lang) != "IDR1,000" {
		t.Errorf("Expected formatted money to be IDR 1,000.00, got %s", m.Format(lang))
	}

	lang = language.Indonesian
	if m.Format(lang) != "Rp1.000" {
		t.Errorf("Expected formatted money to be IDR 1.000, got %s", m.Format(lang))
	}
}

func TestMoneySpellOut(t *testing.T) {
	amount := decimal.NewFromFloat(2_123_673_595)
	m := NewMoney(amount, "IDR")

	lang := language.Indonesian
	expected := "Dua Miliar Seratus Dua Puluh Tiga Juta Enam Ratus Tujuh Puluh Tiga Ribu Lima Ratus Sembilan Puluh Lima Rupiah"

	if m.SpellOut(V1, lang) != expected {
		t.Errorf("Expected formatted money to be %s, got %s", expected, m.SpellOut(V1, lang))
	}

	// test with default spelling
	m = NewMoney(amount, "USD")
	lang = language.Ukrainian
	expected = "Two Billion One Hundred Twenty Three Million Six Hundred Seventy Three Thousand Five Hundred Ninety Five USD"
	if m.SpellOut(V1, lang) != expected {
		t.Errorf("Expected formatted money to be %s, got %s", expected, m.SpellOut(V1, lang))
	}

	// test wit custom spelling that version already exists
	lang = language.English
	err := m.RegisterSpellings(V1, lang, func(amount int64) string {
		return fmt.Sprintf("CST %d", amount)
	})

	assert.EqualError(t, err, "spelling in en with version v1 already registered")

	// test with custom spelling
	lang = language.Ukrainian
	expected = "CST 2123673595"
	err = m.RegisterSpellings(V1, lang, func(amount int64) string {
		return fmt.Sprintf("CST %d", amount)
	})

	assert.NoError(t, err)

	if m.SpellOut(V1, lang) != expected {
		t.Errorf("Expected formatted money to be %s, got %s", expected, m.SpellOut(V1, lang))
	}
}
