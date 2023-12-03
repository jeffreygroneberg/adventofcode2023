package calibration

import (
	"strconv"
	"strings"

	"github.com/jeffreygroneberg/adventofcode2023/util"
)

var Replacer = strings.NewReplacer(

	// edge cases with overlapping first

	"oneight", "18",
	"twone", "21",
	"threeight", "38",
	"fiveight", "58",
	"sevenine", "79",
	"eightwo", "82",
	"eighthree", "83",
	"nineight", "98",

	"one", "1",
	"eight", "8",
	"five", "5",
	"four", "4",
	"three", "3",
	"two", "2",
	"seven", "7",
	"six", "6",
	"nine", "9",
)

// CalculateCalibrationForLine calculates the calibration value for a given line.
// If replace is true, it replaces certain characters in the line before calculating the calibration.
// It loops through each character in the line and adds the numbers to a string.
// Finally, it returns the sum of the first and last digit of the resulting string as an integer.
// If there is an error during the conversion, it returns an error.
func CalculateCalibrationForLine(line string, replace bool) (int, error) {

	var builder strings.Builder

	if replace {
		line = Replacer.Replace(line)
	}

	// loop through each character in the string
	for i := 0; i < len(line); i++ {
		// if the character is a number, add it to the string
		if util.IsNumber(line[i]) {
			builder.WriteByte(line[i])
		}
	}

	return strconv.Atoi(builder.String()[0:1] + builder.String()[builder.Len()-1:])

}
