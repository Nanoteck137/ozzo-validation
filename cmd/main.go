package main

import (
	"log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/newer"
	"github.com/kr/pretty"
)

type Test struct {
	Test string `json:"test"`
}

func (b Test) Validate(validator newer.Validator) error {
	return validator.Struct(
		&b,
		validator.Field(&b.Test, validation.Required, validation.Length(5, 10)),
	)
}

type Body struct {
	A    string `json:"a"`
	Test []Test `json:"test"`
}

func (b Body) Validate(validator newer.Validator) error {
	return validator.Struct(
		&b,
		validator.Field(&b.A, validation.Required, validation.Length(5, 10)),
		validator.Field(&b.Test),
	)
}

func main() {
	b := Body{
		A: "hell",
		Test: []Test{
			{
			},
			{
			},
		},
	}

	v := newer.NormalValidator{}
	err := b.Validate(&v)
	if err != nil {
		pretty.Println(err)
		log.Fatal(err)
	}
}
