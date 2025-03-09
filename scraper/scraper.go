package scraper

import (
	"fmt"
	"net/http"
	"news/logger"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	Category    string
	ArticleName string
	Summary     string
	Link        string
	TimePosted  string
}

func ScrapeNBC() ([]Article, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// GET request
	req, err := http.NewRequest("GET", "https://www.nbcnews.com/latest-stories/", nil)
	if err != nil {
		logger.Red(fmt.Sprintf("Error making GET request to URL: %v", err))
		return nil, fmt.Errorf("Error making GET request to URL: %w", err)
	}

	// Set headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")

	// Make request with headers
	res, err := client.Do(req)
	if err != nil {
		logger.Red(fmt.Sprintf("Error making request with headers: %v", err))
		return nil, fmt.Errorf("Error making request with headers: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		logger.Red(fmt.Sprintf("Error scraping data: %d %s", res.StatusCode, res.Status))
		return nil, fmt.Errorf("Error scraping data: %d %s", res.StatusCode, res.Status)
	}

	// Parse html data
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		logger.Red(fmt.Sprintf("error parsing HTML: %v", err))
		return nil, fmt.Errorf("error parsing HTML: %w", err)
	}

	var articles []Article

	// Scraping data
	doc.Find("div.wide-tease-item__wrapper").Each(func(i int, s *goquery.Selection) {
		article := Article{
			Category:    extractText(s, "h2.unibrow"),
			ArticleName: extractText(s, "h2.wide-tease-item__headline"),
			Summary:     extractText(s, "div.wide-tease-item__description"),
			Link:        extractLink(s),
			TimePosted:  extractText(s, "div.wide-tease-item__timestamp"),
		}
		articles = append(articles, article)
	})

	return articles, nil
}

// Extract text from HTML element
// "Selector" is the html element we are extracting text from
func extractText(s *goquery.Selection, selector string) string {
	elem := s.Find(selector).First()
	if elem.Length() == 0 {
		return "N/A"
	}
	return elem.Text()
}

// Extract link from HTML element
func extractLink(s *goquery.Selection) string {
	link := s.Find("a[data-testid='wide-tease-image']")
	href, exists := link.Attr("href")
	if !exists {
		return "N/A"
	}
	return href
}
