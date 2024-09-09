package main

import (
	"fmt"
)

func numTilePossibilities(tiles string) int {
	used := make([]bool, len(tiles))
	result := map[string]bool{}
	dfs(tiles, "", &used, &result)
	return len(result)
}

func dfs(tiles, current string, used *[]bool, result *map[string]bool) {
	if current != "" {
		(*result)[current] = true
	}

	for i := 0; i < len(tiles); i++ {
		if (*used)[i] {
			continue
		}

		(*used)[i] = true
		dfs(tiles, current+string(tiles[i]), used, result)
		(*used)[i] = false
	}
}

func main() {
	tiles := "AAB"
	result := numTilePossibilities(tiles)
	fmt.Println(result)
}
