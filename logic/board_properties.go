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
	var count int
	var originX int
	var originY int
	var i int
	var j int
	var max_i int
	var max_j int

	if posX != 0 && posY != 0 && posX != board.SizeX - 1 && posY != board.SizeY - 1 {
		originX = posX - 1
		originY = posY - 1
		max_i , max_j = 3,3
	} else {
		if posX == 0{
			originX = posX
			max_i = 2
		} else if posX == board.SizeX - 1 {
			originX = posX - 1
			max_i = 2
		} else {
			originX = posX - 1
			max_i = 3
		}
		if posY == 0 {
			originY = posY
			max_j = 2
		} else if posY == board.SizeY - 1 {
			originY = posY - 1
			max_j = 2
		} else {
			originY = posY - 1
			max_j = 3
		}
	}

	for j = 0; j < max_j ; j++ {
		for i = 0; i < max_i ; i++ {
			if originX+i == posX && originY+j == posY {
				continue
			} else if board.getPosition(originX+i, originY+j) == 1 {
				count++
			}
		}
	}

	return count
}

func (board *Board) generateRealBoard() {
	// Iterate through each byte array

}
