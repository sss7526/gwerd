package main

   import (
	   "fmt"
      "os"
      "github.com/sss7526/gwerd/internal/cli"
      "github.com/sss7526/gwerd/internal/processor"
      "github.com/sss7526/gwerd/internal/constants"
   )

func main() {
   parsedArgs := cli.ParseArgs()
   
   listLangs, ok := parsedArgs["langs"].(bool)
   if ok && listLangs {
      listLanguages()
      os.Exit(0)
   } 

   verbose, ok := parsedArgs["verbose"].(bool)
   if !ok || !verbose {
      verbose = false
   }
   
   srcLang, _ := parsedArgs["source-lang"].(string)
   outLang, _ := parsedArgs["output-lang"].(string)

   text, ok := parsedArgs["text"].([]string)
   if !ok || len(text) == 0 {
      text = nil
   }


	// Call the translation function with chromedp
	processor.Translate(srcLang, outLang, text, verbose)
}

// listLanguages prints available languages
func listLanguages() {
	fmt.Println("Available Languages:")
	for code, name := range constants.LanguageCodes {
		fmt.Printf("%s: %s\n", code, name)
	}
}