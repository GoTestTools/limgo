package gosrc_test

import (
	"fmt"
	"testing"

	"github.com/GoTestTools/limgo/pkg/model/gosrc"
)

//nolint:funlen
func TestStatementIsBranchStmt(t *testing.T) {

	testcases := []struct {
		stmt           gosrc.Statement
		expectIsBranch bool
	}{
		{
			stmt: gosrc.Statement{
				Type: gosrc.ExprTypeBlock,
			},
			expectIsBranch: false,
		},
		{
			stmt: gosrc.Statement{
				Type: gosrc.ExprTypeIf,
			},
			expectIsBranch: true,
		},
		{
			stmt: gosrc.Statement{
				Type: gosrc.ExprTypeSwitch,
			},
			expectIsBranch: true,
		},
		{
			stmt: gosrc.Statement{
				Type: gosrc.ExprTypeTypeSwitch,
			},
			expectIsBranch: true,
		},
		{
			stmt: gosrc.Statement{
				Type: gosrc.ExprTypeCaseClause,
			},
			expectIsBranch: true,
		},
		{
			stmt: gosrc.Statement{
				Type: gosrc.ExprTypeRange,
			},
			expectIsBranch: true,
		},
		{
			stmt: gosrc.Statement{
				Type: gosrc.ExprTypeFor,
			},
			expectIsBranch: true,
		},
		{
			stmt: gosrc.Statement{
				Type: gosrc.ExprTypeAssign,
			},
			expectIsBranch: false,
		},
		{
			stmt: gosrc.Statement{
				Type: gosrc.ExprTypeDeclaration,
			},
			expectIsBranch: false,
		},
		{
			stmt: gosrc.Statement{
				Type: gosrc.ExprTypeDefer,
			},
			expectIsBranch: false,
		},
		{
			stmt: gosrc.Statement{
				Type: gosrc.ExprTypeExpr,
			},
			expectIsBranch: false,
		},
		{
			stmt: gosrc.Statement{
				Type: gosrc.ExprTypeGenDecl,
			},
			expectIsBranch: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("IsBranch for stmt type %s", testcase.stmt.Type), func(t *testing.T) {
			got := testcase.stmt.IsBranchStmt()
			if testcase.expectIsBranch != got {
				t.Fatalf("Expected statement %s on IsBranchStmt to return %t, but got %t",
					testcase.stmt.Type, testcase.expectIsBranch, got)
			}
		})
	}
}
