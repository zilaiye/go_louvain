package main

import (
	"flag"
	"fmt"

	"./louvain"
)

func main() {
	var inputFilename = flag.String("i", "", "Input graph filename for loouvain. [Required]")
	var showCommunityIdOfEachLayer = flag.Bool("l", false, "Show community id of each layer as result. Default setting is false.")
	flag.Parse()

	if *inputFilename == "" {
		fmt.Print("Input filename [-i] is required.")
		return
	}

	graphReader := louvain.NewGraphReader()
	graph := graphReader.Load(*inputFilename)
	louvain := louvain.NewLouvain(graph)
	louvain.Compute()

	fmt.Printf("Best Modularity Q: %f\n", louvain.BestModularity())

	if *showCommunityIdOfEachLayer == false {
		fmt.Printf("Nodes to communities.\n")
		for nodeId, commId := range louvain.GetBestPertition() {
			fmt.Printf("nodeId: %s communityId: %d \n", graphReader.GetNodeLabel(nodeId), commId)
		}
	} else {
		fmt.Println("[NodeId] [CommunityId in each layer]")
		for nodeId := 0; nodeId != graph.GetNodeSize(); nodeId++ {
			fmt.Print(graphReader.GetNodeLabel(nodeId) + " ")
			fmt.Println(louvain.GetNodeToCommunityInEachLevel(nodeId))
		}
	}
}
