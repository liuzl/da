package da

import (
	"bufio"
	"fmt"
	"github.com/liuzl/cedar-go"
	"io"
	"os"
	"strings"
)

func Build(fileName string) (*Dict, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	trie := cedar.New()
	values := [][]string{}
	br := bufio.NewReader(file)
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		items := strings.Fields(line)
		if len(items) < 2 {
			continue
		}
		err = trie.Insert([]byte(items[0]), len(values))
		if err != nil {
			return nil, err
		}
		values = append(values, items[1:])
	}
	return &Dict{Trie: trie, Values: values}, nil
}

func (self *Dict) PrefixMatch(str string) (map[string][]string, error) {
	if self.Trie == nil {
		return nil, fmt.Errorf("Trie is nil")
	}
	ret := make(map[string][]string)
	for _, id := range self.Trie.PrefixMatch([]byte(str), 0) {
		key, err := self.Trie.Key(id)
		if err != nil {
			return nil, err
		}
		value, err := self.Trie.Value(id)
		if err != nil {
			return nil, err
		}
		ret[string(key)] = self.Values[value]
	}
	return ret, nil
}

func (self *Dict) Get(str string) ([]string, error) {
	if self.Trie == nil {
		return nil, fmt.Errorf("trie is nil")
	}
	id, err := self.Trie.Get([]byte(str))
	if err != nil {
		return nil, err
	}
	value, err := self.Trie.Value(id)
	if err != nil {
		return nil, err
	}
	return self.Values[value], nil
}
