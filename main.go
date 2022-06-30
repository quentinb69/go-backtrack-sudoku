package main

import ( 
  "fmt"
  "os"
  "bufio"
  "strconv"
)

const (
  SIZE = 9
  SQUARE_SIZE = 3
 )
 
type Sudoku [SIZE][SIZE]int

/* used for coding game */

// Print Sudoku grid
func (s Sudoku) String() string {
	str := ""
	for i, _ := range s {
		for j, _ := range s[i] {
			str += strconv.Itoa(s[i][j])
		}
		if i <= len(s) -2 {
			str += "\n"
		}
	}
	return str
}

// Log in stderr
func log(s... interface{}) {
	fmt.Fprintln(os.Stderr, s)
}

// validate a line
func validLine(s *Sudoku, x, val int) bool {
	for col, _ := range (s[x]) {
		if val == s[x][col] {
			return false
		}
	}
	return true
}

// validate a column
func validCol(s *Sudoku, y, val int) bool {
	for line, _ := range (s) {
		if val == s[line][y] {
			return false
		}
	}
	return true
}

// get top left position in square and validate unicity of val in whole square
func validSquare(s *Sudoku, x, y, val int) bool {
	topLeftX:= int(x/SQUARE_SIZE) * SQUARE_SIZE
	topLeftY:= int(y/SQUARE_SIZE) * SQUARE_SIZE
	for i := 0 ; i < SQUARE_SIZE ; i ++ {
		for j := 0 ; j < SQUARE_SIZE ; j++ {
			if s[i+topLeftX][j+topLeftY] == val {
				return false
			}
		}
	}
	return true
}

// find free case (containing 0)
func findfree(s* Sudoku) (int, int) {
	for i, _ := range (s){
		for j, _ := range (s[i]){
			if s[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

// recursive backtracking
func backtrack(s Sudoku) (bool) {
	x, y := findfree(&s)

	// full
	if x==-1 && y==-1 {
		fmt.Println( s)
		return true
	}

	for nb := 1; nb <= SIZE; nb++ {
		if validLine(&s, x,nb) && validCol(&s, y, nb) && validSquare(&s, x, y, nb) {
			s[x][y] = nb
			if ret := backtrack (s); ret {
				return true
			}
			// if not ok then revert
			s[x][y] = 0
		}
	}
	// if nothing work backtrack
	return false
}

// main for coding game
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)
 
	grid := Sudoku{}

    for i := 0; i < SIZE; i++ {
        scanner.Scan()
        line := scanner.Text()
		for j, v := range []rune(line) {
			grid[i][j] = int(v -'0')
		}
    }

	log( "INPUT")
	log( grid)	

	backtrack( grid)
}
