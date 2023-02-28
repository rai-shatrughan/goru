package main

import (
	"fmt"
	"log"
	"time"

	"github.com/playwright-community/playwright-go"
)

var (
	browser playwright.Browser
	page    playwright.Page
	pw      *playwright.Playwright
	err     error
)

func main() {
	setup()
	defer teardown()
	gnews()
	cnn()
}

func setup() {
	err = playwright.Install()
	if err != nil {
		log.Fatalf("could not install playwright: %v", err)
	}
	pw, err = playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err = pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err = browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
}

func teardown() {
	if err := browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err := pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}

func cnn() {
	if _, err = page.Goto("https://edition.cnn.com/"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	time.Sleep(60 * time.Second)
	entries, err := page.QuerySelectorAll("xpath=//a/span")
	if err != nil {
		log.Fatalf("could not get entries: %v", err)
	}
	for i, entry := range entries {
		titleElement, err := entry.TextContent()
		if err != nil {
			log.Fatalf("could not get title element: %v", err)
		}
		fmt.Printf("%d: %s\n", i+1, titleElement)
	}

}

func gnews() {
	if _, err := page.Goto("https://news.google.co.in"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	time.Sleep(30 * time.Second)
	entries, err := page.QuerySelectorAll("xpath=//article//a")
	if err != nil {
		log.Fatalf("could not get entries: %v", err)
	}
	for i, entry := range entries {
		titleElement, err := entry.GetAttribute("aria-label")
		if err != nil {
			log.Fatalf("could not get title element: %v", err)
		}
		fmt.Printf("%d: %s\n", i+1, titleElement)
	}
}
