package netconfig

import (
	"bytes"
	"testing"

	"github.com/go-kit/log"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {

	buf := &bytes.Buffer{}
	logger := log.NewLogfmtLogger(buf)

	nc, err := New(Config{}, logger)
	require.Error(t, err)
	require.Nil(t, nc)
}
