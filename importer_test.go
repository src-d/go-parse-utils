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
	_, paths, err := NewImporter().getSourceFiles(project, goPath, FileFilters{})
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
	pkg, err := imp.Import(project)
	require.Nil(t, err)
	require.Equal(t, "parseutil", pkg.Name())
	require.NotNil(t, pkg.Scope().Lookup("FileType"), "finds an object in importer.go")
}

func TestImportWithFilters(t *testing.T) {
	imp := NewImporter()
	_, err := imp.ImportWithFilters(project, FileFilters{
		func(pkgPath, file string, typ FileType) bool {
			return file != "importer.go"
		},
	})
	require.NotNil(t, err, "excluding importer.go makes package unimportable")
}

func TestImportFrom(t *testing.T) {
	imp := NewImporter()
	pkg, err := imp.ImportFrom(project, goSrc, 0)
	require.Nil(t, err)
	require.Equal(t, "parseutil", pkg.Name())
}

func TestFileFilters(t *testing.T) {
	fs := FileFilters{
		func(pkgPath, file string, typ FileType) bool {
			return pkgPath == "a"
		},
		func(pkgPath, file string, typ FileType) bool {
			return file == "a"
		},
		func(pkgPath, file string, typ FileType) bool {
			return typ == GoFile
		},
	}

	require.True(t, fs.KeepFile("a", "a", GoFile))
	require.False(t, fs.KeepFile("b", "a", GoFile))
	require.False(t, fs.KeepFile("a", "b", GoFile))
	require.False(t, fs.KeepFile("a", "a", CgoFile))
}
