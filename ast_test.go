package parseutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPackageAST(t *testing.T) {
	pkg, err := PackageAST(project)
	require.Nil(t, err)
	require.Equal(t, "parseutil", pkg.Name)
}
