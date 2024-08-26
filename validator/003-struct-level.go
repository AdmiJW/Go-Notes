package validator

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ======================================================
// ============ Struct Definitions ======================
// ======================================================
// Wraps the normal validator.FieldError struct to provide custom error handling,
// including the generation of error messages in JSON format.
type validationError struct {
	Namespace		string			`json:"namespace"` 		// can differ when a custom TagnameFunc is registered or
	Field			string			`json:"field"`			// by passing alt name to ReportError like below
	StructNamespace	string			`json:"structNamespace"`
	StructField		string			`json:"structField"`
	Tag				string			`json:"tag"`
	ActualTag		string			`json:"actualTag"`
	Kind			string			`json:"kind"`
	Type			string			`json:"type"`
	Value			interface{}		`json:"value"`
	Param			string			`json:"param"`
	Message			string			`json:"message"`
}

// Gender enum, with three values `Male`, `Female`, and `Intersex`.
// `String()` method provides a string representation of the enum value. Returns "unknown" if the value is invalid.
type gender uint
const (
	Male gender = iota + 1
	Female
	Intersex
)
func (g gender) String() string {
	terms := []string {"Male", "Female", "Intersex"}
	if (g < Male || g > Intersex) { return "unknown " }
	return terms[g]
}

type user2 struct {
	FirstName		string		`json:"fname"`
	LastName		string		`json:"lname"`
	Age				uint8		`validate:"gte=0,lte=130"`
	Email			string		`json:"e-mail" validate:"required,email"`
	FavouriteColor	string		`validate:"hexcolor|rgb|rgba"`
	Addresses		[]*addr2	`validate:"required,dive,required"`
	// Uses custom validation
	Gender			gender		`json:"gender" validate:"required,gender_custom_validation"`
}
type addr2 struct {
	Street		string		`validate:"required"`
	City		string		`validate:"required"`
	Planet		string		`validate:"required"`
	Phone		string		`validate:"required"`
}

// =============================================
// ============ Functions ======================
// =============================================
func StructLevel() {
	validate = validator.New()

	// Registers a function that retrieves field names from the JSON tags
	// instead of using the struct field names directly.
	// This allows validation error messages to use the same field names
	// as the JSON representation.
	validate.RegisterTagNameFunc(func (fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" { return "" }
		return name	
	})

	// Register validation for User
	// NOTE: only have to register a non-pointer type for User, validator
	// internally dereferences during it's type checks
	validate.RegisterStructValidation(UserStructLevelValidation, user2{})

	// Register a custom validation for user's gender field
	// validates that an enum is within the internal (string representation is not "unknown")
	err := validate.RegisterValidation(
		"gender_custom_validation", 
		func(fl validator.FieldLevel) bool {
			value := fl.Field().Interface().(gender)
			return value.String() != "unknown"
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Build User info to mock normally posted data
	address := &addr2 {
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone: "none",
		City: "Unknown",
	}
	user := &user2 {
		FirstName: "",
		LastName: "",
		Age: 45,
		Email: "Badger.Smith@gmail",
		FavouriteColor: "#000",
		Addresses: []*addr2{address},
	}

	// Returns InvalidValidationError for bad validation input, nil or ValidationErrors ( []FieldError )
	err = validate.Struct(user)
	if err != nil {
		// This check is only needed when your code could produce an invalid
		// value for validation, such as interface with nil value
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		// If errors are found, they are wrapped in the `validationError` struct
		// then marshalled into pretty-printed JSON format for easy reading
		for _, err := range err.(validator.ValidationErrors) {
			e := validationError{
				Namespace: err.Namespace(),
				Field: err.Field(),
				StructNamespace: err.StructNamespace(),
				StructField: err.StructField(),
				Tag: err.Tag(),
				ActualTag: err.ActualTag(),
				Kind: fmt.Sprintf("%v", err.Kind()),
				Type: fmt.Sprintf("%v", err.Type()),
				Value: fmt.Sprintf("%v", err.Value()),
				Param: err.Param(),
				Message: err.Error(),
			}

			indent, err := json.MarshalIndent(e, "", "  ")
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(indent))
		}

		// From here, you can create your own error messages in whatever format, language you wish
	}
	// Save user to database
}

// UserStructLevelValidation contains custom struct level validations that don't always make sense
// at the field validation level. For example, this function validates that either FirstName or
// LastName exist; could have done this with a custom field validation, but then would have had to
// add it to both fields duplicating the logic + overhead. This way it's only validated once.
//
// You may ask why wouldn't this be done outside of validator?
// Because doing this way hooks right into validator and you can combine with validation tags and
// still have a common error output format.
func UserStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(user2)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		sl.ReportError(user.FirstName, "fname", "FirstName", "fnameorlname", "")
		sl.ReportError(user.LastName, "lname", "LastName", "fnameorlname", "")
	}

	// Can do more, even with different tag than "fnameorlname"
}