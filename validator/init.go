package validator

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// use a single instance, it caches struct info
var validate *validator.Validate
var uni *ut.UniversalTranslator

func Init() {}