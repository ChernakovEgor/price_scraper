package scraper

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/ChernakovEgor/price_scraper/internal/database"
	"golang.org/x/net/html"
)

const agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

type Scraper struct {
	db *database.Queries
}

func NewScraper(db *database.Queries) Scraper {
	return Scraper{db}
}

func (s *Scraper) ProcessURLs() error {
	urls, err := s.db.GetURLs(context.Background())
	if err != nil {
		return fmt.Errorf("getting urls from db: %v", err)
	}

	log.Print("Starting to process URLs...")
	for _, url := range urls {
		_, err := s.FetchURL(int(url.ID), url.Url)
		if err != nil {
			log.Printf("error fetching '%s...': %v", url.Url[:10], err)
		}
	}

	return nil
}

func (s *Scraper) FetchURL(urlID int, url string) (int, error) {
	log.Printf("Begin fetching %s", url)
	client := http.DefaultClient
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, fmt.Errorf("creating request: %v", err)
	}

	req.Header.Set("User-Agent", agent)

	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("fetching url '%s': %w", url, err)
	}
	defer resp.Body.Close()

	rawHTML, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("reading response body of GET '%s': %w", url, err)
	}

	params := database.AddParsingResultParams{
		UrlID:      int64(urlID),
		StatusCode: int64(resp.StatusCode),
		RawBody:    sql.NullString{String: string(rawHTML), Valid: true},
	}
	result, err := s.db.AddParsingResult(context.Background(), params)

	return int(result.ID), nil
}

func findPricePath(parent *html.Node, price string) {
	for elem := range parent.Descendants() {
		// fmt.Printf("inspecting %v\n", elem.DataAtom)
		if elem.Data == price {
			fmt.Println("Match!")
			extractPath(elem)
		}
	}
}

func extractPath(node *html.Node) {
	var path []string
	for {
		if node != nil {
			if node.Type == html.ElementNode {
				attrString := attributesToString(node.Attr)
				elem := fmt.Sprintf("<%v %s>", node.DataAtom, attrString)
				path = append(path, elem)
			}
		} else {
			break
		}
		node = node.Parent
	}
	slices.Reverse(path)
	fmt.Println(path)
}

func attributesToString(attrs []html.Attribute) string {
	var extracted []string
	for _, attr := range attrs {
		extracted = append(extracted, fmt.Sprintf(`%s="%s"`, attr.Key, attr.Val))
	}

	return strings.Join(extracted, " ")
}
