package minegw

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func (board *Board) createEachCell() {
	// Allocate a 0 to each empty cell (will add mines later)
	var totalCells [][]byte
	for row := 0; row < board.SizeX; row++ {
		var rowContent []byte
		for column := 0; column < board.SizeY; column++ {
			rowContent = append(rowContent, board.generateMine())
		}
		totalCells = append(totalCells, rowContent)
	}
	board.Cells = totalCells
}

func (board *Board) generateMine() byte {
	// return board.randInt(board.getSize())
	probability := board.randInt(100)
	if probability < board.Probability {
		return 1
	}
	return 0
}

func (board *Board) randInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func (board *Board) getSize() int {
	return board.SizeX*board.SizeY
}

func (board *Board) RevealBoard() {
	fmt.Println(board.Cells)
}

func (board *Board) getPosition(posX,posY int) byte{
	row := board.Cells[posY]
	value := row[posX]
	return value
}

func (board *Board) Change(posX,posY int, newVal byte) {
	board.Cells[posY][posX] = newVal
}

func (board *Board) Print() {
	fmt.Print("   ")
	for i := 0; i < board.SizeY; i++ {
		fmt.Print(" "+strconv.Itoa(i))
	}
	fmt.Println(" ")
	for index, row := range board.Cells {
		fmt.Print(index)
		fmt.Print(": ")
		fmt.Println(row)
	}
}

func (board *Board) checkValidCell(posX,posY int) (bool, bool) {
	var posXvalid bool
	var posYvalid bool
	if posX < 0 ||  posX > board.SizeY {
		posXvalid = false
	} else {
		posXvalid = true
	}

	if posY < 0 ||  posY > board.SizeX {
		posYvalid = false
	} else {
		posYvalid = true
	}
	return posXvalid, posYvalid
}

func (board *Board) locateNearMines(posX,posY int) int {
	// X axis:
	if posX != 0 && posY != 0 && posX != board.SizeX && posY != board.SizeY {
		originX := posX - 1
		originY := posY - 1
		var i int
		var j int
		var around [][]byte
		var rowX []byte
		for j = 0; j < 3 ; j++ {
			for i = 0; i < 3 ; i++ {
				rowX = append(rowX, board.getPosition(originX+i, originY+j))
			}
			around = append(around, rowX)
			fmt.Println(rowX)
			rowX = nil
		}
	}
	return 0
}

func (board *Board) generateRealBoard() {
	// Iterate through each byte array
}
