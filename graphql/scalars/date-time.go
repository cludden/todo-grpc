package scalars

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalDateTime custom scalar marshaler
func MarshalDateTime(d time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		fmt.Fprintf(w, `"%s"`, d.Format(time.RFC3339Nano))
	})
}

// UnmarshalDateTime custom scalar unmarshal
func UnmarshalDateTime(v interface{}) (time.Time, error) {
	str, ok := v.(string)
	if !ok {
		fmt.Println("DateTime values must be strings")
		return time.Time{}, errors.New("DateTime values must be strings")
	}

	t, err := time.Parse(time.RFC3339Nano, str)
	if err != nil {
		fmt.Printf("invalid DateTime string: %v", err)
		return t, fmt.Errorf("invalid DateTime string: %v", err)
	}
	return t, nil
}
