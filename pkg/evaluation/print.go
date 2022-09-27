package evaluation

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/engelmi/limgo/pkg"
)

func PrintPretty(output io.Writer, errs []CoverageError) {
	w := tabwriter.NewWriter(output, 1, 1, 1, ' ', 0)

	pkg.ColorizedFPrintln(w, "\nExpected coverage thresholds not met:", pkg.ColorRed)
	for _, err := range errs {
		txt := fmt.Sprintf(
			"   '%s':\texpected coverage threshold for %s of %.2f%%, but only got %.2f%%",
			err.AffectedFile, err.Type, err.ExpectedThreshold, err.ActualCovered,
		)
		pkg.ColorizedFPrintln(w, txt, pkg.ColorRed)

	}

	w.Flush()
}
