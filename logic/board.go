package minegw

import (
	"fmt"
	"strings"
	"strconv"
	"os"
	"bufio"
)

type Board struct {
	SizeX int
	SizeY int
	Mines int
	Probability int
	Cells [][]byte
	Table [][]int
}

func Spawn() *Board {
	return new(Board)
}


var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func (board *Board) Create(size string) {
	// Parse board size from the string , following the pattern: Size_XxSize_Y
	sizeArray := strings.Split(size,"x")
	sizeXstring,sizeYstring := sizeArray[0], sizeArray[1]
	// Parse size to integers
	sizeX, err := strconv.Atoi(sizeXstring)
	if err != nil {
		panic("Size X not valid")
	}
	sizeY, err := strconv.Atoi(sizeYstring)
	if err != nil {
		panic("Size Y not valid")
	}
	fmt.Println("Creating board....")
	sizeString := fmt.Sprintf("size_x:%d size_y:%d\n", sizeX,sizeY)
	fmt.Println(sizeString)
	for hyphen := 0; hyphen <= 2+(sizeX*2); hyphen++ {
		fmt.Print("-")
	}
	fmt.Println("")
	board.SizeX, board.SizeY = sizeX, sizeY
	// Set default probability
	board.Probability = 50
	board.createEachCell()
	board.Print()
	board.generateRealBoard()
	board.PrintTable()
	board.Prompt()
}

func (board *Board) Flag(posX,posY int) {
}

func (board *Board) Mine(posX,posY int) {
	board.Change(posX,posY,2)
	board.Print()
}

func (board *Board) Prompt() {
	fmt.Println("Specify Pos: ")
	pos,_ := reader.ReadString('\n')
	pos = strings.TrimRight(pos,"\n")
	posTotal := strings.Split(pos, ",")
	posXstring, posYstring := posTotal[0], posTotal[1]
	posX,err := strconv.Atoi(posXstring)
	if err != nil {
		board.Prompt()
	}
	posY, err := strconv.Atoi(posYstring)
	if err != nil {
		board.Prompt()
	}
	fmt.Println("Select what to do: [M]ine [F]lag")
	option,_ := reader.ReadString('\n')
	option = strings.TrimRight(option, "\n")
	switch strings.ToLower(option) {
	case "m","mine":
		board.Mine(posX, posY)
	case "f","flag":
		board.Flag(posX, posY)
	default:
		board.Prompt()
	}
	board.Prompt()
}
