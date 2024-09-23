package errlist

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

// If we do not set custom errors in golang,
// we can not get any errors which passed from service
var (
	ErrBadRequest            = errors.New("bad request")
	ErrBadQueryParams        = errors.New("bad query params")
	ErrUnauthorized          = errors.New("unauthorized")
	ErrForbidden             = errors.New("forbidden")
	ErrFieldValidation       = errors.New("field validation error")
	ErrNotFound              = errors.New("not found")
	ErrNoRecord              = errors.New("no records found")
	ErrMethodNotAllowed      = errors.New("method not allowed")
	ErrNotAcceptable         = errors.New("not acceptable")
	ErrRequestTimeOut        = errors.New("request timeout")
	ErrConflict              = errors.New("conflict")
	ErrGone                  = errors.New("gone")
	ErrLengthRequired        = errors.New("length required")
	ErrTooManyRequests       = errors.New("too many request")
	ErrInternalServer        = errors.New("internal server error")
	ErrServiceUnavailable    = errors.New("service unavailable")
	ErrNoCookie              = errors.New("no cookie")
	ErrPermissionDenied      = errors.New("permission denied")
	ErrNoSuchUser            = errors.New("no such user")
	ErrEmailAlreadyExists    = errors.New("email already exist")
	ErrTransactionFailed     = errors.New("failed to perform transaction")
	ErrInvalidJWTToken       = errors.New("invalid JWT Token")
	ErrInvalidJWTClaims      = errors.New("invalid JWT Claims")
	ErrNotAllowedImageHeader = errors.New("not allowed image header")
	ErrNotRequiredFields     = errors.New("missing required fields")
	ErrRange                 = errors.New("value out of range")
	ErrSyntax                = errors.New("invalid syntax")
	ErrBucketNotFound        = errors.New("bucket not found")
	ErrObjectNotFound        = errors.New("object not found")
	ErrNoSuchFile            = errors.New("http: no such file")
)

// RestErr is the interface for custom error handling
type RestErr interface {
	Status() int
	Error() string // note that err method in go implements Error() string method, so we must use that method
	Causes() interface{}
}

// RestError struct to implement the RestErr interface
type RestError struct {
	ErrStatus  int         `json:"err_status,omitempty"`
	ErrMessage string      `json:"err_msg,omitempty"`
	ErrCauses  interface{} `json:"err_cause,omitempty"`
}

// Status is
func (e RestError) Status() int {
	return e.ErrStatus
}

// Causes is
func (e RestError) Causes() interface{} {
	return e.ErrCauses
}

// Message is
func (e RestError) Error() string {
	return fmt.Sprintf("err_status: %d - err_msg: %s - err_cause: %v", e.ErrStatus, e.ErrMessage, e.ErrCauses)
}

// NewRestError is
func NewRestError(status int, err_msg string, causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  status,
		ErrMessage: err_msg,
		ErrCauses:  causes,
	}
}

// NewBadRequestError creates a new 400 Bad Request error
func NewBadRequestError(causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: ErrBadRequest.Error(),
		ErrCauses:  causes,
	}
}

// NewNotFoundError creates a new 404 Not Found error
func NewNotFoundError(causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusNotFound,
		ErrMessage: ErrNotFound.Error(),
		ErrCauses:  causes,
	}
}

// NewUnAuthorizedError creates a new 401 Unauthorized error
func NewUnAuthorizedError(causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusUnauthorized,
		ErrMessage: ErrUnauthorized.Error(),
		ErrCauses:  causes,
	}
}

// NewForbiddenError creates a new 403 Forbidden error
func NewForbiddenError(causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusForbidden,
		ErrMessage: ErrForbidden.Error(),
		ErrCauses:  causes,
	}
}

// NewInternalServerError creates a new 500 Internal Server Error
func NewInternalServerError(causes interface{}) RestErr {
	result := RestError{
		ErrStatus:  http.StatusInternalServerError,
		ErrMessage: ErrInternalServer.Error(),
		ErrCauses:  causes,
	}
	return result
}

// NewRequestTimedOutError creates a new 408 Request Timeout error
func NewRequestTimedOutError(causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusRequestTimeout,
		ErrMessage: ErrRequestTimeOut.Error(),
		ErrCauses:  causes,
	}
}

// NewConflictError creates a new 409 Conflict error
func NewConflictError(causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusConflict,
		ErrMessage: ErrConflict.Error(),
		ErrCauses:  causes,
	}
}

// NewTooManyRequestError creates a new 429 TooManyRequest error
func NewTooManyRequestError(causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusTooManyRequests,
		ErrMessage: ErrTooManyRequests.Error(),
		ErrCauses:  causes,
	}
}

// NewBadQueryParamsError creates a new 400 Bad Request error for bad query parameters
func NewBadQueryParamsError(causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: ErrBadQueryParams.Error(),
		ErrCauses:  causes,
	}
}

// Parsing errors with switch case
func ParseErrors(err error) RestErr {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewNotFoundError(err.Error())
	case errors.Is(err, ErrTooManyRequests):
		return NewTooManyRequestError(err.Error())
	case errors.Is(err, context.DeadlineExceeded):
		return NewRequestTimedOutError(err.Error())
	// Handle AWS errors
	case strings.Contains(err.Error(), ErrBucketNotFound.Error()),
		strings.Contains(err.Error(), ErrObjectNotFound.Error()):
		return NewBadRequestError(err.Error())

	// Handle SQLSTATE errors
	case strings.Contains(err.Error(), "SQLSTATE"), strings.Contains(err.Error(), "1062"):
		return ParseSqlErrors(err)

	// Handle strconv.Atoi errors
	case strings.Contains(err.Error(), ErrSyntax.Error()):
		return NewBadRequestError(err.Error())
	case strings.Contains(err.Error(), ErrRange.Error()):
		return NewBadRequestError(err.Error())

	// Handle Validation errors from go-validator
	case errors.As(err, &validator.ValidationErrors{}):
		return ParseValidatorError(err)

	// Handle Token or Cookie errors
	case strings.Contains(strings.ToLower(err.Error()), ErrNoCookie.Error()),
		strings.Contains(strings.ToLower(err.Error()), ErrInvalidJWTToken.Error()),
		strings.Contains(strings.ToLower(err.Error()), ErrInvalidJWTClaims.Error()):
		return NewUnAuthorizedError(ErrUnauthorized.Error() + err.Error())

	default:
		// If already a RestErr, return as-is
		if restErr, ok := err.(RestErr); ok {
			return restErr
		}
		// For any other error, return an internal server error
		return NewInternalServerError(err.Error())
	}
}

// ParseSQL errors is
func ParseSqlErrors(err error) RestErr {
	// Handle sql.ErrNoRows
	if errors.Is(err, sql.ErrNoRows) {
		// No rows found, return a 404 Not Found error
		return NewNotFoundError(ErrNoRecord.Error() + err.Error())
	}

	// Handle MySQL-specific errors by checking SQLSTATE codes and error numbers
	// Error numbers can be found here: https://dev.mysql.com/doc/refman/8.0/en/server-error-reference.html
	switch {
	// MySQL error 1062: Duplicate entry (unique constraint violation)
	case strings.Contains(err.Error(), "1062"):
		return NewConflictError("duplicate entry (unique constraint violated): " + err.Error())

	// MySQL error 1452: Cannot add or update a child row (foreign key constraint violation)
	case strings.Contains(err.Error(), "1452"):
		return NewBadRequestError("foreign key constraint violated: " + err.Error())

	// MySQL error 1048: Column cannot be null
	case strings.Contains(err.Error(), "1048"):
		return NewBadRequestError("column cannot be null: " + err.Error())

	// MySQL error 1054: Unknown column
	case strings.Contains(err.Error(), "1054"):
		return NewBadRequestError("unknown column in field list: " + err.Error())

	// MySQL error 1366: Incorrect string value (invalid character encoding, etc.)
	case strings.Contains(err.Error(), "1366"):
		return NewBadRequestError("incorrect string value: " + err.Error())

	// MySQL error 1146: Table doesn't exist
	case strings.Contains(err.Error(), "1146"):
		return NewInternalServerError("table does not exist: " + err.Error())

	// MySQL error 1264: Out of range value for column
	case strings.Contains(err.Error(), "1264"):
		return NewBadRequestError("value out of range for column: " + err.Error())

	// MySQL error 1406: Data too long for column
	case strings.Contains(err.Error(), "1406"):
		return NewBadRequestError("data too long for column: " + err.Error())
	}

	// Handle SQLX-specific errors (if necessary)
	if strings.Contains(err.Error(), "sqlx") {
		// Handle SQLX-specific errors (such as transaction errors or other SQLX-related failures)
		// For example, sqlx may return errors with "sqlx: ...", you can customize this section further if needed
		return NewInternalServerError("SQLX operation failed: " + err.Error())
	}

	// If no specific error handling is matched, return a generic internal server error
	return NewInternalServerError(err.Error())
}

// parseValidatorError parses validation errors and returns corresponding RestErr
func ParseValidatorError(err error) RestErr {
	validationErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return NewBadRequestError(err.Error()) // If not a validation error, fallback to generic error
	}

	// Collect detailed validation error messages
	var errorMessages []string
	for _, fieldErr := range validationErrs {
		// For each validation error, create a message
		errorMessage := fmt.Sprintf("Field Validation for %s failed on the %s tag", fieldErr.Field(), fieldErr.Tag())

		// Append the message to the error list
		errorMessages = append(errorMessages, errorMessage)
	}

	// Combine all messages into one string and return a Bad Request Error
	return NewBadRequestError(strings.Join(errorMessages, ", "))
}

// Response returns is ErrorResponse, for clean syntax I took function name Response
// Because in every package i call this errlst package httpError, serviceErr, repoErr
// Then easily call this httpError.Response(err), serviceErr.Response(err), repoErr.Response(err)
func Response(err error) error {
	return ParseErrors(err)
}
