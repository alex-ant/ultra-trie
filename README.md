# ultra-trie
Radix tree with additional interface containers.

### Why ultra-trie?

In addition to classical algorithm of radix tree this package also allows you to store any data type inside each node.

### What is Trie anyway?

A common application of a trie is storing a predictive text or autocomplete dictionary. *From Wikipedia:*

```
Trie is a kind of search tree—an ordered tree data structure that is used to store a dynamic set or
associative array. The position of a node in the tree defines the key with which it is associated.
All the descendants of a node have a common prefix of the string associated with that node, and the
root is associated with the empty string.
```

![diagram](https://github.com/LoudRun/ultra-trie/raw/master/non-code/trie.png "trie")

Trie data structure.

![smaple](https://github.com/LoudRun/ultra-trie/raw/master/non-code/autocomplete.png "autocomplete")

Text autocomplete is one of the common use cases.

### Usage

```bash
$ go get github.com/LoudRun/ultra-trie
```

```Go
package main

import (
	"fmt"
	"log"

	"github.com/LoudRun/ultra-trie"
)

type sampleData struct {
	field1 string
	field2 int
}

func main() {
	// create new tree
	tr := trie.New()

	// add sample data 1
	tr.Add("to", sampleData{
		field1: "a",
		field2: 1,
	})

	// add sample data 2
	tr.Add("tea", sampleData{
		field1: "b",
		field2: 2,
	})

	// add sample data 3
	tr.Add("azm", sampleData{
		field1: "c",
		field2: 3,
	})

	// validate valid prefixes
	if !tr.PrefixExists("t") ||
		!tr.PrefixExists("te") ||
		!tr.PrefixExists("tea") ||
		!tr.PrefixExists("az") ||
		!tr.PrefixExists("a") {
		log.Fatal("valid prefix doesn't exist")
	}

	// get prefix members for "t"
	mmT, mmTErr := tr.GetPrefixMembers("t")
	if mmTErr != nil {
		log.Fatal(mmTErr)
	}

	// get prefix members for "a"
	mmA, mmAErr := tr.GetPrefixMembers("a")
	if mmAErr != nil {
		log.Fatal(mmAErr)
	}

	// print results
	fmt.Println("mmT:", mmT) // mmT: [{to {a 1}} {tea {b 2}}]
	fmt.Println("mmA:", mmA) // mmA: [{azm {c 3}}]
}

```
