package main

import (
	"fmt"
	"os"

	"./louvain"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Print("usage: louvain_runner [graph_filename]")
		return
	}

	graphReader := louvain.NewGraphReader()
	graph := graphReader.Load(os.Args[1])
	louvain := louvain.NewLouvain(graph)
	louvain.Compute()

	fmt.Printf("Modularity Q: %f\n", louvain.BestModularity())
	fmt.Printf("Nodes to communities.\n")
	for nodeId, commId := range louvain.GetBestPertition() {
		fmt.Printf("nodeId: %s commId: %d \n", graphReader.GetNodeLabel(nodeId), commId)
	}
}
