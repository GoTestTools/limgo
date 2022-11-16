package domain_test

import (
	"path"
	"testing"

	"github.com/GoTestTools/limgo/pkg/domain"
	"github.com/GoTestTools/limgo/pkg/model/gosrc"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

//nolint:dupl,funlen
func TestExploreFunctions(t *testing.T) {

	testcases := []struct {
		name       string
		file       string
		expectFuns []gosrc.Function
		expectErr  bool
	}{
		{
			name:       "Returns error on non-existing file",
			file:       "does-not-exist",
			expectFuns: nil,
			expectErr:  true,
		},
		{
			name:       "Returns error on file containing only package header",
			file:       "empty_file_go",
			expectFuns: nil,
			expectErr:  true,
		},
		{
			name: "Successfully analyzes hello world example",
			file: "hello_world_go",
			expectFuns: []gosrc.Function{
				{
					Name:     "main",
					Position: gosrc.Position{LineFrom: 7, LineTo: 9, ColumnFrom: 1, ColumnTo: 2},
					Statements: []gosrc.Statement{
						{
							Type:             gosrc.ExprTypeExpr,
							Position:         gosrc.Position{LineFrom: 8, LineTo: 8, ColumnFrom: 2, ColumnTo: 28},
							NestedStatements: nil,
						},
					},
				},
			},
			expectErr: false,
		},
		{
			name: "Successfully analyzes hello world example with if statement",
			file: "hello_world_with_if_go",
			expectFuns: []gosrc.Function{
				{
					Name:     "main",
					Position: gosrc.Position{LineFrom: 8, LineTo: 12, ColumnFrom: 1, ColumnTo: 2},
					Statements: []gosrc.Statement{
						{
							Type:     gosrc.ExprTypeIf,
							Position: gosrc.Position{LineFrom: 9, LineTo: 11, ColumnFrom: 2, ColumnTo: 3},
							NestedStatements: []gosrc.Statement{
								{
									Type:             gosrc.ExprTypeExpr,
									Position:         gosrc.Position{LineFrom: 10, LineTo: 10, ColumnFrom: 3, ColumnTo: 40},
									NestedStatements: nil,
								},
							},
						},
					},
				},
			},
			expectErr: false,
		},
		{
			name: "Successfully analyzes hello world example with if-else statement",
			file: "hello_world_with_if_else_go",
			expectFuns: []gosrc.Function{
				{
					Name:     "main",
					Position: gosrc.Position{LineFrom: 8, LineTo: 14, ColumnFrom: 1, ColumnTo: 2},
					Statements: []gosrc.Statement{
						{
							Type:     gosrc.ExprTypeIf,
							Position: gosrc.Position{LineFrom: 9, LineTo: 13, ColumnFrom: 2, ColumnTo: 3},
							NestedStatements: []gosrc.Statement{
								{
									Type:     gosrc.ExprTypeExpr,
									Position: gosrc.Position{LineFrom: 10, LineTo: 10, ColumnFrom: 3, ColumnTo: 40},
								},
								{
									Type:     gosrc.ExprTypeBlock,
									Position: gosrc.Position{LineFrom: 11, LineTo: 13, ColumnFrom: 9, ColumnTo: 3},
									NestedStatements: []gosrc.Statement{
										{
											Type: gosrc.ExprTypeExpr, Position: gosrc.Position{LineFrom: 12, LineTo: 12, ColumnFrom: 3, ColumnTo: 34},
										},
									},
								},
							},
						},
					},
				},
			},
			expectErr: false,
		},
		{
			name: "Successfully analyzes hello world example with nested if statements",
			file: "hello_world_with_nested_if_go",
			expectFuns: []gosrc.Function{
				{
					Name:     "main",
					Position: gosrc.Position{LineFrom: 8, LineTo: 27, ColumnFrom: 1, ColumnTo: 2},
					Statements: []gosrc.Statement{
						{
							Type:     gosrc.ExprTypeIf,
							Position: gosrc.Position{LineFrom: 9, LineTo: 26, ColumnFrom: 2, ColumnTo: 3},
							NestedStatements: []gosrc.Statement{
								{
									Type:     gosrc.ExprTypeIf,
									Position: gosrc.Position{LineFrom: 10, LineTo: 24, ColumnFrom: 3, ColumnTo: 4},
									NestedStatements: []gosrc.Statement{
										{
											Type:     gosrc.ExprTypeExpr,
											Position: gosrc.Position{LineFrom: 11, LineTo: 11, ColumnFrom: 4, ColumnTo: 41},
										},
										{
											Type:     gosrc.ExprTypeBlock,
											Position: gosrc.Position{LineFrom: 12, LineTo: 24, ColumnFrom: 10, ColumnTo: 4},
											NestedStatements: []gosrc.Statement{
												{
													Type:     gosrc.ExprTypeExpr,
													Position: gosrc.Position{LineFrom: 13, LineTo: 13, ColumnFrom: 4, ColumnTo: 46},
												},
												{
													Type:     gosrc.ExprTypeAssign,
													Position: gosrc.Position{LineFrom: 14, LineTo: 14, ColumnFrom: 4, ColumnTo: 23},
												},
												{
													Type:     gosrc.ExprTypeIf,
													Position: gosrc.Position{LineFrom: 15, LineTo: 23, ColumnFrom: 4, ColumnTo: 5},
													NestedStatements: []gosrc.Statement{
														{
															Type:     gosrc.ExprTypeExpr,
															Position: gosrc.Position{LineFrom: 16, LineTo: 16, ColumnFrom: 5, ColumnTo: 55},
														},
														{
															Type:     gosrc.ExprTypeIf,
															Position: gosrc.Position{LineFrom: 17, LineTo: 23, ColumnFrom: 11, ColumnTo: 5},
															NestedStatements: []gosrc.Statement{
																{
																	Type:     gosrc.ExprTypeExpr,
																	Position: gosrc.Position{LineFrom: 18, LineTo: 18, ColumnFrom: 5, ColumnTo: 44},
																},
																{
																	Type:     gosrc.ExprTypeIf,
																	Position: gosrc.Position{LineFrom: 19, LineTo: 23, ColumnFrom: 11, ColumnTo: 5},
																	NestedStatements: []gosrc.Statement{
																		{
																			Type:     gosrc.ExprTypeExpr,
																			Position: gosrc.Position{LineFrom: 20, LineTo: 20, ColumnFrom: 5, ColumnTo: 44},
																		},
																		{
																			Type:     gosrc.ExprTypeBlock,
																			Position: gosrc.Position{LineFrom: 21, LineTo: 23, ColumnFrom: 11, ColumnTo: 5},
																			NestedStatements: []gosrc.Statement{
																				{
																					Type: gosrc.ExprTypeExpr, Position: gosrc.Position{LineFrom: 22, LineTo: 22, ColumnFrom: 5, ColumnTo: 46},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								{
									Type:     gosrc.ExprTypeExpr,
									Position: gosrc.Position{LineFrom: 25, LineTo: 25, ColumnFrom: 3, ColumnTo: 40},
								},
							},
						},
					},
				},
			},
			expectErr: false,
		},
		{
			name: "Successfully analyzes example with for statement",
			file: "for_go",
			expectFuns: []gosrc.Function{
				{
					Name:     "main",
					Position: gosrc.Position{LineFrom: 7, LineTo: 12, ColumnFrom: 1, ColumnTo: 2},
					Statements: []gosrc.Statement{
						{
							Type:     gosrc.ExprTypeFor,
							Position: gosrc.Position{LineFrom: 8, LineTo: 11, ColumnFrom: 2, ColumnTo: 3},
							NestedStatements: []gosrc.Statement{
								{
									Type:     gosrc.ExprTypeAssign,
									Position: gosrc.Position{LineFrom: 9, LineTo: 9, ColumnFrom: 3, ColumnTo: 46},
								},
								{
									Type:     gosrc.ExprTypeExpr,
									Position: gosrc.Position{LineFrom: 10, LineTo: 10, ColumnFrom: 3, ColumnTo: 19},
								},
							},
						},
					},
				},
			},
			expectErr: false,
		},
		{
			name: "Successfully analyzes example with range statement",
			file: "range_go",
			expectFuns: []gosrc.Function{
				{
					Name:     "main",
					Position: gosrc.Position{LineFrom: 7, LineTo: 12, ColumnFrom: 1, ColumnTo: 2},
					Statements: []gosrc.Statement{
						{
							Type:     gosrc.ExprTypeAssign,
							Position: gosrc.Position{LineFrom: 8, LineTo: 8, ColumnFrom: 2, ColumnTo: 54},
						},
						{
							Type:     gosrc.ExprTypeRange,
							Position: gosrc.Position{LineFrom: 9, LineTo: 11, ColumnFrom: 2, ColumnTo: 3},
							NestedStatements: []gosrc.Statement{
								{
									Type:     gosrc.ExprTypeExpr,
									Position: gosrc.Position{LineFrom: 10, LineTo: 10, ColumnFrom: 3, ColumnTo: 19},
								},
							},
						},
					},
				},
			},
			expectErr: false,
		},
		{
			name: "Successfully analyzes example with type switch",
			file: "type_switch_go",
			expectFuns: []gosrc.Function{
				{
					Name:     "main",
					Position: gosrc.Position{LineFrom: 7, LineTo: 20, ColumnFrom: 1, ColumnTo: 2},
					Statements: []gosrc.Statement{
						{
							Type:     gosrc.ExprTypeDeclaration,
							Position: gosrc.Position{LineFrom: 8, LineTo: 8, ColumnFrom: 2, ColumnTo: 13},
						},
						{
							Type:     gosrc.ExprTypeTypeSwitch,
							Position: gosrc.Position{LineFrom: 9, LineTo: 19, ColumnFrom: 2, ColumnTo: 3},
							NestedStatements: []gosrc.Statement{
								{
									Type:     gosrc.ExprTypeCaseClause,
									Position: gosrc.Position{LineFrom: 10, LineTo: 14, ColumnFrom: 2, ColumnTo: 4},
									NestedStatements: []gosrc.Statement{
										{
											Type:     gosrc.ExprTypeExpr,
											Position: gosrc.Position{LineFrom: 11, LineTo: 11, ColumnFrom: 3, ColumnTo: 17},
										},
										{
											Type:     gosrc.ExprTypeIf,
											Position: gosrc.Position{LineFrom: 12, LineTo: 14, ColumnFrom: 3, ColumnTo: 4},
											NestedStatements: []gosrc.Statement{
												{
													Type:     gosrc.ExprTypeExpr,
													Position: gosrc.Position{LineFrom: 13, LineTo: 13, ColumnFrom: 4, ColumnTo: 34},
												},
											},
										},
									},
								},
								{
									Type:     gosrc.ExprTypeCaseClause,
									Position: gosrc.Position{LineFrom: 15, LineTo: 16, ColumnFrom: 2, ColumnTo: 17},
									NestedStatements: []gosrc.Statement{
										{
											Type:     gosrc.ExprTypeExpr,
											Position: gosrc.Position{LineFrom: 16, LineTo: 16, ColumnFrom: 3, ColumnTo: 17},
										},
									},
								},
								{
									Type:     gosrc.ExprTypeCaseClause,
									Position: gosrc.Position{LineFrom: 17, LineTo: 18, ColumnFrom: 2, ColumnTo: 17},
									NestedStatements: []gosrc.Statement{
										{
											Type:     gosrc.ExprTypeExpr,
											Position: gosrc.Position{LineFrom: 18, LineTo: 18, ColumnFrom: 3, ColumnTo: 17},
										},
									},
								},
							},
						},
					},
				},
			},
			expectErr: false,
		},
		{
			name: "Successfully analyzes example with switch-case",
			file: "switch_go",
			expectFuns: []gosrc.Function{
				{
					Name:     "main",
					Position: gosrc.Position{LineFrom: 7, LineTo: 20, ColumnFrom: 1, ColumnTo: 2},
					Statements: []gosrc.Statement{
						{
							Type:     gosrc.ExprTypeDeclaration,
							Position: gosrc.Position{LineFrom: 8, LineTo: 8, ColumnFrom: 2, ColumnTo: 13},
						},
						{
							Type:     gosrc.ExprTypeSwitch,
							Position: gosrc.Position{LineFrom: 9, LineTo: 19, ColumnFrom: 2, ColumnTo: 3},
							NestedStatements: []gosrc.Statement{
								{
									Type:     gosrc.ExprTypeCaseClause,
									Position: gosrc.Position{LineFrom: 10, LineTo: 14, ColumnFrom: 2, ColumnTo: 4},
									NestedStatements: []gosrc.Statement{
										{
											Type:     gosrc.ExprTypeExpr,
											Position: gosrc.Position{LineFrom: 11, LineTo: 11, ColumnFrom: 3, ColumnTo: 17},
										},
										{
											Type:     gosrc.ExprTypeIf,
											Position: gosrc.Position{LineFrom: 12, LineTo: 14, ColumnFrom: 3, ColumnTo: 4},
											NestedStatements: []gosrc.Statement{
												{
													Type:     gosrc.ExprTypeExpr,
													Position: gosrc.Position{LineFrom: 13, LineTo: 13, ColumnFrom: 4, ColumnTo: 34},
												},
											},
										},
									},
								},
								{
									Type:     gosrc.ExprTypeCaseClause,
									Position: gosrc.Position{LineFrom: 15, LineTo: 16, ColumnFrom: 2, ColumnTo: 17},
									NestedStatements: []gosrc.Statement{
										{
											Type:     gosrc.ExprTypeExpr,
											Position: gosrc.Position{LineFrom: 16, LineTo: 16, ColumnFrom: 3, ColumnTo: 17},
										},
									},
								},
								{
									Type:     gosrc.ExprTypeCaseClause,
									Position: gosrc.Position{LineFrom: 17, LineTo: 18, ColumnFrom: 2, ColumnTo: 17},
									NestedStatements: []gosrc.Statement{
										{
											Type:     gosrc.ExprTypeExpr,
											Position: gosrc.Position{LineFrom: 18, LineTo: 18, ColumnFrom: 3, ColumnTo: 17},
										},
									},
								},
							},
						},
					},
				},
			},
			expectErr: false,
		},
		{
			name: "Successfully analyzes example with defer",
			file: "defer_go",
			expectFuns: []gosrc.Function{
				{
					Name:     "main",
					Position: gosrc.Position{LineFrom: 7, LineTo: 14, ColumnFrom: 1, ColumnTo: 2},
					Statements: []gosrc.Statement{
						{
							Type:     gosrc.ExprTypeDefer,
							Position: gosrc.Position{LineFrom: 8, LineTo: 13, ColumnFrom: 2, ColumnTo: 5},
							NestedStatements: []gosrc.Statement{
								{
									Type:     gosrc.ExprTypeExpr,
									Position: gosrc.Position{LineFrom: 9, LineTo: 9, ColumnFrom: 3, ColumnTo: 39},
								},
								{
									Type:     gosrc.ExprTypeIf,
									Position: gosrc.Position{LineFrom: 10, LineTo: 12, ColumnFrom: 3, ColumnTo: 4},
									NestedStatements: []gosrc.Statement{
										{
											Type:     gosrc.ExprTypeExpr,
											Position: gosrc.Position{LineFrom: 11, LineTo: 11, ColumnFrom: 4, ColumnTo: 43},
										},
									},
								},
							},
						},
					},
				},
			},
			expectErr: false,
		},
		{
			name: "Successfully analyzes example with func declaration",
			file: "func_go",
			expectFuns: []gosrc.Function{
				{
					Name:     "main",
					Position: gosrc.Position{LineFrom: 7, LineTo: 12, ColumnFrom: 1, ColumnTo: 2},
					Statements: []gosrc.Statement{
						{
							Type:     gosrc.ExprTypeAssign,
							Position: gosrc.Position{LineFrom: 8, LineTo: 10, ColumnFrom: 2, ColumnTo: 3},
							NestedStatements: []gosrc.Statement{
								{
									Type:     gosrc.ExprTypeExpr,
									Position: gosrc.Position{LineFrom: 9, LineTo: 9, ColumnFrom: 3, ColumnTo: 19},
								},
							},
						},
						{
							Type:     gosrc.ExprTypeExpr,
							Position: gosrc.Position{LineFrom: 11, LineTo: 11, ColumnFrom: 2, ColumnTo: 10},
						},
					},
				},
				{
					Name:     "print",
					Position: gosrc.Position{LineFrom: 14, LineTo: 16, ColumnFrom: 1, ColumnTo: 2},
					Statements: []gosrc.Statement{
						{
							Type:     gosrc.ExprTypeExpr,
							Position: gosrc.Position{LineFrom: 15, LineTo: 15, ColumnFrom: 2, ColumnTo: 22},
						},
					},
				},
			},
			expectErr: false,
		},
		{
			name: "someother",
			file: "func_decl_go",
			expectFuns: []gosrc.Function{
				{
					Name:     "main",
					Position: gosrc.Position{LineFrom: 7, LineTo: 11, ColumnFrom: 1, ColumnTo: 2},
					Statements: []gosrc.Statement{
						{
							Type:     gosrc.ExprTypeExpr,
							Position: gosrc.Position{LineFrom: 8, LineTo: 8, ColumnFrom: 2, ColumnTo: 28},
						},
						{
							Type:     gosrc.ExprTypeExpr,
							Position: gosrc.Position{LineFrom: 9, LineTo: 9, ColumnFrom: 2, ColumnTo: 9},
						},
						{
							Type:     gosrc.ExprTypeExpr,
							Position: gosrc.Position{LineFrom: 10, LineTo: 10, ColumnFrom: 2, ColumnTo: 13},
						},
					},
				},
				{
					Name:     "DoSth",
					Position: gosrc.Position{LineFrom: 13, LineTo: 15, ColumnFrom: 24, ColumnTo: 2},
					Statements: []gosrc.Statement{
						{
							Type:     gosrc.ExprTypeExpr,
							Position: gosrc.Position{LineFrom: 14, LineTo: 14, ColumnFrom: 2, ColumnTo: 47},
						},
					},
				},
				{
					Name:     "DoSthElse",
					Position: gosrc.Position{LineFrom: 15, LineTo: 17, ColumnFrom: 4, ColumnTo: 2},
					Statements: []gosrc.Statement{
						{
							Type:     gosrc.ExprTypeExpr,
							Position: gosrc.Position{LineFrom: 16, LineTo: 16, ColumnFrom: 2, ColumnTo: 51},
						},
					},
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			filePath := path.Join("./testdata", "goModuleAnalyzer", testcase.file)
			funcs, err := domain.ExploreFunctions(filePath)

			if err != nil && !testcase.expectErr {
				t.Fatalf("Unexpected error occurred while analyzing file %s: %v", filePath, err)
			}
			if err == nil && testcase.expectErr {
				t.Fatalf("Expected error, but none did occurr while analyzing file %s", filePath)
			}

			if funcs == nil && testcase.expectFuns != nil {
				t.Fatal("Expected detected gosrc.Functions, but got none")
			}
			if funcs != nil && testcase.expectFuns == nil {
				t.Fatalf("Did not expect gosrc.Functions to be detected, but got %d", len(funcs))
			}

			options := cmpopts.IgnoreFields(gosrc.Statement{}, "ParentStatement")
			if diff := cmp.Diff(testcase.expectFuns, funcs, options); diff != "" {
				t.Fatalf("Detected diff in analyzed file %s: \n%s", filePath, diff)
			}
		})
	}
}
