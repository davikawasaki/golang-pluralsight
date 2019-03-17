package main

import (
	"fmt"
)

func main() {
	// key (string) and value (int)
	leagueTitles := make(map[string]int)
	leagueTitles["Sunderland"] = 6
	leagueTitles["Newcastle"] = 4

	recentHead2Head := map[string]int{
		"Sunderland": 5,
		"Newcastle":  0,
	}

	fmt.Printf("\nLeague titles: %v\nRecent head to heads: %v\n",
		leagueTitles, recentHead2Head)
}
