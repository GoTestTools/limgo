package evaluation

type CoverageError struct {
	Type         CoverageErrorType
	AffectedFile string

	ExpectedThreshold float64
	ActualCovered     float64
}
