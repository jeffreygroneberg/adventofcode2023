package seed

import (
	"strconv"
	"strings"
)

type MappingTable struct {
	SourceName      string
	DestinationName string
	Destination     []int
	Source          []int
	Range           []int
}

func mapTo(input int, table MappingTable) int {

	for i := 0; i < len(table.Source); i++ {

		sourceValue := table.Source[i]
		destinationValue := table.Destination[i]
		rangeValue := table.Range[i]

		if input >= sourceValue && input <= sourceValue+rangeValue {
			// determine gap source and input
			gapSource := input - sourceValue
			return destinationValue + gapSource
		}
	}

	return input

}

func mapToWithPipe(input int, startMapName string, endMapName string, maps map[string]MappingTable) int {

	currentMap := maps[startMapName]

	for {

		input = mapTo(input, currentMap)
		currentMap = maps[currentMap.DestinationName]

		if currentMap.DestinationName == endMapName {

			// map one last time
			input = mapTo(input, currentMap)
			break
		}
	}

	return input

}

func createMappingTable(mapTitle string, values []string) MappingTable {

	firstPart := strings.Split(mapTitle, " ")[0]
	fullNameSplit := strings.Split(firstPart, "-")

	sourceName := fullNameSplit[0]
	destinationName := fullNameSplit[2]

	table := MappingTable{

		SourceName:      sourceName,
		DestinationName: destinationName,
		Destination:     make([]int, len(values)),
		Source:          make([]int, len(values)),
		Range:           make([]int, len(values)),
	}

	for i, value := range values {

		// trim white space
		value = strings.TrimSpace(value)
		parts := strings.Split(value, " ")

		// store in Mapping table
		table.Destination[i], _ = strconv.Atoi(parts[0])
		table.Source[i], _ = strconv.Atoi(parts[1])
		table.Range[i], _ = strconv.Atoi(parts[2])

	}

	return table

}

func CreateMappingTables(lines []string) []MappingTable {

	var tables []MappingTable

	for i, line := range lines {

		// trim white space
		line = strings.TrimSpace(line)

		// skip empty lines
		if line == "" {
			continue
		}

		// check if line is a map title
		if strings.HasSuffix(line, " map:") {

			// now we need to read the next lines until we reach an empty line
			var values []string

			for j := i + 1; j < len(lines); j++ {

				// trim white space
				line = strings.TrimSpace(lines[j])

				// if we reach an empty line, we are done
				if line == "" {
					break
				}

				values = append(values, line)

			}

			// create mapping table
			table := createMappingTable(lines[i], values)
			// store in slice
			tables = append(tables, table)

		}

	}
	return tables
}

func CreateSeedRanges(line string) []struct {
	from        int
	rangeNumber int
} {

	// seeds: 79 14 55 13 would create two seed pairs: 79-80 and 55-56

	// trim white space
	line = strings.TrimSpace(line)

	// split by whitespace
	parts := strings.Split(line, " ")

	seedRanges := make([]struct {
		from        int
		rangeNumber int
	}, 0)

	for i := 0; i < len(parts); i += 2 {

		from, _ := strconv.Atoi(parts[i])
		rangeNumber, _ := strconv.Atoi(parts[i+1])

		// create seed pair
		seedPair := struct {
			from        int
			rangeNumber int
		}{
			from:        from,
			rangeNumber: rangeNumber,
		}

		// store in slice
		seedRanges = append(seedRanges, seedPair)

	}

	return seedRanges
}
