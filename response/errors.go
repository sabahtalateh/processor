package response

import (
	"fmt"
	"net/http"
)

var (
	defaultOnMarshalError = func(err error, _ http.ResponseWriter, _ any) { fmt.Printf("%s", err) }
	defaultOnWriteError   = func(written int, err error, body []byte) { fmt.Printf("%s", err) }

	onMarshalError = defaultOnMarshalError
	onWriteError   = defaultOnWriteError
)

func UseOnMarshalError(f func(err error, w http.ResponseWriter, body any)) { onMarshalError = f }
func UseOnWriteError(f func(written int, err error, body []byte))          { onWriteError = f }
