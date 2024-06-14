package main

import "fmt"

// test
func main() {
	text := []byte("Пример текста с http://example.com и другими ссылками http://example.org и http://example.net")
	pattern := []byte("http://")

	for i := 0; i < len(text)-len(pattern); i++ {
		if string(text[i:i+len(pattern)]) == "http://" {
			j := i + len(pattern)
			for j < len(text) && text[j] != ' ' {
				text[j] = '*'
				j++
			}
			i = j
		}
	}

	fmt.Println(string(text))
}
