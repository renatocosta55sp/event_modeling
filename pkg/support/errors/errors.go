package errors

import (
	"github.com/jellydator/ttlcache/v3"
	log "github.com/sirupsen/logrus"
	"github.org/eventmodeling/ecommerce/pkg/support/commons"
)

/*
custom error structure with key, message and attributes.
*/
type CoreError struct {
	Key        string               `json:"key"`
	Message    string               `json:"message"`
	Attributes []CoreAttributeError `json:"attributes"`
}

type CoreAttributeError struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type CoreErrorField struct {
	Field   string `json:"field"`
	Key     string `json:"key"`
	Message string `json:"message"`
}

/*
Function responsible for get from cache the message, if error comes with more than one string then will concatenate to message.

key: eg.: "error.validate"
message: eg.: "validate error"
*/
func New(keys ...string) *CoreError {
	var cacheMsg *ttlcache.Item[string, string]
	msgKey := keys[0]
	var message string

	if len(keys) > 1 {
		for i := 1; i < len(keys); i++ {
			message = commons.ConcatenateStrings(message, keys[i])
		}
	}

	cacheMsg = C.Get(msgKey)

	if cacheMsg == nil {
		log.Errorf(commons.ConcatenateStrings("error not in errors.json please contact the dev team:", msgKey))
		return &CoreError{Key: msgKey}
	}

	if commons.StringIsNotEmpty(message) {
		fullMessageError := commons.ConcatenateStrings(cacheMsg.Value(), " ", message)
		return &CoreError{Key: cacheMsg.Key(), Message: fullMessageError}

	}
	return &CoreError{Key: cacheMsg.Key(), Message: cacheMsg.Value()}
}

func MakeErrorField(err error, field string, errorFields *[]CoreErrorField) (hasError bool) {
	if err == nil {
		return
	}

	errOut := ConvertTo(err)
	if errOut == nil {
		return
	}

	*errorFields = append(*errorFields, CoreErrorField{
		Field:   field,
		Key:     errOut.Key,
		Message: errOut.Message,
	})

	return len(*errorFields) > 0
}

/*
Function responsible for convert error to CoreError.
*/
func ConvertTo(err interface{}) *CoreError {
	errOut, ok := err.(*CoreError)
	if !ok {
		errDefault, _ := err.(error)
		errOut = New("error.unmapped", errDefault.Error())
	}

	return errOut
}

/*
Function responsible for adding AttributeError in Error.

field: eg.: "company_owner"
value: eg.: "company express"
*/
func (e *CoreError) AddAttributeError(field, value string) *CoreError {
	e.Attributes = append(e.Attributes, CoreAttributeError{Field: field, Value: value})

	return e
}

/*
Function responsible for concatenating key and message with | .

	Eg.: Error{
			key: "error.validate"
			message: "validate error"
		}

	result:	"error.validate | validate error"
*/
func (e *CoreError) Error() string {
	return commons.ConcatenateStrings(e.Key, " | ", e.Message)
}
