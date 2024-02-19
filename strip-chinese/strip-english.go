// Strips English Unicode ranges from input text
// Example usages include stripping English characters from a transcription file that contains English and Chinese
// References: https://stackoverflow.com/questions/47068770/how-do-i-remove-all-the-chinese-characters-from-a-string
// e.g. `cat input.txt | ./strip-english > output.txt`

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func removeEnglish(input string) string {
	// Define the regular expression pattern to match English letters and common punctuation marks
    pattern := "[a-zA-Z.,!?;:()\\[\\]\\{\\}\"'`\\s-]+"
	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Replace matched characters with an empty string
	output := re.ReplaceAllString(input, "")

	return output
}

func main() {
	// Create a scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Process each line from stdin
	for scanner.Scan() {
		input := scanner.Text()

		// Remove English characters and common punctuation marks
		result := removeEnglish(input)

		// Print the result
		fmt.Println(result)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading standard input: %v\n", err)
		os.Exit(1)
	}
}

