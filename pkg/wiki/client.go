package wiki

import (
	"cgt.name/pkg/go-mwclient"
)

// Client is way to interact with wikipedia
type Client interface {
	Login(username, password string)
	GetPagesInCategory(categoryName string) []string
	GetPage(title string) string
}

// NewClient gets a wiki client
func NewClient() Client {
	w, err := mwclient.New("https://en.wikipedia.org/w/api.php", "github.com/peppage/wikitools")
	handleErr(err)

	return &client{
		mwClient: w,
	}
}

type client struct {
	mwClient *mwclient.Client
}

// Login a user to wikipedia
func (c *client) Login(username, password string) {
	// This throws warnings which are put into the returning error
	c.mwClient.Login(username, password)
	//handleErr(err)
}

// Get all pages in a category
func (c *client) GetPagesInCategory(categoryName string) []string {
	params := map[string]string{
		"action":  "query",
		"list":    "categorymembers",
		"cmtitle": "Category:" + categoryName,
		"cmlimit": "500",
	}

	hasMore := true
	titles := []string{}
	for ok := true; ok; ok = hasMore {
		resp, err := c.mwClient.Get(params)
		handleErr(err)

		cmContinue, err := resp.GetString("continue", "cmcontinue")
		if err != nil {
			hasMore = false
		} else {
			params["cmcontinue"] = cmContinue
		}

		pages, err := resp.GetObjectArray("query", "categorymembers")
		handleErr(err)
		for _, p := range pages {
			m := p.Map()
			v, err := m["title"].String()
			handleErr(err)
			titles = append(titles, v)
		}
	}

	return titles
}

// Get the parsed contents of a page (html)
func (c *client) GetPage(title string) string {
	params := map[string]string{
		"action": "parse",
		"page":   title,
	}

	resp, err := c.mwClient.Get(params)
	handleErr(err)

	value, err := resp.GetString("parse", "text")
	handleErr(err)
	return value
}
