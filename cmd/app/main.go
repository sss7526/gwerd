package main

   import (
	   "fmt"
      "os"
      "sort"
      "github.com/sss7526/gwerd/internal/cli"
      "github.com/sss7526/gwerd/internal/processor"
      "github.com/sss7526/gwerd/internal/constants"
   )

func main() {
   parsedArgs := cli.ParseArgs()
   
   listLangs, ok := parsedArgs["list"].(bool)
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
	type kv struct {
      key string
      Value string
   }
   
   var sorted []kv
   for k, v := range constants.LanguageCodes {
      sorted = append(sorted, kv{k, v})
   }

   sort.Slice(sorted, func(i, j int) bool {
      return sorted[i].Value < sorted[j].Value
   })


   fmt.Println("Available Languages:")

   for _, kv := range sorted {
      fmt.Printf("%s: %s\n", kv.key, kv.Value)
   }

}