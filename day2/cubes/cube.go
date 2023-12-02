package cubegame

import (
	"strconv"
	"strings"
)

func ExtractGameIdAndDraws(line string) (int, string) {

	parts := strings.Split(line, ":")
	gameId, err := strconv.Atoi(strings.TrimSpace(parts[0][4:]))

	if err != nil {
		panic(err)
	}

	return gameId, strings.TrimSpace(parts[1])

}

func GetCubesForSingleDraw(draws string) map[string]int {

	cubeCount := make(map[string]int)

	drawnCubes := strings.Split(draws, ",")

	for _, cube := range drawnCubes {

		cube = strings.TrimSpace(cube)
		cubeParts := strings.Split(cube, " ")

		cubeNumber, err := strconv.Atoi(cubeParts[0])
		if err != nil {
			panic(err)
		}

		cubeColor := cubeParts[1]
		cubeCount[cubeColor] = cubeNumber

	}

	return cubeCount
}

func GetCubesForAllDraws(draws string) map[string]int {

	cubeCount := make(map[string]int)

	draws = strings.TrimSpace(draws)
	cubesPerDraw := strings.Split(draws, ";")

	for _, draw := range cubesPerDraw {

		cubes := GetCubesForSingleDraw(draw)

		for color, number := range cubes {
			if number > cubeCount[color] {
				cubeCount[color] = number
			}
		}

	}

	return cubeCount
}

func GetCubesForAllGames(games []string) map[int]map[string]int {

	cubeCount := make(map[int]map[string]int)

	for _, game := range games {

		gameId, draws := ExtractGameIdAndDraws(game)
		cubes := GetCubesForAllDraws(draws)
		cubeCount[gameId] = cubes

	}

	return cubeCount

}

func GetGamesWithConstraints(games []string, constraints map[string]int, filterByValid bool) map[int]map[string]int {

	validGames := map[int]map[string]int{}

	cubes := GetCubesForAllGames(games)

	for gameId, gameCubes := range cubes {

		if isGameValid(gameCubes, constraints) == filterByValid {
			validGames[gameId] = gameCubes
		}
	}

	return validGames

}

func isGameValid(gameCubes map[string]int, constraints map[string]int) bool {
	valid := true
	for color, number := range gameCubes {
		if number > constraints[color] {
			valid = false
		}
	}
	return valid
}

func GetPowerOfCubes(cubes map[string]int) int {
	power := 1
	for _, number := range cubes {
		power *= number
	}
	return power
}
