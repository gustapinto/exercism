package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	rankSquares, isRankOnMap := cb[rank]
	if !isRankOnMap {
		return 0
	}

	counter := 0
	for _, isSquareOccupied := range rankSquares {
		if isSquareOccupied {
			counter += 1
		}
	}

	return counter
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	if file < 1 || file > 8 {
		return 0
	}

	counter := 0
	for _, rank := range cb {
		isSquareOcuppied := rank[file-1]
		if isSquareOcuppied {
			counter += 1
		}
	}

	return counter
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	counter := 0
	for _, _ = range cb {
		counter += 8
	}

	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	counter := 0
	for key, _ := range cb {
		counter += CountInRank(cb, key)
	}

	return counter
}
