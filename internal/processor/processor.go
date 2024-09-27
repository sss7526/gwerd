package processor

import (

	"fmt"
	"context"
	"log"
	"time"
	"strings"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/network"
)

func Translate(srcLang, outLang string, texts []string, verbose bool) {
	
	if outLang != "" {
		for _, text := range texts {
			err := doTranslate(srcLang, outLang, text, verbose)
			if err != nil {
				fmt.Printf("Error performing translation for text: %s\nError: %v", text, err)
			}
		}
	} else {
		fmt.Println("No output language specified.")
	}
}

// doTranslate handles Google Translate interaction using chromedp
func doTranslate(srcLang, outLang string, text string, verbose bool) error {
	keywordsToBlock := []string{"ads", "tracking", "analytics", "adservice", "counter", "track", "guestbook"}

	blockedURLS := []string{}
	for _, keyword := range keywordsToBlock {
		blockedURLS = append(blockedURLS, fmt.Sprintf("*%s*", keyword))
	}

	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebkit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
	referrer := "https://www.google.com"

	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoDefaultBrowserCheck,
		chromedp.NoFirstRun,
		chromedp.UserAgent(userAgent),
		chromedp.Flag("disable-application-cache", true),
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("ignore-certificate-errors", true),
	)

	// Build the Google Translate URL
	translateURL := fmt.Sprintf("https://translate.google.com/?sl=%s&tl=%s&text=%s&op=translate", srcLang, outLang, text)
	fmt.Printf("Requesting: %s\n", translateURL)

	// Create context with timeout
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		network.Enable(),
		network.SetBlockedURLS(blockedURLS),
	)
	if err != nil {
		return fmt.Errorf("failed to enable network events with blocked URLs: %w", err)
	}

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if verbose {
			switch ev := ev.(type) {
				
			case *network.EventRequestWillBeSent:
				shouldBlock := false
				badword := ""
				fmt.Printf("VALIDATING URL: %s\n\n", ev.Request.URL)
				for _, keyword := range keywordsToBlock {
					if strings.Contains(ev.Request.URL, keyword) {
						shouldBlock = true
						badword = keyword
						break
					}
				}

				if shouldBlock {
					fmt.Printf("BLOCKED Request: %s (contains '%s')\n\n", ev.Request.URL, badword)
				} else {
					fmt.Printf("ALLOWED Request URL: %s\n", ev.Request.URL)
					fmt.Printf("ALLOWED Request METHOD: %s\n", ev.Request.Method)
					fmt.Printf("ALLOWED Request HEADERS: %s\n\n", ev.Request.Headers)
				}

			case *network.EventResponseReceived:
				fmt.Printf("RESPONSE URL: %s\n", ev.Response.URL)
				fmt.Printf("RESPONSE STATUS: %d\n", ev.Response.Status)
				fmt.Printf("RESPONSE HEADERS: %s\n\n", ev.Response.Headers)
			}
		}
	})

	var translatedText string

	// Run chromedp tasks
	err = chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			headers := make(map[string]interface{})
			headers["referer"] = referrer

			return network.SetExtraHTTPHeaders(network.Headers(headers)).Do(ctx)
		}),
		chromedp.Navigate(translateURL),
		chromedp.WaitReady("body"), // Wait until the translation element is visible
		chromedp.Sleep(5 * time.Second),
		chromedp.Text(`span.ryNqvb`, &translatedText, chromedp.ByQuery), // Select the translated text
	)

	if err != nil {
		log.Fatalf("Failed to extract translated text: %v\n", err)
	}

	// Output the translation to the terminal
	fmt.Printf("%s\n", translatedText)

	return nil
}