package psogc

import "github.com/go-playground/validator"

type CharacterClass int

const (
	HUmar CharacterClass = iota + 1
	HUnewearl
	HUcast
	HUcaseal
	RAmar
	RAmarl
	RAcast
	RAcaseal
	FOmar
	FOmarl
	FOnewm
	FOnewearl
)

func (c CharacterClass) String() string {
	if s, ok := characterClassEnum[c]; ok {
		return s
	}

	return "Unknown"
}

func (c CharacterClass) Code() int {
	if _, ok := characterClassEnum[c]; ok {
		return int(c)
	}

	return 0
}

var characterClassEnum = map[CharacterClass]string{
	HUmar:     "humar",
	HUnewearl: "hunewearl",
	HUcast:    "hucast",
	HUcaseal:  "hucaseal",
	RAmar:     "ramar",
	RAmarl:    "ramarl",
	RAcast:    "racast",
	RAcaseal:  "racaseal",
	FOmar:     "fomar",
	FOmarl:    "fomarl",
	FOnewm:    "fonewm",
	FOnewearl: "fonewearl",
}

type CharacterClassEnum struct {
	enum map[CharacterClass]string
}

func NewCharacterClassEnum() *CharacterClassEnum {
	return &CharacterClassEnum{
		enum: characterClassEnum,
	}
}

func (c *CharacterClassEnum) Codes() (codes []int) {
	for code := range c.enum {
		codes = append(codes, int(code))
	}

	return codes
}

func (c *CharacterClassEnum) Values() (values []string) {
	for _, value := range c.enum {
		values = append(values, value)
	}

	return values
}

func (c *CharacterClassEnum) OrderedValues() (values []string) {
	for i := 1; ; i++ {
		value := CharacterClass(i).String()
		if value == "Unknown" {
			break
		}
		values = append(values, value)
	}

	return values
}

func (c *CharacterClassEnum) ValuesValidator() func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		for _, value := range c.Values() {
			if value == fl.Field().String() {
				return true
			}
		}

		return false
	}
}
