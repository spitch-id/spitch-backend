package config

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	Validator  *validator.Validate
	Translator *ut.UniversalTranslator
)

func NewValidator() (*validator.Validate, *ut.UniversalTranslator) {
	validate := validator.New()

	enLocale := en.New()
	idLocale := id.New()

	uni := ut.New(enLocale, enLocale, idLocale)

	return validate, uni
}
