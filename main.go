package main

   import (
	   "context"
	   "flag"
	   "fmt"
	   "log"
	   "time"

	   "github.com/chromedp/chromedp"
   )

   // Mapping of language codes for --list-langs flag
   var languageCodes = map[string]string{
    "af": "Afrikaans",
    "sq": "Albanian",
    "am": "Amharic",
    "ar": "Arabic",
    "hy": "Armenian",
    "az": "Azerbaijani",
    "eu": "Basque",
    "be": "Belarusian",
    "bn": "Bengali",
    "bs": "Bosnian",
    "bg": "Bulgarian",
    "ca": "Catalan",
    "ceb": "Cebuano",
    "ny": "Chichewa",
    "zh-CN": "Chinese (Simplified)",
    "zh-TW": "Chinese (Traditional)",
    "co": "Corsican",
    "hr": "Croatian",
    "cs": "Czech",
    "da": "Danish",
    "nl": "Dutch",
    "en": "English",
    "eo": "Esperanto",
    "et": "Estonian",
    "tl": "Filipino",
    "fi": "Finnish",
    "fr": "French",
    "fy": "Frisian",
    "gl": "Galician",
    "ka": "Georgian",
    "de": "German",
    "el": "Greek",
    "gu": "Gujarati",
    "ht": "Haitian Creole",
    "ha": "Hausa",
    "haw": "Hawaiian",
    "iw": "Hebrew",
    "hi": "Hindi",
    "hmn": "Hmong",
    "hu": "Hungarian",
    "is": "Icelandic",
    "ig": "Igbo",
    "id": "Indonesian",
    "ga": "Irish",
    "it": "Italian",
    "ja": "Japanese",
    "jw": "Javanese",
    "kn": "Kannada",
    "kk": "Kazakh",
    "km": "Khmer",
    "ko": "Korean",
    "ku": "Kurdish (Kurmanji)",
    "ky": "Kyrgyz",
    "lo": "Lao",
    "la": "Latin",
    "lv": "Latvian",
    "lt": "Lithuanian",
    "lb": "Luxembourgish",
    "mk": "Macedonian",
    "mg": "Malagasy",
    "ms": "Malay",
    "ml": "Malayalam",
    "mt": "Maltese",
    "mi": "Maori",
    "mr": "Marathi",
    "mn": "Mongolian",
    "my": "Myanmar (Burmese)",
    "ne": "Nepali",
    "no": "Norwegian",
    "or": "Odia (Oriya)",
    "ps": "Pashto",
    "fa": "Persian",
    "pl": "Polish",
    "pt": "Portuguese",
    "pa": "Punjabi",
    "ro": "Romanian",
    "ru": "Russian",
    "sm": "Samoan",
    "gd": "Scots Gaelic",
    "sr": "Serbian",
    "st": "Sesotho",
    "sn": "Shona",
    "sd": "Sindhi",
    "si": "Sinhala",
    "sk": "Slovak",
    "sl": "Slovenian",
    "so": "Somali",
    "es": "Spanish",
    "su": "Sundanese",
    "sw": "Swahili",
    "sv": "Swedish",
    "tg": "Tajik",
    "ta": "Tamil",
    "tt": "Tatar",
    "te": "Telugu",
    "th": "Thai",
    "tr": "Turkish",
    "tk": "Turkmen",
    "uk": "Ukrainian",
    "ur": "Urdu",
    "ug": "Uyghur",
    "uz": "Uzbek",
    "vi": "Vietnamese",
    "cy": "Welsh",
    "xh": "Xhosa",
    "yi": "Yiddish",
    "yo": "Yoruba",
    "zu": "Zulu",
}


func main() {
	// Define command-line arguments
	srcLang := flag.String("il", "auto", "Source language code (default: auto)")
	outLang := flag.String("ol", "", "Target language code (required)")
	text := flag.String("text", "", "Text to translate (required)")
	listLangsFlag := flag.Bool("list-langs", false, "List available language codes and exit")
	
	// Parse flags
	flag.Parse()

	// Handle --list-langs flag
	if *listLangsFlag {
		listLanguages()
		return
	}

	// Validate the input arguments after flag parsing
	if *outLang == "" {
		log.Fatal("Error: Target language (-ol) is required.")
	}

	if *text == "" {
		log.Fatal("Error: Text to translate is required (use -text).")
	}

	// Call the translation function with chromedp
	doTranslate(*srcLang, *outLang, *text)
}

// listLanguages prints available languages
func listLanguages() {
	fmt.Println("Available Languages:")
	for code, name := range languageCodes {
		fmt.Printf("%s: %s\n", code, name)
	}
}

// doTranslate handles Google Translate interaction using chromedp
func doTranslate(srcLang, outLang, text string) {
	// Build the Google Translate URL
	translateURL := fmt.Sprintf("https://translate.google.com/?sl=%s&tl=%s&text=%s&op=translate", srcLang, outLang, text)
	fmt.Printf("Requesting: %s\n", translateURL)

	// Create context with timeout
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Set a timeout for the entire process
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	var translatedText string

	// Run chromedp tasks
	err := chromedp.Run(ctx,
		chromedp.Navigate(translateURL),
		chromedp.WaitReady("body"), // Wait until the translation element is visible
		chromedp.Text(`span.ryNqvb`, &translatedText, chromedp.ByQuery), // Select the translated text
	)

	if err != nil {
		log.Fatalf("Failed to extract translated text: %v\n", err)
	}

	// Output the translation to the terminal
	fmt.Printf("Translated Text: %s\n", translatedText)
   }