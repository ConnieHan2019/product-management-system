package error

import "errors"

var (
	SUCCESS = 200

	EMAIL_FORMAT_INVALID = 301

	PASSWORD_INVALID = 302

	VALID_CODE_ERROR = 303

	VALID_CODE_NOT_SENT = 304

	EMAIL_HAS_BEEN_USED = 305

	LOGIN_ERROR = 311

	EMAIL_NOT_EXIST = 321

	PASSWORD_ERROR = 322

	PASSWORD_REPEAT = 325

	FAIL_TO_STORE_FILE = 326

	FAIL_TO_GET_UPLOAD_FILE = 327

	FAIL_TO_CREATE_TASK = 363

	FAIL_TO_SUBMIT_TASK = 372

	CREATE_TASK_REQUSET_INVALID = 361

	REFERENCE_NOT_EXIST = 362

	TASK_ID_INALID = 371

	QUERY_PARAM_IVALID = 381

	FAIL_TO_GET_TASKS = 382

	FAIL_TO_GET_TASK_STATUS = 383

	INTERNAL_SERVER_ERROR = 500
)

var (
	ERROR_TOKEN_OUTDATED             = errors.New("token outdated")
	ERROR_TOKEN_NOT_PROVIDED         = errors.New("token not provided")
	ERROR_EMAIL_NOT_PROVIDED         = errors.New("email not provided")
	ERROR_EMAIL_HAS_BEEN_REGISTERED  = errors.New("email has been registered")
	ERROR_PASSWORD_NOT_PROVIDED      = errors.New("password not provided")
	ERROR_TASK_NAME_NOT_PROVIDED     = errors.New("task name not provided")
	ERROR_PASSWORD_NOT_CORRECT       = errors.New("password not correct")
	ERROR_VALIDATE_CODE_NOT_PROVIDED = errors.New("validate code not provided")
	ERROR_USER_NOT_EXIST             = errors.New("user not exist")
	ERROR_REFERENCE_NOT_EXIST        = errors.New("reference not exist")
	ERROR_VALIDATE_CODE              = errors.New("invalid error code")
	ERROR_PASSWORD_REPEAT            = errors.New("password repeat")
	ERROR_LOGIN                      = errors.New("email not exist or password not correct")
)
