package commons

import (
	"bytes"
	"errors"
	"math/rand"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	"github.org/eventmodeling/ecommerce/pkg/support/enums"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

/*
This function is responsible for concatenate strings using bytes.Buffer

values: is the slice of strings that will be concatenated
*/
func ConcatenateStrings(values ...string) string {
	var sb strings.Builder

	for _, str := range values {
		sb.WriteString(str)
	}

	return sb.String()
}

/*
This function is responsible to verify if contains a expected int in a list of ints

integersList: is the list of ints that will be verified

expectedInteger: is the expected int to be verified
*/
func ContainsInt(integersList []int, expectedInteger int) bool {
	for _, item := range integersList {
		if item == expectedInteger {
			return true
		}
	}
	return false
}

/*
This function is responsible to verify if contains a expected int in a list of strings

stringList: is the list of strings that will be verified

expectedString: is the expected string to be verified
*/
func ContainsString(stringList []string, expected string) bool {
	for _, item := range stringList {
		if item == expected {
			return true
		}
	}
	return false
}

/*
This function is responsible to verify if contains a expected uuid in a list of uuids

uuidList: is the list of uuids that will be verified

expected: is the expected uuid to be verified
*/
func ContainsUUID(uuidList []uuid.UUID, expected uuid.UUID) bool {
	for _, item := range uuidList {
		if item == expected {
			return true
		}
	}
	return false
}

/*
This function is responsible for returning the UUID of the process
for a control of the logs.
*/
func GetShortUUID() string {
	return uuid.New().String()[:5]
}

/*
This function function is responsible for returning the last string
of a sequence of strings, separated by the delimiter.
*/

func GetLastString(str, delimiter string) string {
	return str[strings.LastIndex(str, delimiter)+1:]
}

/*
IfThenElse evaluates a condition, if true returns the first parameter otherwise the second
*/
func IfThenElse(condition bool, a any, b any) any {
	if condition {
		return a
	}
	return b
}

/*
Returns the first non-empty string in the list.
Returns an empty string if all values are empty.
*/

func Coalesce(values ...string) string {
	for _, value := range values {
		if StringIsNotEmpty(value) {
			return value
		}
	}
	return ""
}

/*
This function is responsible to verify if a string is not empty

value: is the string value to verify
*/
func StringIsNotEmpty(value string) bool {
	return len(strings.TrimSpace(value)) > 0
}

/*
This function is responsible to verify if a string is empty

value: is the string value to verify
*/
func StringIsEmpty(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

/*
This function is responsible to generate a random string and number

length: is the length value of wanted
*/
func RandomStringAndNumber(length int) string {
	charset := enums.NAPP_COMMONS_RANDOM_STRING_CHARSET
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

/*
This function is responsible to checks if there is a first and last name

fullName: is the full name for verification
*/
func ValidateFullName(fullName string) bool {
	nameSplited := strings.Split(fullName, " ")

	if len(nameSplited) < 2 {
		return false
	}

	for _, name := range nameSplited {
		if len(name) < 1 {
			return false
		}
	}

	return true
}

/*
This function is responsible to checks if e-mail is valid

email: is the e-mail for validation
*/
func ValidateEmail(email string) bool {
	email = strings.TrimSpace(email)
	result, _ := regexp.MatchString(enums.NAPP_COMMONS_REGEX_EMAIL_MATCH_STRING, email)
	return result
}

/*
This function is responsible for concatenate strings and create a new error object.

strings: is the slice of strings that will be concatenated
*/
func ConcatenateErrorStrings(strings ...string) error {
	return errors.New(ConcatenateStrings(strings...))
}

/*
This function is responsible for return the difference of days between two dates. Tip: The biggest value should be passed in the first parameter

date1: is the first date

date2: is the second date
*/
func GetDiffDays(date1, date2 time.Time) int {
	return int(date1.Sub(date2).Hours() / 24)
}

/*
Function responsible for converting strings (false or true) to bool
*/
func ParseBool(value string) bool {
	return StringIsNotEmpty(value) && strings.ToLower(value) == "true"
}

/*
Function responsible removes every all that is not a int
*/
func CleanNonInt(value string) string {
	buf := bytes.NewBufferString("")
	for _, r := range value {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	return buf.String()
}

/*
Function responsible for get content type of filename ex. .png will return image/png
*/
func GetContentTypeOf(fileName string) string {
	extension := filepath.Ext(fileName)

	switch extension {
	case enums.NAPP_COMMONS_CONTENT_TYPE_SUFFIX_PNG:
		return enums.NAPP_COMMONS_CONTENT_TYPE_PNG
	case enums.NAPP_COMMONS_CONTENT_TYPE_SUFFIX_JPG:
		return enums.NAPP_COMMONS_CONTENT_TYPE_JPG
	case enums.NAPP_COMMONS_CONTENT_TYPE_SUFFIX_JPEG:
		return enums.NAPP_COMMONS_CONTENT_TYPE_JPEG
	case enums.NAPP_COMMONS_CONTENT_TYPE_SUFFIX_GIF:
		return enums.NAPP_COMMONS_CONTENT_TYPE_GIF
	case enums.NAPP_COMMONS_CONTENT_TYPE_SUFFIX_TIF:
		return enums.NAPP_COMMONS_CONTENT_TYPE_TIF
	case enums.NAPP_COMMONS_CONTENT_TYPE_SUFFIX_TIFF:
		return enums.NAPP_COMMONS_CONTENT_TYPE_TIFF
	case enums.NAPP_COMMONS_CONTENT_TYPE_SUFFIX_CSS:
		return enums.NAPP_COMMONS_CONTENT_TYPE_CSS
	case enums.NAPP_COMMONS_CONTENT_TYPE_SUFFIX_JS:
		return enums.NAPP_COMMONS_CONTENT_TYPE_JS
	case enums.NAPP_COMMONS_CONTENT_TYPE_SUFFIX_JSON:
		return enums.NAPP_COMMONS_CONTENT_TYPE_JSON
	case enums.NAPP_COMMONS_CONTENT_TYPE_SUFFIX_PDF:
		return enums.NAPP_COMMONS_CONTENT_TYPE_PDF
	case enums.NAPP_COMMONS_CONTENT_TYPE_SUFFIX_PPT:
		return enums.NAPP_COMMONS_CONTENT_TYPE_PPT
	case enums.NAPP_COMMONS_CONTENT_TYPE_SUFFIX_DOC:
		return enums.NAPP_COMMONS_CONTENT_TYPE_DOC
	case enums.NAPP_COMMONS_CONTENT_TYPE_SUFFIX_XLS:
		return enums.NAPP_COMMONS_CONTENT_TYPE_XLS
	case enums.NAPP_COMMONS_CONTENT_TYPE_SUFFIX_XLSX:
		return enums.NAPP_COMMONS_CONTENT_TYPE_XLSX
	case enums.NAPP_COMMONS_CONTENT_TYPE_SUFFIX_XML:
		return enums.NAPP_COMMONS_CONTENT_TYPE_XML
	case enums.NAPP_COMMONS_CONTENT_TYPE_SUFFIX_SVG:
		return enums.NAPP_COMMONS_CONTENT_TYPE_SVG
	}

	return ""
}

/*
This function divides the email into two parts using the @ symbol as a separator.
It then creates a new string with asterisks for the first half of the username and
adds the second half of the username and domain to the end.

For example, if you pass the email "myemail@email.com.br" to the function, it will return "mye***l@email.com.br".

email: email to be encrypted
*/
func MaskEmail(email string) (string, error) {
	if !ValidateEmail(email) {
		return "", errors.New(enums.NAPP_COMMONS_MASK_EMAIL_ERROR_MESSAGE)
	}

	parts := strings.Split(email, "@")
	username := parts[0]
	domain := parts[1]

	n := len(username)
	var maskedUsername string

	if n > 4 {
		maskedUsername = ConcatenateStrings(username[:3], strings.Repeat("*", n-4), username[n-1:])
	} else {
		maskedUsername = ConcatenateStrings(username[:1], strings.Repeat("*", n-1))
	}

	return ConcatenateStrings(maskedUsername, "@", domain), nil
}

/*
Address is a struct that represents an address
This struct is used on FormatAddress function to format the address if following format: Street, Number - Complement - Neighborhood - City - State - PostalCode
*/
type Address struct {
	Street       string
	Number       string
	PostalCode   string
	Neighborhood string
	City         string
	State        string
	Complement   string
	Country      string
}

/*
FormatAddress is responsible for formatting the address

address: is the address to be formatted

return: the formatted address in the following format: Street, Number - Complement - Neighborhood - City - State - PostalCode
*/
func FormatAddress(address Address) string {
	var sb strings.Builder
	if StringIsNotEmpty(address.Street) {
		sb.WriteString(address.Street)
	}
	if StringIsNotEmpty(address.Number) {
		sb.WriteString(", ")
		sb.WriteString(address.Number)
	}
	if StringIsNotEmpty(address.Complement) {
		sb.WriteString(" - ")
		sb.WriteString(address.Complement)
	}
	if StringIsNotEmpty(address.Neighborhood) {
		sb.WriteString(" - ")
		sb.WriteString(address.Neighborhood)
	}
	if StringIsNotEmpty(address.City) {
		sb.WriteString(" - ")
		sb.WriteString(address.City)
	}
	if StringIsNotEmpty(address.State) {
		sb.WriteString(" - ")
		sb.WriteString(address.State)
	}
	if StringIsNotEmpty(address.Country) {
		sb.WriteString(" - ")
		sb.WriteString(address.Country)
	}
	if StringIsNotEmpty(address.PostalCode) {
		sb.WriteString(" - ")
		sb.WriteString(address.PostalCode)
	}
	return sb.String()
}

/*
GetCountryAcronymFromAreaCode is responsible for return the country acronym based on the phone area code

phoneAreaCode: is the phone area code to get the country acronym
*/
func GetCountryAcronymFromAreaCode(phoneAreaCode int) string {
	codes := map[int]string{
		1:   enums.NAPP_COMMONS_COUNTRY_AREA_CODE_US,
		33:  enums.NAPP_COMMONS_COUNTRY_AREA_CODE_FR,
		34:  enums.NAPP_COMMONS_COUNTRY_AREA_CODE_ES,
		39:  enums.NAPP_COMMONS_COUNTRY_AREA_CODE_IT,
		44:  enums.NAPP_COMMONS_COUNTRY_AREA_CODE_GB,
		49:  enums.NAPP_COMMONS_COUNTRY_AREA_CODE_DE,
		54:  enums.NAPP_COMMONS_COUNTRY_AREA_CODE_AR,
		55:  enums.NAPP_COMMONS_COUNTRY_AREA_CODE_BR,
		81:  enums.NAPP_COMMONS_COUNTRY_AREA_CODE_JP,
		86:  enums.NAPP_COMMONS_COUNTRY_AREA_CODE_CN,
		351: enums.NAPP_COMMONS_COUNTRY_AREA_CODE_PT,
	}

	return codes[phoneAreaCode]
}

/*
IsNil is responsible for check generic struct is nil

data: is the struct to check

return: bool if nil or not
*/
func IsNil(data any) bool {
	if data == nil {
		return true
	}

	switch reflect.TypeOf(data).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(data).IsNil()
	}

	return false
}

/*
IsNotNil is responsible for check generic struct is not nil

data: is the struct to check

return: bool if nil or not
*/
func IsNotNil(data any) bool {
	return !IsNil(data)
}

/*
ValueOrZero is responsible for return value or zero if nil of pointer

pointer: is the pointer to check

return: value or zero
*/
func ValueOrZero[T any](pointer *T) T {
	if pointer == nil {
		var v T
		return v
	}
	return *pointer
}

/*
AddTimezoneInTime is responsible add timezone in time and not change time

date: is time to add timezone

timezone: is timezone to add in time

return: time with timezone or error if timezone is invalid
*/
func AddTimezoneInTime(date time.Time, timezone string) (newDate time.Time, err error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return newDate, ConcatenateErrorStrings(enums.NAPP_COMMONS_ADD_TIMEZONE_IN_TIME_ERROR_MESSAGE, err.Error())
	}

	newDate = time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), location)

	return
}

/*
ValidateEan is responsible for validating the ean and returning the type of ean
The possible types of ean of return are: EAN_8, EAN_13
*/
func ValidateEan(ean string) (isValid bool, typeEan string) {
	if len(ean) < 8 || len(ean) > 14 {
		return false, ""
	}

	lengthIs13 := len(ean) == 13

	sumEven := 0
	sumOdd := 0

	for i := 0; i < len(ean)-1; i++ {
		number, err := strconv.Atoi(string(ean[i]))

		if err != nil {
			return false, ""
		}

		if i%2 == 0 {
			sumEven += number
		} else {
			sumOdd += number
		}
	}

	if lengthIs13 {
		sumOdd = sumOdd * 3
	} else {
		sumEven = sumEven * 3
	}

	result := sumOdd + sumEven
	verifyingDigit := (10 - (result % 10)) % 10

	digit, err := strconv.Atoi(string(ean[len(ean)-1]))

	if err != nil {
		return false, ""
	}

	if verifyingDigit == digit {
		if lengthIs13 {
			return true, enums.NAPP_COMMONS_EAN_TYPE_13
		}

		return true, enums.NAPP_COMMONS_EAN_TYPE_8
	}

	return false, ""
}

/*
RemoveSpecialCharacters is responsible for removing any special character, keeping only letters and numbers - alphanumeric

input: is the string with all characters

return: is the input string with only letters and numbers - alphanumeric
*/
func RemoveSpecialCharactersFromString(input string) string {
	regex := regexp.MustCompile("[^a-zA-Z0-9]")

	return regex.ReplaceAllString(input, "")
}

/*
FormatDocumentByType is responsible for formatting a document according to its type

document: is the characters of the document. The value must be alphanumeric

documentType: is the type of the document. Possible values are in enums NAPP_COMMONS_DOCUMENT_TYPE_*

return: formatted document according to its type
*/
func FormatDocumentByType(document, documentType string) string {
	switch documentType {
	case enums.NAPP_COMMONS_DOCUMENT_TYPE_CNPJ:
		if len(document) == 14 {
			return ConcatenateStrings(document[:2], ".", document[2:5], ".", document[5:8], "/", document[8:12], "-", document[12:])
		}
	}

	return document
}

/*
This function is responsible to verify if contains a expected int64 in a list of int64s

integersList: is the list of int64s that will be verified

expectedInteger: is the expected int64 to be verified
*/
func ContainsInt64(integersList []int64, expectedInteger int64) bool {
	for _, item := range integersList {
		if item == expectedInteger {
			return true
		}
	}
	return false
}

/*
This function is responsible to return the acronym of the state

state: state name
*/
func GetAcronymByState(state string) string {
	statesMap := map[string]string{
		"ACRE":                "AC",
		"ALAGOAS":             "AL",
		"AMAPÁ":               "AP",
		"AMAZONAS":            "AM",
		"BAHIA":               "BA",
		"CEARÁ":               "CE",
		"DISTRITO FEDERAL":    "DF",
		"ESPÍRITO SANTO":      "ES",
		"GOIÁS":               "GO",
		"MARANHÃO":            "MA",
		"MATO GROSSO":         "MT",
		"MATO GROSSO DO SUL":  "MS",
		"MINAS GERAIS":        "MG",
		"PARÁ":                "PA",
		"PARAÍBA":             "PB",
		"PARANÁ":              "PR",
		"PERNAMBUCO":          "PE",
		"PIAUÍ":               "PI",
		"RIO DE JANEIRO":      "RJ",
		"RIO GRANDE DO NORTE": "RN",
		"RIO GRANDE DO SUL":   "RS",
		"RONDÔNIA":            "RO",
		"RORAIMA":             "RR",
		"SANTA CATARINA":      "SC",
		"SÃO PAULO":           "SP",
		"SERGIPE":             "SE",
		"TOCANTINS":           "TO",
	}

	state = strings.ToUpper(state)

	return statesMap[state]
}

/*
This function is responsible to normalize string removing spelling accents

value: is the value that be normalize
*/
func NormalizeString(value string) (result string, err error) {
	chain := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, err = transform.String(chain, value)

	return
}

/*
This function is responsible get state from state acronym by country acronym

stateAcronym: state acronym
countryAcronym: country acronym
*/
func GetStateFromStateAcronymByCountryAcronym(stateAcronym, countryAcronym string) string {
	return enums.COUNTRIES_STATES[countryAcronym][stateAcronym]
}

/*
This function is responsible to validate if is a empty uuid

value: uuid to validate
*/
func UUIDIsNil(value uuid.UUID) bool {
	return value == uuid.Nil || StringIsEmpty(value.String())
}

/*
This function is responsible to validate if is not a empty uuid

value: uuid to validate
*/
func UUIDIsNotNil(value uuid.UUID) bool {
	return !UUIDIsNil(value)
}

/*
This function is responsible to validade if a NullUUID is valid and different than nil

value: NullUUID to validate
*/
func NullUUIDIsValid(value uuid.NullUUID) bool {
	if value.Valid && UUIDIsNotNil(value.UUID) {
		return true
	}

	return false
}

/*
This function is responsible to validade if a NullUUID is invalid or equal nil

value: NullUUID to validate
*/
func NullUUIDIsNotValid(value uuid.NullUUID) bool {
	return !NullUUIDIsValid(value)
}

/*
This function is responsible to count special characters

value: string to count special characters
expectedQuantity: quantity expected of special characters
*/
func ValidateSpecialCharacters(value string, expectedQuantity int) bool {
	pattern := `[!@#$%^&*()_+\-=\[\]{}|;':",./<>?]`

	re := regexp.MustCompile(pattern)

	specialChars := re.FindAllString(value, -1)

	return len(specialChars) >= expectedQuantity
}

/*
This function is responsible to count upper characters

value: string to count upper characters
quantityExpected: quantity expected of upper characters
*/
func ValidateUpperCharacters(value string, expectedQuantity int) bool {
	countUppers := 0
	for _, char := range value {
		if unicode.IsUpper(char) {
			countUppers++
			if countUppers >= expectedQuantity {
				break
			}
		}
	}

	return countUppers == expectedQuantity
}

/*
This function is responsible to removes all duplicate elements from a list

list: list of elements

returns: list with only unique elements
*/
func Distinct[T any](list []T) (result []T) {
	seen := make(map[any]bool)

	for _, value := range list {
		if !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}

	return
}
