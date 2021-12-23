package processors

import (
	crypto_rand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// CountLines count number of words in string.
// Example: "line 1\n line 2" = 2
type CountLines struct{}

func (p CountLines) Name() string {
	return "count-lines"
}

func (p CountLines) Alias() []string {
	return nil
}

func (p CountLines) Transform(data []byte, _ ...Flag) (string, error) {
	var lines int
	if len(data) > 0 {
		lines = strings.Count(string(data), "\n") + 1
	}
	return fmt.Sprintf("%d", lines), nil
}

func (p CountLines) Flags() []Flag {
	return nil
}

func (p CountLines) Title() string {
	return "Count Number of Lines"
}

func (p CountLines) Description() string {
	return "Count the number of lines in your text"
}

func (p CountLines) FilterValue() string {
	return p.Title()
}

// SortLines sort given lines, it's not a natural sort.
// Example: 2\n 1\n -> 1\n 2\n
type SortLines struct{}

func (p SortLines) Name() string {
	return "sort-lines"
}

func (p SortLines) Alias() []string {
	return nil
}

func (p SortLines) Transform(data []byte, _ ...Flag) (string, error) {
	sorted := strings.Split(string(data), "\n")
	sort.Strings(sorted)
	return strings.Join(sorted, "\n"), nil
}

func (p SortLines) Flags() []Flag {
	return nil
}

func (p SortLines) Title() string {
	return "Sort Lines"
}

func (p SortLines) Description() string {
	return "Sort lines alphabetically"
}

func (p SortLines) FilterValue() string {
	return p.Title()
}

// ShuffleLines sort given lines, in random order.
type ShuffleLines struct{}

func (p ShuffleLines) Name() string {
	return "shuffle-lines"
}

func (p ShuffleLines) Alias() []string {
	return nil
}

func (p ShuffleLines) Transform(data []byte, _ ...Flag) (string, error) {
	seed, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(int64(time.Now().Nanosecond())))
	if err != nil {
		return "", err
	}
	rand.Seed(seed.Int64())

	shuffle := strings.Split(string(data), "\n")
	rand.Shuffle(len(shuffle), func(i, j int) {
		shuffle[i], shuffle[j] = shuffle[j], shuffle[i]
	})
	return strings.Join(shuffle, "\n"), nil
}

func (p ShuffleLines) Flags() []Flag {
	return nil
}

func (p ShuffleLines) Title() string {
	return "Shuffle Lines"
}

func (p ShuffleLines) Description() string {
	return "Shuffle lines randomly"
}

func (p ShuffleLines) FilterValue() string {
	return p.Title()
}

// UniqueLines sort given lines, in random order.
type UniqueLines struct{}

func (p UniqueLines) Name() string {
	return "unique-lines"
}

func (p UniqueLines) Alias() []string {
	return nil
}

type uniqueLinesModel struct {
	Key   string
	Value int
}
type uniqueLinesList []uniqueLinesModel

func (p uniqueLinesList) Len() int           { return len(p) }
func (p uniqueLinesList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p uniqueLinesList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func (p UniqueLines) Transform(data []byte, _ ...Flag) (string, error) {
	var outputLines []string
	unique := make(map[string]int)
	splitLines := strings.Split(string(data), "\n")

	for k, v := range splitLines {
		unique[v] = k
	}

	linesList := make(uniqueLinesList, len(unique))
	i := 0
	for k, v := range unique {
		linesList[i] = uniqueLinesModel{k, v}
		i++
	}
	sort.Sort(linesList)

	for _, line := range linesList {
		outputLines = append(outputLines, line.Key)
	}

	return strings.Join(outputLines, "\n"), nil
}

func (p UniqueLines) Flags() []Flag {
	return nil
}

func (p UniqueLines) Title() string {
	return "Unique Lines"
}

func (p UniqueLines) Description() string {
	return "Unique Lines"
}

func (p UniqueLines) FilterValue() string {
	return p.Title()
}
