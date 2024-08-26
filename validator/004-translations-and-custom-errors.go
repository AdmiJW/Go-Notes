package validator

import (
	"fmt"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

// This demonstrates how to perform validations with custom error messages and translations.
// `universal-translator` is a package that helps manage translations for different languages.
// `validator/v10/translations/en` provides English translations for the default error messages.

// Change this to your desired language `en` or `zh`
const ACCEPT_LANGUAGE = "zh"

func TranslationsAndCustomErrors() {
	// NOTE: omitting a lot of error checking for brevity

	en := en.New()
	zh := zh.New()
	// Creates a `universal-translator` instance that supports English and Chinese as the supported language.
	// English is selected as the fallback language.
	uni = ut.New(en, en, zh)

	// This is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	en_trans, _ := uni.GetTranslator("en")
	zh_trans, _ := uni.GetTranslator("zh")
	translator := en_trans
	if ACCEPT_LANGUAGE == "zh" { translator = zh_trans }

	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, en_trans)
	zh_translations.RegisterDefaultTranslations(validate, zh_trans)

	// You can specify your own in whichever language you like.
	translateAll(translator)
	translateIndividual(translator)
	translateOverride(translator)
}

func translateAll(trans ut.Translator) {
	type user struct {
		Username 	string 		`validate:"required"`
		Tagline 	string 		`validate:"required,lt=10"`
		Tagline2 	string 		`validate:"required,gt=1"`
	}

	u := user{
		Username: "",
		Tagline: "This tagline is way too long",
		Tagline2: "1",
	}

	err := validate.Struct(u)
	if err != nil {
		// Translate all error at once
		errs := err.(validator.ValidationErrors)

		// Returns a map with key = namespace & value = translated error
		// NOTICE: 2 errors are returned and you'll see something surprising
		// translations are i18n aware
		// eg: '10 characters' vs '1 character'
		fmt.Println(errs.Translate(trans))
	}
}

func translateIndividual(trans ut.Translator) {
	type user struct {
		Username 	string 		`validate:"required"`
	}

	var u user

	err := validate.Struct(u)
	if err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			fmt.Println(e.Translate(trans))
		}
	}
}

// Registers a custom translation for the `required` validation tag.
func translateOverride(trans ut.Translator) {
	validate.RegisterTranslation(
		"required", 
		trans, 
		func(ut ut.Translator) error {
			// see universal-translator for details
			if ut.Locale() == "en" {
				ut.Add("required", "{0} must have a value!", true)
			}
			return ut.Add("required", "{0}必填!", true)
		},
		func (ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		},
	)

	type user struct {
		Username 	string 		`validate:"required"`
	}

	var u user

	err := validate.Struct(u)
	if err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			// Can translate each error one at a time.
			fmt.Println(e.Translate(trans))
		}
	}
}