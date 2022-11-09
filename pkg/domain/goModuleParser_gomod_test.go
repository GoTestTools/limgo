package domain_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/GoTestTools/limgo/pkg/domain"
	"github.com/GoTestTools/limgo/pkg/model/module"
	"github.com/google/go-cmp/cmp"
)

func TestParseGoMod(t *testing.T) {

	testcases := []struct {
		file        string
		expectGoMod module.GoMod
		expectError bool
	}{
		{
			file:        "empty_file",
			expectError: true,
		},
		{
			file:        "missing_go_version",
			expectError: true,
		},
		{
			file:        "missing_module",
			expectError: true,
		},
		{
			file: "valid_gomod",
			expectGoMod: module.GoMod{
				ModuleName: "github.com/GoTestTools/limgo",
				GoVersion:  "1.19",
			},
			expectError: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("with %s", testcase.file), func(t *testing.T) {
			file, err := os.Open(path.Join("testdata", "goModuleParser_gomod", testcase.file))
			if err != nil {
				t.Fatalf("Unexpected error occurred when opening file '%s': %v", testcase.file, err)
			}

			goMod, err := domain.ParseGoMod(file)
			if testcase.expectError && err == nil {
				t.Fatalf("Expected error, but got none")
			}
			if !testcase.expectError && err != nil {
				t.Fatalf("Expected no error, but got %v", err)
			}

			if diff := cmp.Diff(testcase.expectGoMod, goMod); diff != "" {
				t.Fatalf("Detected difference in parsed GoMod: %s", diff)
			}
		})
	}
}
