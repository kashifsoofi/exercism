package letter

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

// ConcurrentFrequency counts the frequency of each rune in a given list of texts and returns this
// data as a FreqMap.
func ConcurrentFrequency(textList []string) FreqMap {
	// make channel to keep hold of frequence in each string
	channel := make(chan FreqMap, len(textList))
	// loop over passed string
	for _, text := range textList {
		// foreach string call a gorouting to execute Frequence and put result in channel
		go func(s string) {
			channel <- Frequency(s)
		}(text)
	}

	m := FreqMap{}
	// loop to iterate over each item in channel (can use i as well)
	for range textList {
		// loop over each result of call to Frequency and update new FreqMap
		for key, value := range <-channel {
			m[key] += value
		}
	}
	return m
}
