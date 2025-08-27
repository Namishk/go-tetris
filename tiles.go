package main

type Tile struct {
	Tile      [][]int
	TilePrime [][]int
}

var TileO = Tile{
	Tile: [][]int{
		{1, 1},
		{1, 1},
	},

	TilePrime: [][]int{
		{1, 1},
		{1, 1},
	},
}

var TileI = Tile{
	Tile: [][]int{
		{1, 1, 1, 1},
	},

	TilePrime: [][]int{
		{1},
		{1},
		{1},
		{1},
	},
}

var TileT = Tile{
	Tile: [][]int{
		{0, 1, 0},
		{1, 1, 1},
	},

	TilePrime: [][]int{
		{1, 0},
		{1, 1},
		{1, 0},
	},
}

var TileL = Tile{
	Tile: [][]int{
		{1, 0},
		{1, 0},
		{1, 1},
	},

	TilePrime: [][]int{
		{1, 1, 1},
		{1, 0, 0},
	},
}

var TileJ = Tile{
	Tile: [][]int{
		{0, 1},
		{0, 1},
		{1, 1},
	},

	TilePrime: [][]int{
		{1, 0, 0},
		{1, 1, 1},
	},
}

var TileS = Tile{
	Tile: [][]int{
		{0, 1, 1},
		{1, 1, 0},
	},

	TilePrime: [][]int{
		{1, 0},
		{1, 1},
		{0, 1},
	},
}

var TileZ = Tile{
	Tile: [][]int{
		{1, 1, 0},
		{0, 1, 1},
	},

	TilePrime: [][]int{
		{0, 1},
		{1, 1},
		{1, 0},
	},
}
