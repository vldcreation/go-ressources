package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

var directions = []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func main() {
	// Open input file
	inputFile, err := os.Open("input3.in")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Open output file for writing
	outputFile, err := os.Create("input3.out")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Redirect standard input to read from the input file
	stdin := os.Stdin
	os.Stdin = inputFile
	defer func() {
		os.Stdin = stdin
	}()

	// Redirect standard output to write to the output file
	stdout := os.Stdout
	os.Stdout = outputFile
	defer func() {
		os.Stdout = stdout
		outputFile.Sync()
	}()

	T := 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &T)

	for caseNum := 1; caseNum <= T; caseNum++ {
		scanner.Scan()
		N, M := 0, 0
		fmt.Sscanf(scanner.Text(), "%d", &N)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d", &M)

		grid := make([][]rune, N)
		for i := 0; i < N; i++ {
			scanner.Scan()
			grid[i] = []rune(scanner.Text())
		}

		factionRegions := make(map[rune]int)
		contestedRegions := 0
		visited := make([][]bool, N)
		for i := range visited {
			visited[i] = make([]bool, M)
		}

		var dfs func(x, y int) (map[rune]struct{}, bool)
		dfs = func(x, y int) (map[rune]struct{}, bool) {
			stack := []Point{{x, y}}
			factions := make(map[rune]struct{})
			isContested := false

			for len(stack) > 0 {
				p := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if visited[p.x][p.y] {
					continue
				}
				visited[p.x][p.y] = true

				if grid[p.x][p.y] == '#' {
					continue
				}

				if grid[p.x][p.y] >= 'a' && grid[p.x][p.y] <= 'z' {
					factions[grid[p.x][p.y]] = struct{}{}
				}

				for _, d := range directions {
					nx, ny := p.x+d.x, p.y+d.y
					if nx >= 0 && nx < N && ny >= 0 && ny < M && !visited[nx][ny] {
						stack = append(stack, Point{nx, ny})
					}
				}
			}

			if len(factions) > 1 {
				isContested = true
			}
			return factions, isContested
		}

		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				if !visited[i][j] && grid[i][j] != '#' {
					factionsInRegion, isContested := dfs(i, j)
					if isContested {
						contestedRegions++
					} else if len(factionsInRegion) == 1 {
						for faction := range factionsInRegion {
							factionRegions[faction]++
						}
					}
				}
			}
		}

		fmt.Printf("Case %d:\n", caseNum)

		// Sort and print factions controlling regions
		for faction := 'a'; faction <= 'z'; faction++ {
			if count, exists := factionRegions[faction]; exists {
				fmt.Printf("%c %d\n", faction, count)
			}
		}

		fmt.Printf("contested %d\n", contestedRegions)
	}
}
