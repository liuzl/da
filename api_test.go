package da

import (
	"testing"
)

func TestBuildAndSave(t *testing.T) {
	d, err := BuildFromFile("test_data/STPhrases.txt")
	if err != nil {
		t.Error(err)
	}
	err = d.Save("cedar")
	if err != nil {
		t.Error(err)
	}
}

func TestLoad(t *testing.T) {
	d, err := Load("cedar")
	if err != nil {
		t.Error(err)
	}
	s := `一丝不挂的一分钟，是一前一后，有点意思吧`
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		ret, err := d.PrefixMatch(string(r[i:]))
		if err != nil {
			t.Error(err)
		}
		t.Log(ret)
	}
}
