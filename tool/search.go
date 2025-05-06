package tool

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type Search struct{}

func (s *Search) Name() string {
	return "Search"
}

func (s *Search) Description() string {
	return "Performs a web search using DuckDuckGo (scraped)."
}

func (s *Search) Run(input string) (string, error) {
	query := url.QueryEscape(strings.TrimSpace(input))
	searchURL := fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s", query)

	resp, err := http.Get(searchURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch search results, status: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Extract first result link using regex (simple but fragile)
	re := regexp.MustCompile(`<a rel="nofollow" class="result__a" href="(.*?)">`)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		return "No results found.", nil
	}

	topResult := matches[1]
	return fmt.Sprintf("Top result: %s", topResult), nil
}
