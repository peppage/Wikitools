package formatters

import "io"

type Formatter interface {
	ListPages(pageNames []string, writer io.Writer)
}
