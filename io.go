package da

import (
	"bufio"
	"encoding/gob"
	"github.com/liuzl/cedar-go"
	"os"
	"path/filepath"
)

func Load(dir string) (*Dict, error) {
	trieFile := filepath.Join(dir, "trie")
	valueFile := filepath.Join(dir, "values")
	trie := cedar.New()
	err := trie.LoadFromFile(trieFile, "gob")
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(valueFile, os.O_RDONLY, 0600)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	in := bufio.NewReader(file)
	dataDecoder := gob.NewDecoder(in)
	var values []string
	err = dataDecoder.Decode(&values)
	if err != nil {
		return nil, err
	}
	return &Dict{Trie: trie, Values: values}, nil
}

func (self *Dict) Save(dir string) error {
	self.Lock()
	defer self.Unlock()
	trieFile := filepath.Join(dir, "trie")
	valueFile := filepath.Join(dir, "values")
	err := self.Trie.SaveToFile(trieFile, "gob")
	if err != nil {
		return err
	}
	file, err := os.OpenFile(valueFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	out := bufio.NewWriter(file)
	defer out.Flush()
	dataEncoder := gob.NewEncoder(out)
	return dataEncoder.Encode(&self.Values)
}
