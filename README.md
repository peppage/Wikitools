# Wikitools [![Go Report Card](https://goreportcard.com/badge/github.com/peppage/Wikitools)](https://goreportcard.com/report/github.com/peppage/Wikitools)

A collection of command line tools that help me keep track of categories.

## Wiki Category Comparison

wikicc is a command line application to check if pages in a category are listed
on a page.

For example:

```
./wikicc -c "Spoken articles" -p "Wikipedia:Spoken_articles" -w
```

The spoken articles are listed on the page and if someone forgets to add them
it would be a lot of work go through the category and check that they are
listed.

The app outputs either to stdout or an html file that links to the page.
