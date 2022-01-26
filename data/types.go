package data

import "fmt"

type Source interface {
	String() string
}

type BookSource struct {
	isbn  string
	title string
	page  int
}

func (book *BookSource) String() string {
	//TODO: implement
	panic("not implemented")
}

type Website struct {
	url string
}

func (w *Website) String() string {
	//TODO: implement
	panic("not implemented")
}

type UnitPrefix int

const (
	Milli UnitPrefix = iota
	Centi
	Deci
	None
	Deca
	Hecto
	Kilo
)

func (u UnitPrefix) String() string {
	switch u {
	case Milli:
		return "m"
	case Centi:
		return "c"
	case Deci:
		return "d"
	case None:
		return ""
	case Deca:
		return "da"
	case Hecto:
		return "h"
	case Kilo:
		return "k"
	default:
		panic("unknown unit prefix")
	}
}

type Unit struct {
	ShortUnit  string
	LongUnit   string
	UnitPrefix UnitPrefix
}

func Liter() Unit {
	return Unit{
		ShortUnit:  "l",
		LongUnit:   "liter",
		UnitPrefix: None,
	}
}

func Gram() Unit {
	return Unit{
		ShortUnit:  "g",
		LongUnit:   "gram",
		UnitPrefix: None,
	}
}

func Tablespoon() Unit {
	return Unit{
		ShortUnit:  "tbsp",
		LongUnit:   "tablespoon",
		UnitPrefix: None,
	}
}

func Teaspoon() Unit {
	return Unit{
		ShortUnit:  "tsb",
		LongUnit:   "teaspoon",
		UnitPrefix: None,
	}
}

func (u Unit) String() string {
	return fmt.Sprintf("%s%s", u.UnitPrefix.String(), u.ShortUnit)
}

type Volume struct {
	amount int
	unit   Unit
}

type Ingredient struct {
	name        string
	description string
}

type Recipe struct {
	source      Source
	ingredients map[string]struct {
		volume     Volume
		ingredient Ingredient
	}
}

type Session struct {
	SessionId        string
	SessionTokenHash string
	XsrfTokenHash    string
	UserId           string
}
