package main

import (
	"fmt"
	"strings"
)

func getInput() string{

	return`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
	
}


func main(){
	lines := strings.Split(getInput(), "\n")
	colLen := len(lines[0])-1

	treeCount := 0
	currPos := 0
	for _, line := range lines{
		currPos = (currPos+3)%colLen
		if line[currPos] == '#'{
			treeCount += 1
		}
	}
	fmt.Println("tree count",treeCount)
}
