package evaluation

type CoverageErrorType string

const (
	CoverageErrorStmt CoverageErrorType = "statements"
	CoverageErrorFunc CoverageErrorType = "functions"
)

type CoverageError struct {
	Type         CoverageErrorType
	AffectedFile string

	ExpectedThreshold float64
	ActualCovered     float64
}
