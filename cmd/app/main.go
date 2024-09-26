package main

   import (
	   // "context"
	   // "flag"
	   "fmt"
	   // "log"
	   // "time"
      "os"

	   // "github.com/chromedp/chromedp"
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

// // doTranslate handles Google Translate interaction using chromedp
// func doTranslate(srcLang, outLang string, text []string) {
// 	// Build the Google Translate URL
// 	translateURL := fmt.Sprintf("https://translate.google.com/?sl=%s&tl=%s&text=%s&op=translate", srcLang, outLang, text)
// 	fmt.Printf("Requesting: %s\n", translateURL)

// 	// Create context with timeout
// 	ctx, cancel := chromedp.NewContext(context.Background())
// 	defer cancel()

// 	// Set a timeout for the entire process
// 	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
// 	defer cancel()

// 	var translatedText string

// 	// Run chromedp tasks
// 	err := chromedp.Run(ctx,
// 		chromedp.Navigate(translateURL),
// 		chromedp.WaitReady("body"), // Wait until the translation element is visible
// 		chromedp.Text(`span.ryNqvb`, &translatedText, chromedp.ByQuery), // Select the translated text
// 	)

// 	if err != nil {
// 		log.Fatalf("Failed to extract translated text: %v\n", err)
// 	}

// 	// Output the translation to the terminal
// 	fmt.Printf("Translated Text: %s\n", translatedText)
//    }