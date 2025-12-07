package usecases

import "net/http"

type UseCaseHandler interface {
	Handle(request *http.Request) (status int, data any)
}
