package scraper

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"

	"github.com/ChernakovEgor/price_scraper/internal/database"
	"golang.org/x/net/html"
)

func FetchURL(db *database.Queries, urlID int, url string) (int, error) {
	resp, err := http.Get(url)
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
	result, err := db.AddParsingResult(context.Background(), params)

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
