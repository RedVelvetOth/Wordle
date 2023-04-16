package wordle

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Takes the Data From the dictionnary.txt and turns it into a map

func getdata() map[string]string {

	dict := make(map[string]string)

	content, err := os.Open("Data/dictionnary.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer content.Close()

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		// do something with a line
		line := strings.Split(scanner.Text(), ",")
		dict[line[0]] = strings.ReplaceAll(line[1], "\"", "")
	}

	return dict

}
