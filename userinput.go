package tomtypes

import (
	"encoding/json"
	"errors"
	"strconv"
)

// StringInt can decode a number from a string, useful when accepting user input
// Forked from https://codesahara.com/blog/golang-cannot-unmarshal-string-into-go-value-of-type-int/
type StringInt int

func (st *StringInt) UnmarshalJSON(b []byte) error {
	// convert the bytes into an interface, to easily check the type
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*st = StringInt(v)
	case float64:
		*st = StringInt(int(v))
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			// the string might not be of integer type
			return err

		}
		*st = StringInt(i)
	default:
		return errors.New("value is not a string, float or int")
	}
	return nil
}
