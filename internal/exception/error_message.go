package exception

import "errors"

var (
	GeneralExceptionMessage        = errors.New("Please wait a few minutes before you try again.")
	AWSExceptionMessage            = errors.New("Error connecting to AWS")
	EnvExceptionMessage            = errors.New("Error get environment variables")
	CallApiExceptionMessage        = errors.New("Error Call API, please wait a few minutes before you try again.")
	JSONParseExceptionMessage      = errors.New("Error while parsing JSON, please wait a few minutes before you try again.")
	ClientDoExceptionMessage       = errors.New("Error Execute API, please wait a few minutes before you try again.")
	RedisExceptionMessage          = errors.New("Redis Error")
	TokenExceptionMessage          = errors.New("Session expired")
	JSONParseRedisExceptionMessage = errors.New("Error while parsing JSON redis, please wait a few minutes before you try again.")
)
