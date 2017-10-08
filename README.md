# go louvain

This package implements the louvain algorithm in go.

# usage

```bash
./louvain_runner [input file]
```

An example is shown below.
```bash
git clone https://github.com/ken57/go_louvain.git

cd go_louvain
go build src/louvain_runner.go
./louvain_runner src/louvain/resource/karate.txt
```

and it will be obtained.
```
Modularity Q: 0.418803
Nodes to communities.
nodeId: 10 commId: 0
nodeId: 34 commId: 1
nodeId: 14 commId: 0
nodeId: 15 commId: 1
...
```

# A format of input file

The input file is in the following format.
```
[source],[dest]
```
An example is shown below.
```
1,2
1,3
1,4
1,5
1,6
1,7
...
```

The input file is Interpreted as an undirected graph.
