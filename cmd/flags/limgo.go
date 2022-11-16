package flags

import "flag"

const (
	FlagCoverFile  string = "coverfile"
	FlagOutFile    string = "outfile"
	FlagOutFormat  string = "outfmt"
	FlagConfigFile string = "config"
	FlagVerbosity  string = "v"
)

type Flags struct {
	CoverageFile string
	OutputFile   string
	OutputFormat string
	ConfigFile   string
	Verbosity    uint
}

func ParseFlags() Flags {
	coverFile := flag.String(FlagCoverFile, "coverage.cov",
		`Coverage file. Required. File name, including the path, of the coverage results from go test.`,
	)

	configFile := flag.String(FlagConfigFile, ".limgo.json",
		`Threshold configuration file. File name, including the path, of the coverage results from go test.`,
	)

	outFile := flag.String(FlagOutFile, "",
		`Output file. Name of the file to write the coverage results to. If left blank, output will be printed to stdout.`,
	)

	outFormat := flag.String(FlagOutFormat, "tab",
		`Output format. Supported values: tab, json, md.`,
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
		CoverageFile: *coverFile,
		OutputFile:   *outFile,
		OutputFormat: *outFormat,
		ConfigFile:   *configFile,
		Verbosity:    *verbosity,
	}
}
