package gosrc_test

import (
	"testing"

	"github.com/GoTestTools/limgo/pkg/model/gosrc"
)

func TestAnalyzedFileDirectory(t *testing.T) {
	file := gosrc.AnalyzedFile{
		FileName: "json.go",
		FilePath: "pkg/dto/json.go",
	}

	expected := "pkg/dto"
	dir := file.Directory()
	if dir != expected {
		t.Fatalf("Expected directory to be '%s', but got '%s'", expected, dir)
	}
}
