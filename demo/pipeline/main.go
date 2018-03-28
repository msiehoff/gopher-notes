package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {
	pageNumberCh := make(chan int)
	playerURLsCh := make(chan []string)
	playerURLCh := make(chan string)
	playerCh := make(chan Player)
	playersSavedCh := make(chan bool)

	go p_getPlayerURLs(pageNumberCh, playerURLsCh)
	go p_getPlayerFromURL(playerURLsCh, playerURLCh)
	go p_getPlayerURL(playerURLCh, playerCh)
	go p_savePlayer(playerCh, playersSavedCh)

	for i := 0; i < 5; i++ {
		pageNumberCh <- i
	}
	close(pageNumberCh)

	for _ = range playersSavedCh {
		fmt.Printf("\n ******* players saved received *******")
	}

	fmt.Printf("\n\n\nDONE")
}

func p_getPlayerURLs(pageNumberCh chan int, playerURLCh chan []string) {
	for pageNumber := range pageNumberCh {
		timer := time.NewTimer(2 *time.Second)
		<- timer.C
		playerURLCh <- playerURLs

		fmt.Printf("\nPlayer URLS Page: %d", pageNumber)
	}

	close(playerURLCh)
}

func p_getPlayerFromURL(playerURLsCh chan []string, playerURLCh chan string) {
	for urls := range playerURLsCh {
		for _, url := range urls {
			playerURLCh <- url
		}
	}

	close(playerURLsCh)
}

func p_getPlayerURL(playerURLCh chan string, playerCh chan Player) {
	for _ = range playerURLCh {
		simulateWait()
		player := chooseRandomPlayer()
		playerCh <- player
		fmt.Printf("\nPlayer Gotten: %v", player)
	}

	close(playerCh)
}

func p_savePlayer(playerCh chan Player, playerSavedCh chan bool) {
	var playersToSave []Player
	minToSave := 5
	var queuedForSaving int

	for player := range playerCh {
		queuedForSaving++
		playersToSave = append(playersToSave, player)
		fmt.Printf("\nPlayer queued for Saving: %v", player.String())

		if queuedForSaving > minToSave {
			fmt.Printf("\nSaving Players")
			for _, p := range playersToSave {
				fmt.Printf("\n --- Player Saved ---: %v", p.String())
			}

			playersToSave = []Player{}
			queuedForSaving = 5
		}
	}

	defer close(playerSavedCh)
}


var playerURLs = []string{
	"https://www.mlssoccer.com/players/josef-martinez",
	"https://www.mlssoccer.com/players/nicolas-lodeiro",
	"https://www.mlssoccer.com/players/jack-jones",
	"https://www.mlssoccer.com/players/mike-smith",
	"https://www.mlssoccer.com/players/thomas-wayne",
	"https://www.mlssoccer.com/players/jacob-ellerby",
	"https://www.mlssoccer.com/players/blah-foo",
	"https://www.mlssoccer.com/players/kyle-guilstre",
	"https://www.mlssoccer.com/players/rico-rodriguez",
	"https://www.mlssoccer.com/players/jefferson-savarino",
}

var Players = []string {
	"Zlatan Ibrahimovic",
	"josef martinez",
	"nicolas lodeiro",
	"bastian schweinsteiger",
	"johan kappelhoff",
	"miguel almiron",
}

var Teams = []string {
	"LA Galaxy",
	"Chicago Fire",
	"Seattle Sounders",
	"Portland Timbers",
	"LAFC",
	"NYRB",
}

type Player struct {
	Name string
	Team string
}

func (p *Player) String() string {
	return fmt.Sprintf("%s (%s)", p.Name, p.Team)
}

func simulateWait() {
	timer := time.NewTimer(1 * time.Second)

	// Wait until timer is done
	// to simulate saving to DB
	<- timer.C
}

func chooseRandomPlayer() Player {
	return Player {
		Name: Players[rand.Intn(5)],
		Team: Teams[rand.Intn(5)],
	}
}