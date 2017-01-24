package parseutil

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const project = "gopkg.in/src-d/go-parse-utils.v1"

var projectPath = filepath.Join(goSrc, project)

func projectFile(path string) string {
	return filepath.Join(projectPath, path)
}

func TestGetSourceFiles(t *testing.T) {
	_, paths, err := NewImporter().getSourceFiles(project, goPath)
	require.Nil(t, err)
	expected := []string{
		projectFile("ast.go"),
		projectFile("importer.go"),
	}

	require.Equal(t, expected, paths)
}

func TestParseSourceFiles(t *testing.T) {
	paths := []string{
		projectFile("ast.go"),
		projectFile("importer.go"),
	}

	pkg, err := NewImporter().parseSourceFiles(projectPath, paths)
	require.Nil(t, err)

	require.Equal(t, "parseutil", pkg.Name())
}

func TestImport(t *testing.T) {
	imp := NewImporter()
	pkg, err := imp.ImportFrom("srcd.works/go-parse-utils.v1", goSrc, 0)
	require.Nil(t, err)
	require.Equal(t, "parseutil", pkg.Name())
}
