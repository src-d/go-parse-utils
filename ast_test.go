package parseutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPackageAST(t *testing.T) {
	pkg, err := PackageAST("srcd.works/go-parse-utils.v1")
	require.Nil(t, err)
	require.Equal(t, "parseutil", pkg.Name)
}
