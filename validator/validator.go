package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"reflect"
	"strings"
	"time"
)

const MaxAgeLimit = 100

func ValidateDOB(fl validator.FieldLevel) bool {
	dobStr := fl.Field().String()
	const layout = "2006-01-02"

	dob, err := time.Parse(layout, dobStr)
	if err != nil {
		return false
	}

	today := time.Now()

	if dob.After(today) {
		return false
	}

	age := today.Year() - dob.Year()

	if today.YearDay() < dob.YearDay() {
		age--
	}

	if age > MaxAgeLimit {
		return false
	}

	return true
}

func genderValidation(fl validator.FieldLevel) bool {
	gender := strings.ToLower(fl.Field().String())
	return gender == "male" || gender == "female" || gender == "other"
}

func maritalStatusValidation(fl validator.FieldLevel) bool {
	status := strings.ToLower(fl.Field().String())
	return status == "single" || status == "married"
}

func RegisterCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		v.RegisterValidation("valid_dob", ValidateDOB)
		v.RegisterValidation("gender", genderValidation)
		v.RegisterValidation("maritalstatus", maritalStatusValidation)
	}
}
