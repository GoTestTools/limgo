package gosrc_test

import (
	"testing"

	"github.com/GoTestTools/limgo/pkg/model/gosrc"
)

func TestStatementsPopEmpty(t *testing.T) {
	stmts := gosrc.Statements{}
	_, err := stmts.Pop()
	if err == nil {
		t.Fatal("Expected error when executing Pop() on empty statement stack")
	}
}

func TestStatementsPush(t *testing.T) {
	stmts := gosrc.Statements{}
	stmts.Push(gosrc.Statements{{}, {}, {}})
	if len(stmts) != 3 {
		t.Fatal("Expected statement stack to contain 3 elements")
	}
}

func TestStatementsPushAndPop(t *testing.T) {
	stmts := gosrc.Statements{}
	stmts.Push(gosrc.Statements{{Type: "first"}, {Type: "second"}, {Type: "third"}})
	elem, err := stmts.Pop()
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v", err)
	}
	if elem.Type != "third" {
		t.Fatalf("Expected the last element put on the stack to be removed first")
	}
}
