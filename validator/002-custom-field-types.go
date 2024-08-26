package validator

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// This function demonstrates custom validation for fields of specific types - specifically those related to
// `sql.Null*` types (like `sql.NullString`, `sql.NullInt64`, etc.).
// THese types are often used in database applications to handle nullable fields in a way that's compatible with
// SQL databases.

// ======================================================
// ============ Struct Definitions ======================
// ======================================================
// DBBackedUser User struct
type dbBackedUser struct {
	Name		sql.NullString	`validate:"required"`
	Age			sql.NullInt64	`validate:"required,gte=0,lte=130"`
}

// =============================================
// ============ Functions ======================
// =============================================
func CustomFieldTypes() {
	validate = validator.New()

	// Register all sql.Null* types to use the `validateValuer` CustomTypeFunc
	// Instead of using default validation for these fields, the validator will use the logic in `validateValuer`
	// to extract the actual value from these types and validate it.
	validate.RegisterCustomTypeFunc(
		validateValuer, 
		sql.NullString{}, 
		sql.NullInt64{}, 
		sql.NullBool{}, 
		sql.NullFloat64{},
	)

	// Validate the struct
	user := dbBackedUser {
		Name: sql.NullString{String: "", Valid: true},
		Age: sql.NullInt64{Int64: 0, Valid: true},
	}
	err := validate.Struct(user)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func validateValuer(field reflect.Value) interface{} {
	// First, check if the field implements the `driver.Valuer` interface using a type assertion
	if valuer, ok := field.Interface().(driver.Valuer); ok {
		// If it does, use the `Value()` method to get the actual value
		value, err := valuer.Value()

		// If there's no error, return the value to the validator, which will then
		// apply the appropriate validation rules to it
		if err == nil { return value }

		// otherwise, handle the error as you need
	}
	return nil
}