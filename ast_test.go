package parseutil_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/src-d/go-parse-utils.v1"
)

func TestPackageAST(t *testing.T) {
	pkg, err := parseutil.PackageAST(project)
	require.Nil(t, err)
	require.Equal(t, "parseutil", pkg.Name)
}

func TestPackageTestAST(t *testing.T) {
	pkg, err := parseutil.PackageTestAST(project)
	require.Nil(t, err)
	require.Equal(t, "parseutil_test", pkg.Name)
}
