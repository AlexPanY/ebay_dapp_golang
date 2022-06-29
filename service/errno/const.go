package errno

var (
	Success                = &Err{Code: "Success", Message: "Success"}
	ErrUnauthorized        = &Err{Code: "Unauthorized", Message: "Unauthorized account."}
	ErrMissingToken        = &Err{Code: "MissingToken", Message: "Request `X-Auth-Token` is missing in Header."}
	ErrTokenInvalid        = &Err{Code: "InvalidToken", Message: "Invalid X-Auth-Token."}
	ErrParamInvalid        = &Err{Code: "InvalidParameter", Message: "Invalid parameter."}
	ErrParamTrimSpace      = &Err{Code: "TrimSpaceParameter", Message: "TrimSpace parameter invalid."}
	ErrParamValidationF    = &Err{Code: "MissingParameter", Message: "Parameter %s must be required."}
	ErrInternalServerError = &Err{Code: "ServiceUnavailable", Message: "Internal server error."}
	ErrDatabase            = &Err{Code: "ServiceException", Message: "Database operation exception."}
	ErrLookupHost          = &Err{Code: "LookupHostException", Message: "Exception for a given host."}
	ErrJSONException       = &Err{Code: "JsonException", Message: "JSON data parsing failed."}
)
