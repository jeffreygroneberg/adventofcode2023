package ride

import (
	"log"
	"strings"
)

type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

func NewNode(name string) *Node {
	return &Node{Name: name}
}

func travelWithInstructions(root *Node, instructions string, end string) int {

	currentNode := root
	isGoal := false

	sum := 0
	for i := 0; !isGoal; i = ((i + 1) % len(instructions)) {

		currentNode = travelOneStep(currentNode, string(instructions[i]))

		if strings.ContainsAny(currentNode.Name, end) {
			isGoal = true
		}

		// detect loop
		if (currentNode.Left == currentNode) && (currentNode.Right == currentNode) {
			log.Printf("Found a loop at %s", currentNode.Name)
			return sum + 1
		}

		sum++

	}

	return sum
}

func createNodes(nodeLines []string) map[string]*Node {

	createdcreateNodes := make(map[string]*Node)

	// create nodes
	for _, line := range nodeLines {

		// get name
		name := line[0:3]

		// get left and right
		left := line[7:10]
		right := line[12:15]

		// create node
		node := NewNode(name)

		// check if node already exists
		if _, ok := createdcreateNodes[name]; ok {
			node = createdcreateNodes[name]
		} else {
			createdcreateNodes[name] = node
		}

		// check if left node already exists
		if _, ok := createdcreateNodes[left]; ok {
			node.Left = createdcreateNodes[left]
		} else {
			node.Left = NewNode(left)
			createdcreateNodes[left] = node.Left
		}

		// check if right node already exists
		if _, ok := createdcreateNodes[right]; ok {
			node.Right = createdcreateNodes[right]
		} else {
			node.Right = NewNode(right)
			createdcreateNodes[right] = node.Right
		}

	}

	return createdcreateNodes

}

func travelOneStep(node *Node, instruction string) *Node {

	if instruction == "L" {
		return node.Left
	} else {
		return node.Right
	}

}
