package flags

import "flag"

const Some string = "asdf"

const (
	FlagCovFile    string = "covfile"
	FlagConfigFile string = "config"
	FlagVerbosity  string = "v"
)

type Flags struct {
	CoverageFile string
	ConfigFile   string
	Verbosity    uint
}

func ParseFlags() Flags {
	covFile := flag.String(FlagCovFile, "",
		`Coverage file. Required. File name, including the path, of the coverage results from go test.`,
	)

	configFile := flag.String(FlagConfigFile, ".limgo.json",
		`Coverage file. File name, including the path, of the coverage results from go test.`,
	)

	verbosity := flag.Uint(FlagVerbosity, 2,
		`Verbosity level. Defines the depth for the output of the coverage statistic. 
0 = No statistic output.  
1 = Output statistic of global coverage.  
2 = Output statistic of global and directory coverage. 
3 = Output statistic of global, directory and file coverage. 
3 > Output statistic of global, directory, file and function coverage.
`,
	)

	flag.Parse()

	return Flags{
		CoverageFile: *covFile,
		ConfigFile:   *configFile,
		Verbosity:    *verbosity,
	}
}
