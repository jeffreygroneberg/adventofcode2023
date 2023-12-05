package seed

import (
	"reflect"
	"testing"

	"github.com/jeffreygroneberg/adventofcode2023/util"
)

func TestCreateMappingTable(t *testing.T) {

	tests := []struct {
		name     string
		values   []string
		expected MappingTable
	}{
		{
			name:   "seed-to-soil map:",
			values: []string{"1 2 3", "4 5 6", "7 8 9"},
			expected: MappingTable{
				SourceName:      "seed",
				DestinationName: "soil",
				Destination:     []int{1, 4, 7},
				Source:          []int{2, 5, 8},
				Range:           []int{3, 6, 9},
			},
		},
		{
			name:   "fertilizer-to-water map:",
			values: []string{"10 20 30", "40 50 60", "70 80 90"},
			expected: MappingTable{
				SourceName:      "fertilizer",
				DestinationName: "water",
				Destination:     []int{10, 40, 70},
				Source:          []int{20, 50, 80},
				Range:           []int{30, 60, 90},
			},
		},
		// Add more test cases here...
	}

	for _, test := range tests {
		result := createMappingTable(test.name, test.values)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test Case: %s\nExpected: %+v\nGot: %+v", test.name, test.expected, result)
		}
	}
}

func TestMapTo(t *testing.T) {
	table := MappingTable{
		SourceName:      "seed",
		DestinationName: "soil",

		Destination: []int{50, 52},
		Source:      []int{98, 50},
		Range:       []int{2, 48},
	}

	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "input within range",
			input:    79,
			expected: 81,
		},
		{
			name:     "input within range",
			input:    14,
			expected: 14,
		},
		{
			name:     "input within range",
			input:    55,
			expected: 57,
		},
	}

	for _, test := range tests {
		result := mapTo(test.input, table)

		if result != test.expected {
			t.Errorf("Test Case: %s\nExpected: %d\nGot: %d", test.name, test.expected, result)
		}
	}
}

func TestCreateMappingTables(t *testing.T) {

	lines := []string{
		"seed-to-soil map:",
		"1 2 3",
		"4 5 6",
		"7 8 9",
		"",
		"fertilizer-to-water map:",
		"10 20 30",
		"40 50 60",
		"70 80 90",
		"",
		"an-to-other map:",
		"100 200 300",
		"400 500 600",
		"700 800 900",
		"",
	}

	expectedTables := []MappingTable{
		{
			SourceName:      "seed",
			DestinationName: "soil",
			Destination:     []int{1, 4, 7},
			Source:          []int{2, 5, 8},
			Range:           []int{3, 6, 9},
		},
		{
			SourceName:      "fertilizer",
			DestinationName: "water",
			Destination:     []int{10, 40, 70},
			Source:          []int{20, 50, 80},
			Range:           []int{30, 60, 90},
		},
		{
			SourceName:      "an",
			DestinationName: "other",
			Destination:     []int{100, 400, 700},
			Source:          []int{200, 500, 800},
			Range:           []int{300, 600, 900},
		},
	}

	tables := CreateMappingTables(lines)

	if len(tables) != len(expectedTables) {
		t.Errorf("Expected %d tables, got %d", len(expectedTables), len(tables))
	}

	for i, expected := range expectedTables {
		if !reflect.DeepEqual(tables[i], expected) {
			t.Errorf("Table %d:\nExpected: %+v\nGot: %+v", i+1, expected, tables[i])
		}
	}

}

func TestMapToWithPipeExample1(t *testing.T) {

	lines := []string{
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
		"",
	}

	expectedOutput := 35
	tables := CreateMappingTables(lines)

	inputs := []int{79, 14, 55, 13}
	expectedOutputs := []int{82, 43, 86, 35}

	mappingTablesMap := make(map[string]MappingTable)

	for i, table := range tables {
		mappingTablesMap[table.SourceName] = tables[i]
	}

	lowestOutput := 0

	for i, input := range inputs {

		output := mapToWithPipe(input, "seed", "location", mappingTablesMap)

		if output != expectedOutputs[i] {
			t.Errorf("Expected %d, got %d", expectedOutputs[i], output)
		}
		if lowestOutput == 0 || output < lowestOutput {
			lowestOutput = output
		}

	}

	if lowestOutput != expectedOutput {
		t.Errorf("Expected %d, got %d", expectedOutput, lowestOutput)
	}

}

func TestMapToWithPipeData1(t *testing.T) {

	inputs := []int{1778931867, 1436999653, 3684516104, 2759374, 1192793053, 358764985, 1698790056, 76369598, 3733854793, 214008036, 4054174000, 171202266, 3630057255, 25954395, 798587440, 316327323, 290129780, 7039123, 3334326492, 246125391}
	lines, _ := util.ReadFile("testdata/part1.txt")
	tables := CreateMappingTables(lines)

	expectedOutput := 107430936
	mappingTablesMap := make(map[string]MappingTable)

	for i, table := range tables {
		mappingTablesMap[table.SourceName] = tables[i]
	}

	lowestOutput := 0
	for _, input := range inputs {

		output := mapToWithPipe(input, "seed", "location", mappingTablesMap)
		if lowestOutput == 0 || output < lowestOutput {
			lowestOutput = output
		}

	}

	if lowestOutput != expectedOutput {
		t.Errorf("Expected %d, got %d", expectedOutput, lowestOutput)
	}

}

func TestCreateSeedRanges(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected []struct {
			from        int
			rangeNumber int
		}
	}{
		{
			name: "test case 1",
			line: "79 80 55 56",
			expected: []struct {
				from        int
				rangeNumber int
			}{
				{
					from:        79,
					rangeNumber: 80,
				},
				{
					from:        55,
					rangeNumber: 56,
				},
			},
		},
	}

	for _, test := range tests {
		result := CreateSeedRanges(test.line)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test Case: %s\nExpected: %+v\nGot: %+v", test.name, test.expected, result)
		}
	}
}

func TestMapToWithPipeExample2(t *testing.T) {

	lines := []string{
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
		"",
	}

	expectedOutput := 46
	tables := CreateMappingTables(lines)

	seedRanges := CreateSeedRanges("79 14 55 13")

	mappingTablesMap := make(map[string]MappingTable)

	for i, table := range tables {
		mappingTablesMap[table.SourceName] = tables[i]
	}

	lowestOutput := 0

	for _, seedRange := range seedRanges {

		for i := seedRange.from; i <= seedRange.from+seedRange.rangeNumber; i++ {

			output := mapToWithPipe(i, "seed", "location", mappingTablesMap)

			if lowestOutput == 0 || output < lowestOutput {
				lowestOutput = output
			}

		}

	}

	if lowestOutput != expectedOutput {
		t.Errorf("Expected %d, got %d", expectedOutput, lowestOutput)
	}

}

func TestMapToWithPipeData2(t *testing.T) {

	lines, _ := util.ReadFile("testdata/part1.txt")
	tables := CreateMappingTables(lines)

	expectedOutput := 107430936
	mappingTablesMap := make(map[string]MappingTable)

	for i, table := range tables {
		mappingTablesMap[table.SourceName] = tables[i]
	}
	lowestOutput := 0

	seedRanges := CreateSeedRanges("1778931867 1436999653 3684516104 2759374 1192793053 358764985 1698790056 76369598 3733854793 214008036 4054174000 171202266 3630057255 25954395 798587440 316327323 290129780 7039123 3334326492 246125391")

	for _, seedRange := range seedRanges {

		for i := seedRange.from; i <= seedRange.from+seedRange.rangeNumber; i++ {

			output := mapToWithPipe(i, "seed", "location", mappingTablesMap)

			if lowestOutput == 0 || output < lowestOutput {
				lowestOutput = output
			}

		}

	}

	if lowestOutput != expectedOutput {
		t.Errorf("Expected %d, got %d", expectedOutput, lowestOutput)
	}

}
