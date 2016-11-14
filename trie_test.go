package trie

import "testing"

type sampleData struct {
	name   string
	field1 string
	field2 int
}

func TestTrie(t *testing.T) {
	// create new tree
	tr := New()

	// initialize sample data variable
	var sd sampleData

	// add data 1
	sd = sampleData{
		name:   "to",
		field1: "a",
		field2: 1,
	}
	tr.Add(sd.name, sd)

	// add data 2
	sd = sampleData{
		name:   "tea",
		field1: "b",
		field2: 2,
	}
	tr.Add(sd.name, sd)

	// add data 3
	sd = sampleData{
		name:   "azm",
		field1: "c",
		field2: 3,
	}
	tr.Add(sd.name, sd)

	// validate valid prefixes
	if !tr.PrefixExists("t") ||
		!tr.PrefixExists("te") ||
		!tr.PrefixExists("tea") ||
		!tr.PrefixExists("az") ||
		!tr.PrefixExists("a") {
		t.Error("valid prefix doesn't exist")
		return
	}

	// validate invalid prefixes
	if tr.PrefixExists("q") ||
		tr.PrefixExists("qwerty") ||
		tr.PrefixExists("f") ||
		tr.PrefixExists("gf") {
		t.Error("invalid prefix exists")
		return
	}
}
