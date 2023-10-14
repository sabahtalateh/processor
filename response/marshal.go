package response

import (
	"fmt"
)

func defaultMarshal(v any) ([]byte, error) {
	if v == nil {
		return nil, nil
	}

	switch vv := v.(type) {
	case []byte:
		return vv, nil
	case string:
		return []byte(vv), nil
	default:
		return nil, fmt.Errorf("[]byte and string can be marshaled by default. use response.BodyMarshalFn for %T", v)
	}
}
