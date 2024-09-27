package main

   import (
	   "fmt"
      "os"
      "sort"
      "github.com/sss7526/gwerd/internal/cli"
      "github.com/sss7526/gwerd/internal/processor"
      "github.com/sss7526/gwerd/internal/constants"
      "github.com/sss7526/gwerd/internal/file_handler"
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

   blockMode, ok := parsedArgs["block"].(bool)
   if !ok || !blockMode {
      blockMode = false
   }
   
   srcLang, _ := parsedArgs["source-lang"].(string)
   outLang, _ := parsedArgs["output-lang"].(string)
   
   filepath, _ := parsedArgs["file"].(string)

   var text []string

   if filepath != "" {
      absPath, err := file_handler.ResolveFilePath(filepath)
      if err != nil {
         fmt.Printf("Error parsing file: %v\n", err)
         os.Exit(1)
      }
      file, err := os.Open(absPath)
      if err  != nil {
         fmt.Printf("Error opening file: %v\n", err)
         os.Exit(1)
      }
      if !blockMode {
         text, err = file_handler.ReadLines(file)
      } else {
         text, err = file_handler.ReadBlock(file)
      }
      if err != nil {
         fmt.Printf("%v", err)
         os.Exit(1)
      }
      defer file.Close()

   }

   if filepath == "" {
      text, ok = parsedArgs["text"].([]string)
      if !ok || len(text) == 0 {
         text = nil
      }
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