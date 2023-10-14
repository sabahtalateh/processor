package response

type Option interface {
	apply(*response)
}

type MarshalFn func(any) ([]byte, error)

type StatusOption struct{ status int }
type HeaderOption struct{ key, value string }
type BodyOption struct{ body any }
type BodyMarshalFnOption struct{ fn MarshalFn }

func Status(status int) StatusOption                 { return StatusOption{status: status} }
func Header(key, value string) HeaderOption          { return HeaderOption{key: key, value: value} }
func Body(body any) BodyOption                       { return BodyOption{body: body} }
func BodyMarshalFn(fn MarshalFn) BodyMarshalFnOption { return BodyMarshalFnOption{fn: fn} }

func (s StatusOption) apply(r *response) { r.status = s.status }

func (b BodyOption) apply(r *response) { r.body = b.body }

func (h HeaderOption) apply(r *response) {
	if _, ok := r.headers[h.key]; !ok {
		r.headers[h.key] = []string{}
	}
	r.headers[h.key] = append(r.headers[h.key], h.value)
}

func (m BodyMarshalFnOption) apply(r *response) {
	r.marshalFn = m.fn
}

func WithContentType(value string) HeaderOption {
	return Header("Content-Type", value)
}
