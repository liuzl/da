package da

import (
	"github.com/liuzl/cedar-go"
	"sync"
)

type Dict struct {
	sync.RWMutex
	Trie   *cedar.Cedar
	Values map[int]string
}
