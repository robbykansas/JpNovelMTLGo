package exception

type DataNotFoundError struct {
	Message string
}

type GeneralError struct {
	Message string
}

type GetSessionError struct {
	Message string
}

type UnauthorizedError struct {
	Message string
}

type InternalServerError struct {
	Message string
}

func (dataNotFoundError DataNotFoundError) Error() string {
	return dataNotFoundError.Message
}

func (generalError GeneralError) Error() string {
	return generalError.Message
}

func (getSessionError GetSessionError) Error() string {
	return getSessionError.Message
}

func (unauthorizedError UnauthorizedError) Error() string {
	return unauthorizedError.Message
}

func (internalServerError InternalServerError) Error() string {
	return internalServerError.Message
}
