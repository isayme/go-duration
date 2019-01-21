package duration

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDurationUnmarshal(t *testing.T) {
	require := require.New(t)

	var d Duration
	err := json.Unmarshal([]byte(`"1h2m3s"`), &d)
	require.Nil(err)

	v := time.Duration(d)
	require.Equal(v.String(), "1h2m3s")
}

func TestDurationMarshalText(t *testing.T) {
	require := require.New(t)

	var d Duration
	err := json.Unmarshal([]byte(`"1h2m3s"`), &d)
	require.Nil(err)

	data, err := d.MarshalText()
	require.Nil(err)
	require.Equal(string(data), `1h2m3s`)

	data, err = json.Marshal(d)
	require.Nil(err)
	require.Equal(string(data), `"1h2m3s"`)
}

func TestDurationPrint(t *testing.T) {
	require := require.New(t)

	var d Duration
	err := json.Unmarshal([]byte(`"1h2m3s"`), &d)
	require.Nil(err)

	require.Equal(`1h2m3s`, fmt.Sprintf("%v", d))
	require.Equal(`1h2m3s`, fmt.Sprintf("%+v", d))
	require.Equal(`"1h2m3s"`, fmt.Sprintf("%#v", d))
}
