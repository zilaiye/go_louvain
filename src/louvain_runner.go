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

	louvain := louvain.Louvain{}
	louvain.Load(os.Args[1])
	louvain.Compute()

	fmt.Printf("Modularity Q: %f\n", louvain.BestModularity())
	fmt.Printf("Nodes to communities.\n")
	for nodeId, commId := range louvain.GetBestPertition() {
		fmt.Printf("nodeId: %d commId: %d \n", nodeId, commId)
	}

}
