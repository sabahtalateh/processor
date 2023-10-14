package response

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Response struct{ opts []Option }

type response struct {
	status    int
	headers   map[string][]string
	body      any
	marshalFn MarshalFn
}

func New(opts ...Option) Response {
	return Response{opts: opts}
}

func JSON(opts ...Option) Response {
	return Response{opts: append([]Option{BodyMarshalFn(json.Marshal), WithContentType("application/json")}, opts...)}
}

func HTML(opts ...Option) Response {
	return Response{opts: append([]Option{BodyMarshalFn(defaultMarshal), WithContentType("text/html")}, opts...)}
}

func XML(opts ...Option) Response {
	return Response{opts: append([]Option{BodyMarshalFn(xml.Marshal), WithContentType("application/xml")}, opts...)}
}

func (r *Response) Write(w http.ResponseWriter) {
	resp := response{status: -1, headers: map[string][]string{}}
	for _, opt := range r.opts {
		opt.apply(&resp)
	}

	if resp.status == -1 {
		resp.status = http.StatusOK
	}

	if resp.marshalFn == nil {
		resp.marshalFn = defaultMarshal
	}

	for h, hh := range resp.headers {
		for _, v := range hh {
			w.Header().Add(h, v)
		}
	}

	out, err := resp.marshalFn(resp.body)
	if err != nil {
		onMarshalError(err, w, resp.body)
		return
	}

	w.WriteHeader(resp.status)
	written, err := w.Write(out)
	if err != nil {
		onWriteError(written, err, out)
		return
	}
}
