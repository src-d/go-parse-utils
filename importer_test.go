package parseutil_test

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/src-d/go-parse-utils.v1"

	"github.com/stretchr/testify/require"
	_ "google.golang.org/grpc"
)

const project = "gopkg.in/src-d/go-parse-utils.v1"

var (
	goPath = os.Getenv("GOPATH")
	goSrc  = filepath.Join(goPath, "src")
)

var projectPath = filepath.Join(goSrc, project)

func projectFile(path string) string {
	return filepath.Join(projectPath, path)
}

func TestGetSourceFiles(t *testing.T) {
	_, paths, err := parseutil.NewImporter().GetSourceFiles(project, goPath, parseutil.FileFilters{})
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

	pkg, err := parseutil.NewImporter().ParseSourceFiles(projectPath, paths)
	require.Nil(t, err)

	require.Equal(t, "parseutil", pkg.Name())
}

func TestImport(t *testing.T) {
	imp := parseutil.NewImporter()
	pkg, err := imp.Import(project)
	require.Nil(t, err)
	require.Equal(t, "parseutil", pkg.Name())
}

func TestImportWithFilters(t *testing.T) {
	imp := parseutil.NewImporter()
	_, err := imp.ImportWithFilters(project, parseutil.FileFilters{
		func(pkgPath, file string, typ parseutil.FileType) bool {
			return file != "importer.go"
		},
	})
	require.NotNil(t, err, "excluding importer.go makes package unimportable")
}

func TestImportGoogleGrpc(t *testing.T) {
	imp := parseutil.NewImporter()
	_, err := imp.Import("google.golang.org/grpc")
	require.Nil(t, err, "should be able to import this. Was a bug")
}

func TestImportFrom(t *testing.T) {
	imp := parseutil.NewImporter()
	pkg, err := imp.ImportFrom(project, goSrc, 0)
	require.Nil(t, err)
	require.Equal(t, "parseutil", pkg.Name())
}

func TestFileFilters(t *testing.T) {
	fs := parseutil.FileFilters{
		func(pkgPath, file string, typ parseutil.FileType) bool {
			return pkgPath == "a"
		},
		func(pkgPath, file string, typ parseutil.FileType) bool {
			return file == "a"
		},
		func(pkgPath, file string, typ parseutil.FileType) bool {
			return typ == parseutil.GoFile
		},
	}

	require.True(t, fs.KeepFile("a", "a", parseutil.GoFile))
	require.False(t, fs.KeepFile("b", "a", parseutil.GoFile))
	require.False(t, fs.KeepFile("a", "b", parseutil.GoFile))
	require.False(t, fs.KeepFile("a", "a", parseutil.CgoFile))
}
