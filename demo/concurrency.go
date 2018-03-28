package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	savePlayerInfo()
}

func savePlayerInfo() {
	playerURLs := getPlayerURLs()

	doneCh := make(chan bool)
	for _, url := range playerURLs {
		go getAndSavePlayer(url, doneCh)
	}

	for i := 0; i < len(playerURLs); i++ {
		<- doneCh
		fmt.Printf("\n-------- player completed")
	}
}

func getAndSavePlayer(url string, doneCh chan bool) {
	player := getPlayer(url)
	savePlayer(player)
	doneCh <- true
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

func getPlayerURLs() []string {
	return playerURLs
}

func getPlayer(url string) Player {
	p := Player {
		Name: Players[rand.Intn(5)],
		Team: Teams[rand.Intn(5)],
	}

	fmt.Printf("\nPlayer Gotten: %v", p.String())

	return p
}

func savePlayer(player Player) {
	timer := time.NewTimer(2 * time.Second)

	// Wait until timer is done
	// to simulate saving to DB
	<- timer.C

	fmt.Printf("\nPlayer Saved: %v", player.String())
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
