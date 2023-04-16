package wordle

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/enescakir/emoji" // Used To Make The Colored Sqaures
	"golang.org/x/exp/maps"
)

// Gets The Maps That Holds Our Data

var dict = getdata()

// The Struct For The Logic Of The game

type game struct {
	Tries      int
	Guesses    string
	Word       string
	Definition string
	win        int
}

// Randomly Generates a Word From Our map (dict)

func getword() string {
	keys := maps.Keys(dict)
	word := rand.Intn(len(keys))
	return keys[word]
}

// Checks if the word entered by the user is valid or not

func word_in_dict(text string) bool {
	for k := range dict {
		if text == k {
			return true
		}
	}
	return false
}

// Prompts the user to input his guess

func nextturn(g game) game {
	fmt.Print("Enter A 5 Letter Word: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = text[:5]
	for len(text) != 5 && word_in_dict(text) {
		fmt.Print("Please Enter A 5 Letter, Your Word Was Incorrect: ")
		text, _ = reader.ReadString('\n')
		text = text[:5]
	}
	g.Guesses = text
	g.Tries++
	return g
}

// Initialises the game variable for the start of the game

func newgame() game {
	word := getword()
	g := game{0, "", word, dict[word], 0}
	return g
}

// The logic function for wordle

func check(g game) []int {
	wr := make([]int, 5, 5)

	for i, v := range g.Word {
		if string(v) == string(g.Guesses[i]) {
			wr[i] = 1
		} else if strings.Contains(g.Word, string(g.Guesses[i])) {
			wr[i] = -1
		} else {
			wr[i] = 0
		}
	}
	return wr
}

// Prints the result of the users guess

func result(g game) {
	ch := check(g)
	count := 0
	for _, i := range ch {
		if i == 1 {
			fmt.Print(emoji.GreenSquare)
			count += 1
		} else if i == -1 {
			fmt.Print(emoji.YellowSquare)
		} else {
			fmt.Print(emoji.BlackLargeSquare)
		}
	}

	if count == 5 {
		g.win = 1
	}

	fmt.Println("	(Try ", g.Tries, "/6): ", g.Guesses)

}

// The main function for the game (ran in main)

func Game() {
	fmt.Println("!!!Welcome To Wordle!!!")
	fmt.Println("You Have 6 Tries To Guess The Word")
	fmt.Println(emoji.GreenSquare, ": Means Correct Letter In The Correct Spot")
	fmt.Println(emoji.YellowSquare, ": Means Correct Letter In The Wrong Spot")
	fmt.Println(emoji.BlackLargeSquare, ": Means Letter Not Found In The Word")

	g := newgame()

	for g.Tries < 6 {
		g = nextturn(g)
		result(g)
		if g.win == 1 {
			fmt.Println("!!! Congratulation You Guessed The Word!!!")
			break
		}
	}

	fmt.Println("Your Word Was ", g.Word, ": ", g.Definition)
}
