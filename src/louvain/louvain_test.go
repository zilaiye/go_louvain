package louvain

import (
	"math"
	"testing"
)

func TestModularity(t *testing.T) {

	graphReader := NewGraphReader()
	graph := graphReader.Load("resource/karate.txt")

	louvain := NewLouvain(graph)
	louvain.Compute()

	actual := louvain.BestModularity()
	// Python-Louvainでkarate clubをlouvain法で処理したときの最終的なModulality
	expected := WeightType(0.41880345)
	// 順番が変わると、若干最終的なModularityが変わることがあるようなので、0.01以下の差は許容。
	// TODO:もっと良い評価方法がないか検討
	if math.Abs(float64(actual-expected)) > 0.01 {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

}
