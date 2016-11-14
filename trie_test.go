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

	// validate prefix data
	var mm []Member
	var mmErr error
	var mmData sampleData
	var mmDataOK bool

	// prefix "a"
	mm, mmErr = tr.GetPrefixMembers("a")
	if mmErr != nil {
		t.Errorf("error received while retrieving prefix members: %s", mmErr.Error())
		return
	}

	if len(mm) != 1 {
		t.Errorf("invalid number of prefix members received: %v", mm)
		return
	}

	if mm[0].Key != "azm" {
		t.Errorf("invalid number of prefix member key received: %s", mm[0].Key)
		return
	}

	mmData, mmDataOK = mm[0].Data.(sampleData)
	if !mmDataOK {
		t.Error("failed to convert member's data into valid structure")
		return
	}

	if mmData.field1 != "c" ||
		mmData.field2 != 3 ||
		mmData.name != "azm" {
		t.Error("invalid member's data received")
		return
	}

	// prefix "t"
	mm, mmErr = tr.GetPrefixMembers("t")
	if mmErr != nil {
		t.Errorf("error received while retrieving prefix members: %s", mmErr.Error())
		return
	}

	if len(mm) != 2 {
		t.Errorf("invalid number of prefix members received: %v", mm)
		return
	}

	// member 1
	if mm[0].Key != "to" {
		t.Errorf("invalid number of prefix member key received: %s", mm[0].Key)
		return
	}

	mmData, mmDataOK = mm[0].Data.(sampleData)
	if !mmDataOK {
		t.Error("failed to convert member's data into valid structure")
		return
	}

	if mmData.field1 != "a" ||
		mmData.field2 != 1 ||
		mmData.name != "to" {
		t.Error("invalid member's data received")
		return
	}

	// member 2
	if mm[1].Key != "tea" {
		t.Errorf("invalid number of prefix member key received: %s", mm[1].Key)
		return
	}

	mmData, mmDataOK = mm[1].Data.(sampleData)
	if !mmDataOK {
		t.Error("failed to convert member's data into valid structure")
		return
	}

	if mmData.field1 != "b" ||
		mmData.field2 != 2 ||
		mmData.name != "tea" {
		t.Error("invalid member's data received")
		return
	}

}
