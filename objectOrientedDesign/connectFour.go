package main

import (
	"errors"
	"fmt"
)

type GridPosition int

const (
	EMPTY GridPosition = iota
	YELLOW
	RED
)

type Grid struct {
	rows, columns int
	grid          [][]GridPosition
}

func NewGrid(rows, columns int) *Grid {
	grid := make([][]GridPosition, rows)
	for i := range grid {
		grid[i] = make([]GridPosition, columns)
	}
	return &Grid{
		rows:    rows,
		columns: columns,
		grid:    grid,
	}
}

func (g *Grid) PlacePiece(column int, piece GridPosition) (int, error) {
	if column < 0 || column >= g.columns {
		return -1, errors.New("invalid column")
	}
	if piece == EMPTY {
		return -1, errors.New("invalid piece")
	}
	for row := g.rows - 1; row >= 0; row-- {
		if g.grid[row][column] == EMPTY {
			g.grid[row][column] = piece
			return row, nil
		}
	}
	return -1, nil
}

func (g *Grid) CheckWin(connectN, row, col int, piece GridPosition) bool {
	// Check horizontal
	count := 0
	for c := 0; c < g.columns; c++ {
		if g.grid[row][c] == piece {
			count++
		} else {
			count = 0
		}
		if count == connectN {
			return true
		}
	}

	// Check vertical
	count = 0
	for r := 0; r < g.rows; r++ {
		if g.grid[r][col] == piece {
			count++
		} else {
			count = 0
		}
		if count == connectN {
			return true
		}
	}

	// Check diagonal
	count = 0
	for r := 0; r < g.rows; r++ {
		c := row + col - r
		if c >= 0 && c < g.columns && g.grid[r][c] == piece {
			count++
		} else {
			count = 0
		}
		if count == connectN {
			return true
		}
	}

	// Check anti-diagonal
	count = 0
	for r := 0; r < g.rows; r++ {
		c := col - row + r
		if c >= 0 && c < g.columns && g.grid[r][c] == piece {
			count++
		} else {
			count = 0
		}
		if count == connectN {
			return true
		}
	}
	return false
}

type Player struct {
	name       string
	pieceColor GridPosition
}

func NewPlayer(name string, pieceColor GridPosition) *Player {
	return &Player{
		name:       name,
		pieceColor: pieceColor,
	}
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetPieceColor() GridPosition {
	return p.pieceColor
}

type Game struct {
	grid        *Grid
	connectN    int
	players     []*Player
	score       map[string]int
	targetScore int
}

func NewGame(grid *Grid, connectN, targetScore int) *Game {
	players := []*Player{
		NewPlayer("Player 1", YELLOW),
		NewPlayer("Player 2", RED),
	}
	score := make(map[string]int)
	for _, player := range players {
		score[player.GetName()] = 0
	}
	return &Game{
		grid:        grid,
		connectN:    connectN,
		players:     players,
		score:       score,
		targetScore: targetScore,
	}
}

func (g *Game) PlayMove(player *Player) (int, int) {
	g.PrintBoard()
	fmt.Printf("%s's turn\n", player.GetName())
	moveColumn := 0
	fmt.Printf("Enter column between %d and %d to add piece: ", 0, g.grid.columns-1)
	fmt.Scanln(&moveColumn)
	moveRow, _ := g.grid.PlacePiece(moveColumn, player.GetPieceColor())
	return moveRow, moveColumn
}

func (g *Game) PlayRound() *Player {
	for {
		for _, player := range g.players {
			row, col := g.PlayMove(player)
			pieceColor := player.GetPieceColor()
			if g.grid.CheckWin(g.connectN, row, col, pieceColor) {
				g.score[player.GetName()]++
				return player
			}
		}
	}
}

func (g *Game) PrintBoard() {
	fmt.Println("Board:")
	for _, row := range g.grid.grid {
		for _, piece := range row {
			switch piece {
			case EMPTY:
				fmt.Print("0 ")
			case YELLOW:
				fmt.Print("Y ")
			case RED:
				fmt.Print("R ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *Game) Play() {
	maxScore := 0
	var winner *Player
	for maxScore < g.targetScore {
		winner = g.PlayRound()
		fmt.Printf("%s won the round\n", winner.GetName())
		maxScore = Max(g.score[winner.GetName()], maxScore)
		g.grid = NewGrid(g.grid.rows, g.grid.columns) // reset grid
	}
	fmt.Printf("%s won the game\n", winner.GetName())
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mainCF() {
	grid := NewGrid(6, 7)
	game := NewGame(grid, 4, 2)
	game.Play()
}
