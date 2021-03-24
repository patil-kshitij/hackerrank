package main

//https://www.hackerrank.com/challenges/swap-nodes-algo/problem
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the swapNodes function below.
 */
func swapNodes(indexes [][]int32, queries []int32) [][]int32 {
	root, depth := createTree(indexes)
	result := make([][]int32, len(queries), len(queries))
	for index, level := range queries {
		for i := int32(1); i <= depth; i++ {
			InorderSwapNode(root, i*level, 1)
		}

		InorderTraversal(root)
		inorderTraversal := traversal
		result[index] = inorderTraversal
		traversal = make([]int32, 0, 1)
	}
	return result
}

var traversal = make([]int32, 0, 1)

func InorderSwapNode(node *Node, level, currentLevel int32) {
	if node == nil || currentLevel > level {
		return
	}

	InorderSwapNode(node.Left, level, currentLevel+1)
	InorderSwapNode(node.Right, level, currentLevel+1)
	if level == currentLevel {
		temp := node.Left
		node.Left = node.Right
		node.Right = temp
	}
}

func InorderTraversal(node *Node) {
	if node == nil {
		return
	}
	InorderTraversal(node.Left)
	traversal = append(traversal, node.Index)
	InorderTraversal(node.Right)
}

type Node struct {
	Index int32
	Left  *Node
	Right *Node
}

func createTree(indexes [][]int32) (*Node, int32) {
	rootNode := &Node{
		Index: int32(1),
		Left:  nil,
		Right: nil,
	}
	depth := int32(1)
	queue := make([]*Node, 0, 1)
	queue = append(queue, rootNode)
	childAdded := false
	for _, children := range indexes {
		parentNode := queue[0]
		queue = queue[1:]
		var leftChild, rightChild *Node
		if children[0] != -1 {
			leftChild = &Node{
				Index: children[0],
			}
			queue = append(queue, leftChild)
			childAdded = true
		}
		if children[1] != -1 {
			rightChild = &Node{
				Index: children[1],
			}
			queue = append(queue, rightChild)
			childAdded = true
		}

		parentNode.Left = leftChild
		parentNode.Right = rightChild

		if childAdded {
			depth = depth + 1
		}
	}

	return rootNode, depth

}

func main() {
	f, _ := os.Open("trial.txt")
	reader := bufio.NewReaderSize(f, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var indexes [][]int32
	for indexesRowItr := 0; indexesRowItr < int(n); indexesRowItr++ {
		indexesRowTemp := strings.Split(readLine(reader), " ")

		var indexesRow []int32
		for _, indexesRowItem := range indexesRowTemp {
			indexesItemTemp, err := strconv.ParseInt(indexesRowItem, 10, 64)
			checkError(err)
			indexesItem := int32(indexesItemTemp)
			indexesRow = append(indexesRow, indexesItem)
		}

		if len(indexesRow) != int(2) {
			panic("Bad input")
		}

		indexes = append(indexes, indexesRow)
	}

	queriesCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var queries []int32

	for queriesItr := 0; queriesItr < int(queriesCount); queriesItr++ {
		queriesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := swapNodes(indexes, queries)

	for resultRowItr, rowItem := range result {
		for resultColumnItr, colItem := range rowItem {
			fmt.Fprintf(writer, "%d", colItem)

			if resultColumnItr != len(rowItem)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		if resultRowItr != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
