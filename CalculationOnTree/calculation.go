package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Data       int32
	Neighbours map[int32]*Node
}

var NodeMap = make(map[int32]*Node)

func kittyCalculation(edgeList, queries [][]int32) []int32 {
	CreateGraph(edgeList)
	result := make([]int32, 0, 1)
	for _, query := range queries {
		pairs := getQueries(query)
		sum := int32(0)
		for _, pair := range pairs {
			dist := getShortestDistance(pair[0], pair[1])
			sum = sum + pair[0]*pair[1]*dist
		}

		answer := int32(sum % (1000000000 + 7))
		result = append(result, answer)
	}
	return result
}

func getQueries(queryArr []int32) [][]int32 {
	queries := make([][]int32, 0, 1)
	for i := 0; i < len(queryArr); i++ {
		for j := i; j < len(queryArr); j++ {
			pair := []int32{queryArr[i], queryArr[j]}
			queries = append(queries, pair)
		}
	}
	return queries
}

func getShortestDistance(src, dst int32) int32 {
	visited := make(map[int32]bool)
	if src == dst {
		return int32(0)
	}
	nodes := NodeMap[src].Neighbours
	visited[src] = true
	var distance = int32(1)
	for len(nodes) >= 1 {
		toVisit := make([]*Node, 0, 1)
		for i := 0; i < len(nodes); i++ {
			_, nodeAlreadyVisited := visited[nodes[i].Data]
			if !nodeAlreadyVisited {
				if nodes[i].Data == dst {
					return distance
				}
				if len(nodes[i].Neighbours) > 0 {
					toVisit = append(toVisit, nodes[i].Neighbours...)
				}
				visited[nodes[i].Data] = true
			}
		}
		nodes = nil
		nodes = make([]*Node, 0, 1)
		nodes = append(nodes, toVisit...)
		distance = distance + 1
	}
	return -1
}

func CreateGraph(edgeList [][]int32) {
	for _, edge := range edgeList {
		src := edge[0]
		dst := edge[1]

		srcNode, srcNodePresent := NodeMap[src]
		dstNode, dstNodePresent := NodeMap[dst]
		switch {
		case !srcNodePresent && !dstNodePresent:
			srcNode = &Node{
				Data: src,
			}
			dstNode = &Node{
				Data: dst,
			}
			srcNode.Neighbours = []*Node{dstNode}
			dstNode.Neighbours = []*Node{srcNode}
			NodeMap[srcNode.Data] = srcNode
			NodeMap[dstNode.Data] = dstNode

		case srcNodePresent && !dstNodePresent:
			dstNode = &Node{
				Data: dst,
			}
			dstNode.Neighbours = []*Node{srcNode}
			NodeMap[dstNode.Data] = dstNode

			srcNode.Neighbours = append(srcNode.Neighbours, dstNode)

		case !srcNodePresent && dstNodePresent:
			srcNode = &Node{
				Data: src,
			}
			srcNode.Neighbours = []*Node{dstNode}
			NodeMap[srcNode.Data] = srcNode

			dstNode.Neighbours = append(dstNode.Neighbours, srcNode)

		case srcNodePresent && dstNodePresent:
			srcNode.Neighbours = append(srcNode.Neighbours, dstNode)
			dstNode.Neighbours = append(dstNode.Neighbours, srcNode)

		}

	}
}

func main() {
	var n int32 //Number of nodes in a tree
	var q int32 //Number of nodes in a query set

	edgeList := make([][]int32, 0, 1)
	queries := make([][]int32, 0, 1)

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	nqStr := readLine(reader)

	nqStrArr := strings.Split(nqStr, " ")

	n64, err := strconv.ParseInt(nqStrArr[0], 10, 64)
	n = int32(n64)
	checkError(err)

	q64, err := strconv.ParseInt(nqStrArr[1], 10, 64)
	checkError(err)
	q = int32(q64)

	for i := int32(0); i < (n - 1); i++ {
		edgeStr := readLine(reader)
		edgeStrArr := strings.Split(edgeStr, " ")
		src, err := strconv.ParseInt(edgeStrArr[0], 10, 64)
		checkError(err)
		dest, err := strconv.ParseInt(edgeStrArr[1], 10, 64)
		checkError(err)
		edge := []int32{
			int32(src),
			int32(dest),
		}
		edgeList = append(edgeList, edge)
	}

	for i := int32(0); i < q; i++ {
		numStr := readLine(reader)
		num, err := strconv.ParseInt(numStr, 10, 64)
		checkError(err)
		query := make([]int32, 0, 1)
		queryStr := readLine(reader)
		queryStrArr := strings.Split(queryStr, " ")
		for j := 0; j < int(num); j++ {
			node, err := strconv.ParseInt(queryStrArr[j], 10, 64)
			checkError(err)
			query = append(query, int32(node))

		}
		queries = append(queries, query)
	}

	output := kittyCalculation(edgeList, queries)
	for _, answer := range output {
		fmt.Fprintln(writer, answer)
	}
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
