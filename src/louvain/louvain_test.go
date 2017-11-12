package louvain

import (
	"fmt"
	"reflect"
	"testing"
)

func TestModularity(t *testing.T) {

	graphReader := NewGraphReader()
	graph := graphReader.Load("resource/karate.txt")

	louvain := NewLouvain(graph)
	louvain.Compute()

	actual := louvain.BestModularity()
	expected := WeightType(0.41880345)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

}

func TestNeighber(t *testing.T) {

	graphReader := NewGraphReader()
	graph := graphReader.Load("resource/dynamic_test.txt")
	louvain := NewLouvain(graph)
	louvain.Compute()

	for nodeId := 0; nodeId != graph.GetNodeSize(); nodeId++ {
		fmt.Print(graphReader.GetNodeLabel(nodeId) + " ")
		fmt.Println(louvain.GetNodeToCommunityInEachLevel(nodeId))
	}
	changeNodes := graphReader.GetNodeIndices([]string{"6", "8"})

	actual0 := graphReader.GetNodeLabels(louvain.GetNeighbers(changeNodes, 0))
	expected0 := []string{"6", "8"}
	if !reflect.DeepEqual(actual0, expected0) {
		t.Errorf("got %v\nwant %v", actual0, expected0)
	}

	actual1 := graphReader.GetNodeLabels(louvain.GetNeighbers(changeNodes, 1))
	expected1 := []string{"0", "6", "4", "9", "5", "7", "8"}
	if !reflect.DeepEqual(actual1, expected1) {
		t.Errorf("got %v\nwant %v", actual1, expected1)
	}

	actual2 := graphReader.GetNodeLabels(louvain.GetNeighbers(changeNodes, 2))
	expected2 := []string{"0", "1", "2", "3", "6", "4", "9", "5", "7", "8"}
	if !reflect.DeepEqual(actual2, expected2) {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}
	louvain.level[0].graph.AddUndirectedEdge(changeNodes[0], changeNodes[1], 1.0)
	louvain.LocalChange(changeNodes, 0)

	for nodeId := 0; nodeId != graph.GetNodeSize(); nodeId++ {
		fmt.Print(graphReader.GetNodeLabel(nodeId) + " ")
		fmt.Println(louvain.GetNodeToCommunityInEachLevel(nodeId))
	}
}
