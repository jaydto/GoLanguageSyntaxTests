package interfacesgo

import (
	"fmt"
	"math/rand"
)

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "woof"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "meaow"
}

func Interfaces1() {
	var animal Animal
	animal = Dog{}
	fmt.Println(animal.Speak())

	animal = Cat{}
	fmt.Println(animal.Speak())
}

// writing interface for a team of players

type Teams interface {
	TeamInfo()
}
type Team struct {
	teamName   string
	teamScores uint8
	teamGames  uint8
	players    []TeamPlayer
}

type TeamPlayer struct {
	playerName     string
	playerPosition string
	team           *Team
}

// Implement the TeamInfo method for the Team struct
func (t *Team) TeamInfo() {
	fmt.Printf("Team Name: %s\n", t.teamName)
	fmt.Printf("Team Scores: %d\n", t.teamScores)
	fmt.Printf("Team Games: %d\n", t.teamGames)
	fmt.Println("Players:")
	for _, player := range t.players {
		fmt.Printf("  - Name: %s, Position: %v\n", player.playerName, player.playerPosition)
	}
}

// Function to add a player to a team
func (t *Team) AddPlayer(player TeamPlayer) {
	player.team = t // Set the team for the player
	t.players = append(t.players, player)
}

func Players() {

	// Create a team
	team := Team{
		teamName:   "Warriors",
		teamScores: 35,
		teamGames:  5,
	}

	team2 := Team{
		teamName:   "Arsenal",
		teamScores: 95,
		teamGames:  10,
	}

	// Create players
	player1 := TeamPlayer{
		playerName:     "John Doe",
		playerPosition: "Forward",
	}
	player2 := TeamPlayer{
		playerName:     "Jane Smith",
		playerPosition: "Guard",
	}
	player3 := TeamPlayer{
		playerName:     "Mik Smith",
		playerPosition: "MidFilder",
	}

	// Add players to the team
	team.AddPlayer(player1)
	team.AddPlayer(player2)
	team.AddPlayer(player3)

	team2.TeamInfo()

	// Display team information
	// team.TeamInfo()
}

type Overview interface {
	kickBall()
}

type FootbalPlayer struct {
	stamina uint8
	power   uint8
}

type CR7 struct {
	stamina uint8
	power   uint8
	sui     uint8
}

func (f FootbalPlayer) kickBall() {
	shot := f.stamina + f.power
	fmt.Println(" Shot power", shot)

}

func (f CR7) kickBall() {
	shot := f.stamina + f.power + f.sui
	fmt.Println(" Cr7 Shot power", shot)

}
func PlayerEvaluation() {

	team := make([]Overview, 11)

	for i := 0; i < len(team)-1; i++ {
		team[i] = FootbalPlayer{
			stamina: uint8(rand.Intn(10)),
			power:   uint8(rand.Intn(10)),
		}
	}
	team[len(team)-1] = CR7{
		power:   20,
		stamina: 30,
		sui:     10,
	}

	for i := range team {
		team[i].kickBall()
	}

	player := FootbalPlayer{
		stamina: 30,
		power:   50,
	}

	player2 := FootbalPlayer{
		stamina: 50,
		power:   20,
	}

	player.kickBall()
	player2.kickBall()
}
