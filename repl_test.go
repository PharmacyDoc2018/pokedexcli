package main

import (
	"testing"
	"time"

	"github.com/PharmacyDoc2018/pokedexcli/internal/pokecache"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello  world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hi Hello Hey",
			expected: []string{"hi", "hello", "hey"},
		},
		{
			input:    "explore mt-coronet-2f",
			expected: []string{"explore", "mt-coronet-2f"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Fail. Expected: %s, Actual: %s", expectedWord, word)
			}
		}
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + (5 * time.Millisecond)
	stop := make(chan struct{})
	cache := pokecache.NewCache(baseTime, stop)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
