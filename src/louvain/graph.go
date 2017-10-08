package louvain

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
	return len(this.selfs)
}
