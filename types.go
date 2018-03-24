package da

import (
	"github.com/liuzl/cedar-go"
)

// Dict contains the Trie and dict values
type Dict struct {
	Trie   *cedar.Cedar
	Values [][]string
}
