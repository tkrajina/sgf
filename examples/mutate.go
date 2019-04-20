package main

// Example of using MutateTree() with a function argument. Rotates the whole tree.

import (
	"fmt"
	"os"

	sgf ".."
)

func main() {
	original := sgf.LoadArgOrQuit(1)					// Equivalent to sgf.Load(os.Args[1])
	mutated := original.MutateTree(rotate_clockwise)
	original.GetEnd().Board().Dump()					// Unharmed
	fmt.Printf("\n")
	mutated.GetEnd().Board().Dump()
	mutated.Save(os.Args[1] + ".rotated.sgf")
	fmt.Printf("\n")
	fmt.Printf("%d nodes in original, %d nodes in mutated.\n", original.NodesInTree(), mutated.NodesInTree())
}

// The mutator function is shown the original node and must return the properties
// that it wants the mutated node to have...

func rotate_clockwise(original *sgf.Node) map[string][]string {

	props := original.AllProperties()		// Fetches a copy, which is safe to edit.
	board := original.Board()

	for _, key := range []string{"AB", "AW", "AE", "B", "CR", "MA", "SL", "SQ", "TR", "W"} {
		for i, s := range props[key] {
			if len(s) == 2 {
				x, y, onboard := sgf.ParsePoint(s, board.Size)
				if onboard {
					new_x := board.Size - 1 - y
					new_y := x
					props[key][i] = sgf.Point(new_x, new_y)
				}
			}
		}
	}

	return props
}
