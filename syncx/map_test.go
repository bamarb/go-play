package syncx

import (
	"go-play/ortb/model"
	"testing"

	"github.com/puzpuzpuz/xsync"
	"github.com/stretchr/testify/require"
)

func TestConcStrMap(t *testing.T) {
	strmap := xsync.NewMapOf[model.Fcap]()
	require.NotNil(t, strmap)
}
