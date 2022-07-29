package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	frequency := FreqMap{}

	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, letter := range l {
		wg.Add(1)

		go func(s string) {
			mu.Lock()

			for _, r := range s {
				frequency[r]++
			}

			mu.Unlock()
			wg.Done()
		}(letter)
	}

	wg.Wait()

	return frequency
}
