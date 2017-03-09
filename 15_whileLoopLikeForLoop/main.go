package main

import (
	"github.com/Goevex/udemyTraining/15_whileLoopLikeForLoop/stuff"
)

func main() {
	learn := func(way bool) bool {
		return stuff.Learn(way)
	}

	const smart = stuff.Smart
	//cons hard = stuff.Hard //pollution

	var knowEverything bool

	for !knowEverything {
		knowEverything = learn(smart)
	}
}
