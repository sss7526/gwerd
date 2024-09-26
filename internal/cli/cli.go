package cli

import (
	"fmt"
	"os"
	"github.com/sss7526/goparse"
)

func ParseArgs() (map[string]interface{}) {
	parser := goparse.NewParser(
		goparse.WithName("\nGWERD"),
		goparse.WithVersion("1.0.0\n"),
		goparse.WithDescription("CLI text translation utility.\n"),
	)

	parser.AddArgument("verbose", "v", "verbose", "Increase verbosity, shows http requests/responses and allowed/blocked status", "bool", false)
	parser.AddArgument("source-lang", "s", "source-lang", "Source language to translate from. Ex: -s en", "string", false)
	parser.AddArgument("output-lang", "o", "output-lang", "Target language to translate to. Ex: -o fr", "string", false)
	parser.AddArgument("text", "t", "text", "One or more strings (enclosed in double quotes) to translate. Ex: -t \"<your phrase>\"", "[]string", false)
	parser.AddArgument("engine", "e", "engine", "Translation engine to run against (Google, DeePL, Bing, etc). Ex: -e google", "string", false)
	parser.AddArgument("list", "l", "list-langs", "List available language codes", "bool", false)

	parser.AddExclusiveGroup([]string{"targets", "engine"}, false)

	parser.AddExclusiveGroup([]string{"list", "source-lang"}, true)
	parser.AddExclusiveGroup([]string{"list", "output-lang"}, true)
	parser.AddExclusiveGroup([]string{"list", "engine"}, true)

	parsedArgs, shouldExit, err := parser.Parse()
	if err != nil {
		fmt.Printf("Error parsing arguments: %v\n", err)
		if shouldExit {
			os.Exit(1)
		}
	}

	if shouldExit {
		os.Exit(0)
	}

	return parsedArgs
	
}

