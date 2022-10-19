package evaluation

type CoverageErrorType string

const (
	CoverageErrorStmt     CoverageErrorType = "statements"
	CoverageErrorLines    CoverageErrorType = "lines"
	CoverageErrorBranches CoverageErrorType = "branches"
	CoverageErrorFunc     CoverageErrorType = "functions"
)

type CoverageError struct {
	Type         CoverageErrorType
	AffectedFile string

	ExpectedThreshold float64
	ActualCovered     float64
}
