package evaluation

type CoverageErrorType string

const (
	CoverageErrorStmt     CoverageErrorType = "statements"
	CoverageErrorLines    CoverageErrorType = "lines"
	CoverageErrorBranches CoverageErrorType = "branches"
	CoverageErrorFunc     CoverageErrorType = "functions"
)
