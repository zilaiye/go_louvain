package louvain

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type WeightType float32

type Edge struct {
	destId int
	weight WeightType
}

type Edges [][]Edge

type Graph struct {
	incidences Edges
	selfs      []WeightType
}

func (this *Graph) AddUndirectedEdge(sourceId int, destId int, weight WeightType) {
	this.incidences[sourceId] = append(this.incidences[sourceId], Edge{destId, weight})
	this.incidences[destId] = append(this.incidences[destId], Edge{sourceId, weight})
}

func (this *Graph) AddDirectedEdge(sourceId int, destId int, weight WeightType) {
	this.incidences[sourceId] = append(this.incidences[sourceId], Edge{destId, weight})
}

func (this *Graph) AddSelfEdge(nodeId int, weight WeightType) {
	this.selfs[nodeId] = weight
}

func (this *Graph) GetIncidentEdges(nodeId int) []Edge {
	return this.incidences[nodeId]
}

func (this *Graph) GetSelfWeight(nodeId int) WeightType {
	return this.selfs[nodeId]
}

func (this *Graph) GetNodeSize() int {
	return len(this.incidences)
}

func (this *Graph) Load(filename string) {
	this.incidences = make(Edges, 35)
	this.selfs = make([]WeightType, 35)
	fp, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	totalWeight := WeightType(0.0)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		source, _ := strconv.Atoi(line[0])
		dest, _ := strconv.Atoi(line[1])
		weight := WeightType(1.0)
		this.AddUndirectedEdge(source, dest, weight)
		totalWeight += weight
	}

}
