package da

import (
	"github.com/liuzl/cedar-go"
)

type Dict struct {
	Trie   *cedar.Cedar
	Values [][]string
}
