/*
 * wikicc is a utility to help list pages in a category. For example
 * the Wikipedia:Spoken_articles page holds a list of all spoken wikipedia articles
 * but editors forget to add new files.
 *
 * wikicc will get all pages in a category and make sure the title is listed on
 * the page.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	flag "github.com/ogier/pflag"
	"github.com/peppage/wikitools/pkg/formatters"
	"github.com/peppage/wikitools/pkg/wiki"
)

func main() {
	cat := flag.StringP("category", "c", "", "Category name without prefix")
	page := flag.StringP("page", "p", "", "Page that lists page titles in category")
	web := flag.BoolP("web", "w" , false, "Output to an html file")
	file := flag.StringP("file", "f", "output.html", "Filename for html output")
	flag.Usage = usage
	flag.Parse()

	

	client := wiki.NewClient()
	titles := client.GetPagesInCategory(*cat)
	pageContents := strings.ToLower(client.GetPage(*page))

	notFoundTitles := []string{}
	for _, t := range titles {
		if !strings.Contains(pageContents, strings.ToLower(t)) {
			notFoundTitles = append(notFoundTitles, t)
		}
	}

	if *web {
		htmlOut(notFoundTitles, *file)
	} else {
		stOut(notFoundTitles)
	}
}

func stOut(titles []string) {
	for _, t := range titles {
		fmt.Println(t)
	}
}

func htmlOut(titles []string, filename string) {
	format := formatters.NewHtml()

	fi, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer fi.Close()

	w := bufio.NewWriter(fi)

	format.ListPages(titles, w)

	if err = w.Flush(); err != nil {
		panic(err)
	}
}

func usage() {
	fmt.Println(`wikicc is a tool for comparing pages in a category to a list.`)
	flag.PrintDefaults()
}
