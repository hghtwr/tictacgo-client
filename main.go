package main

import (
	"fmt"
	"reflect"
)

type Board struct {
	height int32
	width  int32
}

func (board *Board) printBoard(player1 *Player, player2 *Player) {

	for i := int32(1); i <= board.height; i++ {
		line := "|"
		for e := int32(1); e <= board.width; e++ {
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

type Stones [][]int32 // [x-axis,y-axis]
type Player struct {
	name       string
	stones     Stones
	color      string //
	stonesLeft int32
	symbol     string
}

func (player *Player) getFieldValue(xAxis int32, yAxis int32) bool {
	for i := range player.stones {
		if reflect.DeepEqual(player.stones[i], []int32{xAxis, yAxis}) {
			return true
		}
	}
	return false
}
func (player *Player) setColor(color string) {
	player.color = color
}

func (player *Player) setStone(xAxis int32, yAxis int32) {
	//coordinates := []int32{xAxis, yAxis}
	player.stones = append(player.stones, []int32{xAxis, yAxis})
}

func main() {
	player1 := Player{name: "Joe", stonesLeft: 3, symbol: "x"}
	player2 := Player{name: "Jane", stonesLeft: 3, symbol: "z"}
	player1.setColor("blue")
	player2.setColor("red")
	board := Board{height: 3, width: 3}
	player1.setStone(2, 2)
	player2.setStone(1, 3)
	player1.setStone(2, 1)
	player2.setStone(2, 3)
	player1.setStone(3, 3)
	player2.setStone(1, 1)
	player1.setStone(1, 2)
	player2.setStone(3, 2)

	fmt.Println(player1)
	fmt.Println(player2)
	board.printBoard(&player1, &player2)
}
