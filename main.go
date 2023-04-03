package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Gamemaster struct {
	turn    int
	players []Player
	board   Board
}

func (gm *Gamemaster) welcome() {
	fmt.Println("Welcome to this tiny command line tictacgo. I'm you're gamemaster!")
}
func (gm *Gamemaster) readInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSuffix(text, "\n")

}
func (gm *Gamemaster) setupUsers() {
	fmt.Println("Player 1, what's your username?")
	gm.players = append(gm.players, Player{name: gm.readInput(), symbol: "x", stonesLeft: 4})
	fmt.Println("Ok, your symbol for the game is ===> 'x'")
	fmt.Println("Whats the 2nd player's name?")
	gm.players = append(gm.players, Player{name: gm.readInput(), symbol: "z", stonesLeft: 4})
	fmt.Println("Ok, your symbol for the game is ===> 'z'")

}
func (gm *Gamemaster) startGame() {
	gm.board = Board{height: 3, width: 3}
	fmt.Println("Ok, let's start the game!")
	fmt.Println("-------------------------")
	gm.board.printBoard(&gm.players[0], &gm.players[1], gm.board.height)
	gm.turn = 1
}
func (gm *Gamemaster) handleTurn() {
	fmt.Println()
	fmt.Println("------ TURN " + fmt.Sprint(gm.turn) + " --------")
	gm.playerTurn(&gm.players[0])
	gm.board.printBoard(&gm.players[0], &gm.players[1], gm.board.height)
	gm.playerTurn(&gm.players[1])
	gm.board.printBoard(&gm.players[0], &gm.players[1], gm.board.height)
	if gm.players[0].stonesLeft > 0 {
		gm.handleTurn()
	}
}

func (gm *Gamemaster) playerTurn(player *Player) {
	fmt.Println(player.name + ", it's your turn. Enter the field to place a stone like 'x-coordinate y-coordinate'")
	coordinate := strings.Split(gm.readInput(), " ")
	xcoordinate, xerr := strconv.Atoi(coordinate[0])
	ycoordinate, yerr := strconv.Atoi(coordinate[1])
	if xerr == nil && yerr == nil {
		player.setStone(xcoordinate, ycoordinate)
	} else {
		fmt.Println("There was an error!")
	}

}

type Board struct {
	height int
	width  int
}

func (board *Board) printBoard(player1 *Player, player2 *Player, turns int) {
	line := "0"
	for j := 1; j <= turns; j++ {
		line = line + " " + fmt.Sprint(j)
	}
	fmt.Println(line)
	for i := int(1); i <= board.height; i++ {
		line = fmt.Sprint(i) + "|"
		for e := int(1); e <= board.width; e++ {

			if player1.getFieldValue(e, i) {
				line = line + player1.symbol + "|"
			} else if player2.getFieldValue(e, i) {
				line = line + player2.symbol + "|"
			} else {
				line = line + "-|"
			}
		}
		fmt.Println(line)
	}
}

type Stones [][]int // [x-axis,y-axis]
type Player struct {
	name       string
	stones     Stones
	stonesLeft int
	symbol     string
}

func (player *Player) getFieldValue(xAxis int, yAxis int) bool {
	for i := range player.stones {
		if reflect.DeepEqual(player.stones[i], []int{xAxis, yAxis}) {
			return true
		}
	}
	return false
}

func (player *Player) setStone(xAxis int, yAxis int) {
	//coordinates := []int{xAxis, yAxis}
	player.stones = append(player.stones, []int{xAxis, yAxis})
	player.stonesLeft = player.stonesLeft - 1
}

func main() {

	/**app := &cli.App{
		Name:  "TicTacGo",
		Usage: "Play a little",
		Action: func(*cli.Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}**/
	gm := Gamemaster{}
	gm.welcome()
	gm.setupUsers()
	gm.startGame()
	gm.handleTurn()
}
