package http

import "net/http"

type Handler interface {
	Handle(request *http.Request) (status int, data any)
}
