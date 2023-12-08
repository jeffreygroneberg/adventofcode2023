package ride

import (
	"strings"
	"sync"
	"testing"

	"github.com/jeffreygroneberg/adventofcode2023/util"
)

func TestCreateNodes(t *testing.T) {
	nodeLines := []string{
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	root := createNodes(nodeLines)["AAA"]

	if root.Name != "AAA" {
		t.Errorf("Expected AAA, but got %s", root.Name)
	}

	if root.Left.Name != "BBB" {
		t.Errorf("Expected BBB, but got %s", root.Left.Name)
	}

	if root.Right.Name != "CCC" {

		t.Errorf("Expected CCC, but got %s", root.Right.Name)
	}

	if root.Left.Left.Name != "DDD" {
		t.Errorf("Expected DDD, but got %s", root.Left.Left.Name)
	}

	if root.Left.Right.Name != "EEE" {

		t.Errorf("Expected EEE, but got %s", root.Left.Right.Name)
	}

	if root.Right.Left.Name != "ZZZ" {

		t.Errorf("Expected ZZZ, but got %s", root.Right.Left.Name)
	}

	if root.Right.Right.Name != "GGG" {

		t.Errorf("Expected GGG, but got %s", root.Right.Right.Name)
	}

	if root.Left.Left.Left.Name != "DDD" {

		t.Errorf("Expected DDD, but got %s", root.Left.Left.Left.Name)
	}

	if root.Left.Left.Right.Name != "DDD" {

		t.Errorf("Expected DDD, but got %s", root.Left.Left.Right.Name)
	}

	if root.Left.Right.Left.Name != "EEE" {

		t.Errorf("Expected EEE, but got %s", root.Left.Right.Left.Name)
	}

	if root.Left.Right.Right.Name != "EEE" {

		t.Errorf("Expected EEE, but got %s", root.Left.Right.Right.Name)
	}

	if root.Right.Left.Left.Name != "ZZZ" {

		t.Errorf("Expected ZZZ, but got %s", root.Right.Left.Left.Name)
	}

	if root.Right.Left.Right.Name != "ZZZ" {

		t.Errorf("Expected ZZZ, but got %s", root.Right.Left.Right.Name)
	}

	if root.Right.Right.Left.Name != "GGG" {

		t.Errorf("Expected GGG, but got %s", root.Right.Right.Left.Name)
	}

	if root.Right.Right.Right.Name != "GGG" {

		t.Errorf("Expected GGG, but got %s", root.Right.Right.Right.Name)
	}

}

func TestTravelWithInstructionsExample1(t *testing.T) {
	nodeLines := []string{

		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	root := createNodes(nodeLines)["AAA"]

	instructions := "RL"

	expected := 2
	result := travelWithInstructions(root, instructions, "ZZZ")

	if result != expected {

		t.Errorf("Expected %d, but got %d", expected, result)

	}
}

func TestTravelWithInstructionsData1(t *testing.T) {
	nodeLines, _ := util.ReadFile("testdata/part1.txt")

	root := createNodes(nodeLines)["AAA"]

	instructions := "LRLRRLRLRLLRRRLLLRLLRRLLRRRLRLRLRRRLRRLRLRRRLRRRLRRRLRRRLRRLRRRLRRRLRRLLLRLLRLRRRLRRRLRRLRLRLRLRRRLRRRLRRRLRRLRRLRRLLRRLRRRLLRRLRRRLRRRLRRRLRLRRLRLRRRLRRLLRLRLLRLRLRRRLRLRRLLRRRLRLRLRLRLRLRRLRLRRLLLLRRLRRLRRRLRRLRRLRRRLRRLRRRLLRLRRLLRLRRLRRLRRLLRRRLRLRLRRRLRRLRLLRLRRRR"

	expected := 19637
	result := travelWithInstructions(root, instructions, "ZZZ")

	if result != expected {

		t.Errorf("Expected %d, but got %d", expected, result)

	}
}

func TestTravelWithInstructionsData2(t *testing.T) {

	nodeLines, _ := util.ReadFile("testdata/part2.txt")

	nodes := createNodes(nodeLines)

	nodesWithAMap := make(map[string]*Node)
	var nodesWithASlice []*Node

	for _, node := range nodes {
		if strings.Contains(node.Name, "A") {
			nodesWithAMap[node.Name] = node
			nodesWithASlice = append(nodesWithASlice, node)
		}
	}

	expected := 8811050362409
	instruction := "LRLRRLRLRLLRRRLLLRLLRRLLRRRLRLRLRRRLRRLRLRRRLRRRLRRRLRRRLRRLRRRLRRRLRRLLLRLLRLRRRLRRRLRRLRLRLRLRRRLRRRLRRRLRRLRRLRRLLRRLRRRLLRRLRRRLRRRLRRRLRLRRLRLRRRLRRLLRLRLLRLRLRRRLRLRRLLRRRLRLRLRLRLRLRRLRLRRLLLLRRLRRLRRRLRRLRRLRRRLRRLRRRLLRLRRLLRLRRLRRLRRLLRRRLRLRLRRRLRRLRLLRLRRRR"

	sumsChan := make(chan int, len(nodesWithASlice))
	var wg sync.WaitGroup

	for _, node := range nodesWithASlice {
		wg.Add(1)
		go func(node Node) {
			defer wg.Done()
			sum := travelWithInstructions(&node, instruction, "Z")
			sumsChan <- sum
		}(*node)
	}

	go func() {
		wg.Wait()
		close(sumsChan)
	}()

	var sums []int
	for sum := range sumsChan {
		sums = append(sums, sum)
	}

	// calculate LCM of all sums
	lcm := LCM(sums[0], sums[1], sums[2:]...)

	if lcm != expected {
		t.Errorf("Expected %d, but got %d", expected, lcm)
	}

}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
