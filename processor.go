package processor

import (
	"net/http"

	"github.com/sabahtalateh/processor/response"
)

func HandlerFunc(f func(r *http.Request) response.Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := f(r)
		resp.Write(w)
	}
}
