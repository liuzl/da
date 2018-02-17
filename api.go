package da

func (self *Dict) Add(k, v string) error {
	self.Lock()
	defer self.Unlock()
	v, err := self.Trie.Get([]byte(k))
	if err != nil {
		return err
	}

	return nil
}
