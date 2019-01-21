package duration

import (
	"strings"
	"time"
)

// Duration parse time.Duration from string
type Duration time.Duration

// UnmarshalText for json.Unmarshal
func (d *Duration) UnmarshalText(data []byte) (err error) {
	s := string(data)

	t, err := time.ParseDuration(s)
	if err != nil {
		return
	}

	*d = Duration(t)
	return nil
}

// MarshalText ...
func (d Duration) MarshalText() ([]byte, error) {
	s := d.String()
	return []byte(s), nil
}

// String for %v
func (d Duration) String() string {
	v := time.Duration(d)
	return v.String()
}

// GoString for %#v
func (d Duration) GoString() string {
	return strings.Join([]string{"\"", "\""}, d.String())
}
