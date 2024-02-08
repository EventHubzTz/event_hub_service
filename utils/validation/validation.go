package validation

import (
	"reflect"
	"regexp"
	"strings"
	"unicode"

	"github.com/EventHubzTz/event_hub_service/database"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func IsUnique(table string, field string, value string) bool {
	var db = database.DB()

	rows, _ := db.Raw("select * from "+table+" where "+field+" = ?", value).Rows()
	defer rows.Close()
	if rows.Next() {
		return false
	}
	return true
}

func Validate(i interface{}) map[string][]string {
	validate := validator.New()
	//REGISTER VALIDATOR "UNIQUE"
	validate.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
		return IsUnique(strings.Split(fl.Param(), ".")[0], strings.Split(fl.Param(), ".")[1], fl.Field().String())
	})
	//REGISTER VALIDATOR "COUNTRY CODE"
	validate.RegisterValidation("country_code", func(fl validator.FieldLevel) bool {
		return IsPhoneNumberMatchingCountryCode(fl.Param(), fl.Field().String())
	})

	//REGISTER FUNCTION
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	errs := validate.Struct(i)

	errors := make(map[string][]string)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			errors[err.Field()] = append(errors[err.Field()], formatErrorMessage(err.Field(), err.Tag(), err.Param()))
		}
		return errors
	}

	return nil
}

func formatErrorMessage(fieldName, tag, value string) string {
	var errorMessage = ""
	var name = strings.ReplaceAll(fieldName, "_", " ")
	switch tag {
	case "required":
		errorMessage = "The " + name + " field is required."
		break
	case "min":
		errorMessage = "The " + name + " field minimum length is " + value + "."
		break
	case "max":
		errorMessage = "The " + name + " field maximum length is " + value + "."
		break
	case "email":
		errorMessage = "Invalid email address.."
		break
	case "unique":
		errorMessage = "The " + name + " field required to be unique."
		break
	case "country_code":
		errorMessage = "The " + name + " field required to use the correct country code."
		break

	}
	return errorMessage
}

func IsEmail(email string) bool {
	if IsBlank(email) {
		return false
	}
	pattern := `^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`
	matched, _ := regexp.MatchString(pattern, email)
	if !matched {
		return false
	}
	return true
}
func IsBlank(str string) bool {
	strLen := len(str)
	if str == "" || strLen == 0 {
		return true
	}
	for i := 0; i < strLen; i++ {
		if unicode.IsSpace(rune(str[i])) == false {
			return false
		}
	}
	return true
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

func IsPhoneNumberMatchingCountryCode(countryCode string, phoneNumber string) bool {
	switch countryCode {
	case "TZ":
		return IsTanzaniaPhoneNumber(phoneNumber)
	}
	return true
}

func IsTanzaniaPhoneNumber(phoneNumber string) bool {
	pattern := `^255([0-9\s\-\+\(\)]*)$`
	matched, _ := regexp.MatchString(pattern, phoneNumber)
	print(matched)
	if !matched {
		return false
	}
	return true
}
