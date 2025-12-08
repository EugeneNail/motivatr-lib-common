package http

import "net/http"

type HandlerFunc func(request *http.Request) (status int, data any)
