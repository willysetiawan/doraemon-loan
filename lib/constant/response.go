package constant

const (
	// code rc
	CODE_SUCCESS     = "200"
	CODE_SUCCESS_MSG = "SUCCESS"

	CODE_CREATED     = "201"
	CODE_CREATED_MSG = "CREATE"

	CODE_REJECT = "400"
	// CODE_REJECT_MSG = "Bad Request"
	CODE_REJECT_MSG = "FAILED"

	CODE_FAILED     = "400"
	CODE_FAILED_MSG = "FAILED"

	CODE_NOT_FOUND_REJECT     = "404"
	CODE_NOT_FOUND_REJECT_MSG = "Not Found"

	CODE_INTERNAL_SERVER_ERR = "500"
	// CODE_INTERNAL_SERVER_ERR_MSG = "Internal Server Error"
	CODE_INTERNAL_SERVER_ERR_MSG = "FAILED"

	CODE_PENDING     = "504"
	CODE_PENDING_MSG = "PENDING"

	EXPCRED_CODE_SUCCESS  = "EC-200"
	PROCOTP_CODE_SUCCESS  = "OTP-200"
	PROCOTP_CODE_EXPIRED  = "OTP-464"
	PROCOTP_CODE_CREATED  = "OTP-201"
	PROCCRED_CODE_SUCCESS = "PC-200"

	MESSAGE_SUCCESS                        = "Successful"
	MESSAGE_NOT_FOUND                      = "Transaction not found"
	MESSAGE_INVALID_TOKEN                  = "Invalid Token"
	MESSAGE_UNAUTHORIZED                   = "Unauthorized"
	MESSAGE_INSUFFICIENT_FUND              = "Insufficient Fund"
	MESSAGE_BACKEND_SYSTEM_FAILURE         = "Backend system failure"
	MESSAGE_ACCESS_TOKEN_INVALID           = "Access token invalid"
	MESSAGE_PAYMENT_TOKEN_NOT_FOUND        = "Token not found in the system"
	MESSAGE_INACTIVE_ACCOUNT               = "Account or User status is abnormal"
	MESSAGE_ACCOUNT_INFORMATION_INVALID    = "Account information invalid"
	MESSAGE_NOT_BINDING                    = "Account not binding yet"
	MESSAGE_TRANSACTION_NOT_FOUND          = "Transaction not found"
	MESSAGE_INVALID_FIELD_FORMAT           = "Invalid field format"
	MESSAGE_TRANSACTION_EXPIRED            = "Transaction expired"
	MESSAGE_PAYMENT_SUCCESS                = "Payment success"
	MESSAGE_PAYMENT_FAILED                 = "Payment failed"
	MESSAGE_INVALID_MERCHANT               = "Invalid merchant"
	MESSAGE_INCONSISTENT_REQUEST           = "Inconsistent Request"
	MESSAGE_CONFLICT                       = "Conflict"
	MESSAGE_DUPLICATE_PARTNER_REFERENCE_NO = "Duplicate partnerReferenceNo"
)
