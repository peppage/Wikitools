/*
 * wikicsort is a utility to help sort pages in a category.
 * Can sorty by size, quality, and views
 */

package main

import (
	"fmt"
	"strings"
	"strconv"
	"sort"

	flag "github.com/ogier/pflag"
	"github.com/peppage/wikitools/pkg/wiki"
)

func main() {
	cat := flag.StringP("category", "c", "", "Category name without prefix")
	//quality := flag.BoolP("importance", "i", true, "Sort category by quality")
	size := flag.BoolP("size", "s", false, "Sort category by page size")
	web := flag.BoolP("web", "w" , false, "Output to an html file")
	//file := flag.StringP("file", "f", "output.html", "Filename for html output")
	// Add decending or ascending
	flag.Usage = usage
	flag.Parse()

	client := wiki.NewClient()
	titles := client.GetPagesInCategory(*cat)

	pages := []page{}
	for _, t := range titles {
		pageContents := strings.ToLower(client.GetPage(t))
		bs := []byte(pageContents)
		pages = append(pages, page{
			title: t,
			size: len(bs),
		})
	}

	if(*size) {
		sort.Sort(bySizeAsc(pages))
	}

	if *web {
		//htmlOut(notFoundTitles, *file)
	} else {
		stOut(pages)
	}
}

func stOut(pages []page) {
	for _, p := range pages {
		i := strconv.Itoa(p.size)
		fmt.Println(p.title + " - (" + i + ")")
	}
}

func usage() {
	fmt.Println(`wikicsort is a tool for sorting pages in a category.`)
	flag.PrintDefaults()
}

type page struct{
	title string
	size int
}

type bySizeAsc []page

func (s bySizeAsc) Len() int {
	return len(s)
}

func (s bySizeAsc) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s bySizeAsc) Less(i, j int) bool {
	return s[i].size < s[j].size
}
