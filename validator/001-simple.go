package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ======================================================
// ============ Struct Definitions ======================
// ======================================================
// user contains user information
type user struct {
	FirstName 		string 		`validate:"required"`
	LastName  		string 		`validate:"required"`
	Age	   			int    		`validate:"gte=0,lte=130"`
	Email	   		string 		`validate:"required,email"`
	Gender			string 		`validate:"oneof=male female prefer_not_to"`
	// Alias for `hexcolor|rgb|rgba|hsl|hsla`
	FavoriteColor 	string 		`validate:"iscolor"`						
	// A person can have multiple addresses. `dive` tells the validator to check each item in the slice individually
	Addresses		[]*addr 	`validate:"required,dive,required"`			
}

// addr houses a user's address information
type addr struct {
	Street		string 	`validate:"required"`
	City		string 	`validate:"required"`
	Planet		string 	`validate:"required"`
	Phone		string 	`validate:"required"`
}

// =============================================
// ============ Functions ======================
// =============================================
func Simple() {
	validate = validator.New(validator.WithRequiredStructEnabled())

	validateStruct()
	validateVariable()
}

// This function demonstrates how to validate a struct
func validateStruct() {
	address := &addr {
		Street: "Eavesdown Docks",
		City: "Persphone",
		Phone: "none",
	}
	user := &user {
		FirstName: "Badger",
		LastName: "Smith",
		Age: 135,
		Gender: "male",
		Email: "Badger.Smith@gmail.com",
		FavoriteColor: "#000-",
		Addresses: []*addr{address},
	}

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(user)

	if err != nil {
		// This check is only needed when your code could produce an invalid 
		// value for validation, such as interface with nil value 
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())			// The full path to the field (e.g. "User.FirstName")
			fmt.Println(err.Field())				// Just the field name (e.g. "FirstName")
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())					// The validation tag that failed (e.g. "required")
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())					// The type of the field (e.g. "string")
			fmt.Println(err.Value())				// The value that failed validation (e.g. "135" for age)
			fmt.Println(err.Param())				// The parameter for the tag (e.g. "0" for gte=0)
			fmt.Println()
		}
	}
}

// This function demonstrates how to validate a variable
func validateVariable() {
	myEmail := "joeybloggs.gmail.com"
	
	errs := validate.Var(myEmail, "required,email")

	if errs != nil {
		fmt.Println(errs)	// output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}
}
