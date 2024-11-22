package constants

type DefaultResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  []string    `json:"errors"`
}

type AuthData struct {
	ID    int64  `json:"id"` // user id
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

const (
	STATUS_SUCCESS        = "0000"
	STATUS_CREATE_SUCCESS = "0001"
	STATUS_READ_SUCCESS   = "0002"
	STATUS_UPDATE_SUCCESS = "0003"
	STATUS_DELETE_SUCCESS = "0004"
	STATUS_USER_FOUND     = "0006"
	STATUS_AUTH_SUCCESS   = "0007"
	STATUS_CRED_SENT      = "0010"
	STATUS_VALID_PROMO    = "0070"

	STATUS_FAILED                 = "5000"
	STATUS_CREATE_FAILED          = "5001"
	STATUS_READ_FAILED            = "5002"
	STATUS_UPDATE_FAILED          = "5003"
	STATUS_DELETE_FAILED          = "5004"
	STATUS_JSON_VALIDATION_FAILED = "5005"
	STATUS_USER_NOT_FOUND         = "5006"
	STATUS_DATA_NOT_FOUND         = "5007"
	STATUS_EMPTY_DATA             = "5008"
	STATUS_INVALID_IDENTIFIER     = "5009"
	STATUS_USER_ALREADY_EXIST     = "5010"
	STATUS_MULTIPLE_IDENTIFIER    = "5011"
	STATUS_INVALID_AUTHORIZATION  = "5020"
	STATUS_EXPIRED_TOKEN          = "5021"
	STATUS_FORBIDDEN              = "5030"
	STATUS_DB_ERROR               = "5040"
	STATUS_VENDOR_REQUEST_FAILED  = "5041"
	STATUS_JSON_UNMARSHAL_FAILED  = "5042"
	STATUS_INVALID_PROMO          = "5070"
)

const (
	MESSAGE_SUCCESS                = "Success"
	MESSAGE_STILL_PROCESS          = "Transaction is being process"
	MESSAGE_FAILED                 = "Something went wrong"
	MESSAGE_INVALID_REQUEST_FORMAT = "Invalid Request Format"
	MESSAGE_UNAUTHORIZED           = "Unauthorized"
	MESSAGE_FORBIDDEN              = "Forbidden"
	MESSAGE_CONFLICT               = "Conflict"
)

func GetCustomResponse(status, message string, data interface{}, errors []string) DefaultResponse {
	return DefaultResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Errors:  errors,
	}
}
