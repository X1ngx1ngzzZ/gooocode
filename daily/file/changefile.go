package main

import (
	"io/ioutil"
	"log"
	"strings"
)

/*
func main() {
	input, err := ioutil.ReadFile("./abc.txt")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "one") {
			lines[i] = "LOL"
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("./abc.txt", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
*/
func main() {
	input, err := ioutil.ReadFile("./abc.txt")
	if err != nil {
		log.Fatalln(err)
	}
	output := strings.Replace(string(input), "five", lolbaba, 1)

	err = ioutil.WriteFile("./abc.txt", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

const (
	lolbaba = `sdhsdhjaskdhkasjhdkjashgdkjskjdgksjadkjsgdkjsjg
	`
)
