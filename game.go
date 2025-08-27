package main

import (
	"fmt"
)

var defaultScreen = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
}

var defaultTile = [][]int{
	{0},
}

type gameState int

const (
	intro gameState = iota
	playing
	paused
	gameOver
)

type Game struct {
	score       int
	speed       int
	board       [][]int
	rowsCleared int
	tile        [][]int
	state       gameState
}

// default UI 10x20 : WxH

func NewGame() *Game {
	g := new(Game)
	g.reset()

	return g
}

func (g *Game) reset() {
	g.score = 0
	g.speed = 1
	g.board = defaultScreen
	g.rowsCleared = 0
	g.tile = defaultTile
	g.state = intro

}

func (g *Game) renderGame() {
	fmt.Print("\x1b[2J") // clear screen
	fmt.Print("\x1b[H")  // bring cursor to start so we can redraw ui
	for _, v := range g.board {
		for _, k := range v {
			fmt.Printf("%d ", k)
		}
		fmt.Print("\r\n")
	}
}

func (g *Game) rotateTile() {
	g.renderGame()
	fmt.Println("ROTATE")
}

func (g *Game) goDown() {
	g.renderGame()
	fmt.Println("DOWN")
}

func (g *Game) moveLeft() {
	g.renderGame()
	fmt.Println("LEFT")
}

func (g *Game) MoveRight() {
	g.renderGame()
	fmt.Println("RIGHT")
}

func (g *Game) exit() {
	fmt.Print("\x1b[2J") // clear screen
	fmt.Print("\x1b[H")  // bring cursor to start so we can redraw ui
	fmt.Println("EXIT")
}
