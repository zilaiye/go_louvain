package louvain

import (
	"testing"
)

func TestModularity(t *testing.T) {

	louvain := Louvain{}
	louvain.Load("resource/karate.txt")
	louvain.Compute()

	actual := louvain.BestModularity()
	expected := WeightType(0.41880345)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

}
