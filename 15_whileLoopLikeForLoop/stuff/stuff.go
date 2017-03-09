package stuff

import (
	"fmt"
	"math"
)

//Smart is fucked up
const Smart = true

//Hard is fucked up too
const Hard = false

const maxKnowledge = math.MaxUint16

var knowledge uint16

//Learn doesnt make fun
func Learn(way bool) bool {
	var success bool
	if way == Hard {
		fmt.Println("Learn better!")
		success = false
	} else if way == Smart {
		knowledge++
		if knowledge != maxKnowledge {
			fmt.Printf("%s \t (%d/%d) \n", "Learn more!", knowledge, uint64(maxKnowledge))
			success = false
		} else {
			fmt.Printf("%s \t (%d/%d) \n", "You made it! What a sudden you got here.", knowledge, uint64(maxKnowledge))
			success = true
		}
	}
	return success
}
