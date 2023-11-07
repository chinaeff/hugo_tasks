package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	ID    int
	Name  string
	Form  string
	Links []*Node
}

func main() {
	rand.Seed(time.Now().UnixNano())

	go func() {
		for {
			numNodes := 5 + rand.Intn(26)
			nodes := generateRandomGraph(numNodes)

			mermaidGraph := generateMermaidGraph(nodes)

			fmt.Println(mermaidGraph)

			time.Sleep(5 * time.Second)
		}
	}()

	select {}
}

func generateRandomGraph(numNodes int) []*Node {
	nodes := make([]*Node, numNodes)

	for i := 0; i < numNodes; i++ {
		node := &Node{
			ID:    i,
			Name:  fmt.Sprintf("Node%d", i),
			Form:  getRandomForm(),
			Links: []*Node{},
		}
		nodes[i] = node
	}

	for i := 0; i < numNodes; i++ {
		node := nodes[i]
		numLinks := rand.Intn(numNodes - 1)
		for j := 0; j < numLinks; j++ {
			linkedNode := nodes[(i+j+1)%numNodes] // Выбираем следующий узел в круговом порядке
			node.Links = append(node.Links, linkedNode)
		}
	}

	return nodes
}

func getRandomForm() string {
	forms := []string{"circle", "rect", "square", "ellipse", "round-rect", "rhombus"}
	return forms[rand.Intn(len(forms))]
}

func generateMermaidGraph(nodes []*Node) string {
	graphText := "graph LR\n"
	for _, node := range nodes {
		graphText += fmt.Sprintf("%s[%s] -->", node.Name, node.Name)
		for _, linkedNode := range node.Links {
			graphText += fmt.Sprintf(" %s[%s],", linkedNode.Name, linkedNode.Name)
		}
		graphText = graphText[:len(graphText)-1]
		graphText += "\n"
	}

	return graphText
}
