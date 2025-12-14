package http

import "net/http"

type WebHandlerFunc func(request *http.Request) (status int, data any)
