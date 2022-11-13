package main

import (
	"container/ring"
	"fmt"
)

func main() {
	players := 473

	fmt.Println(game(players, 70904))
	fmt.Println(game(players, 7090400))
}

func game(players, lastMarble int) int {
	score := make([]int, players)
	curPlayer := 0

	circle := ring.New(1)
	circle.Value = 0

	for marble := 0; marble <= lastMarble; marble++ {
		if marble%23 == 0 {
			// move back 8
			circle = circle.Move(-8)
			// remove a marble
			del := circle.Unlink(1)
			// set the position to after the deleted one
			circle = circle.Next()
			// add the score
			score[curPlayer] += marble + del.Value.(int)

		} else {
			// add new marble in between the next two marbles and set the position to the new marble
			newItem := ring.New(1)
			newItem.Value = marble
			circle = circle.Move(1).Link(newItem).Move(-1)
		}

		// advance player
		curPlayer = (curPlayer + 1) % players
	}

	// find winning score
	maxScore := 0
	for _, s := range score {
		if s > maxScore {
			maxScore = s
		}
	}
	return maxScore
}
