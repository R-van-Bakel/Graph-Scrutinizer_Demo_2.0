package main

import (
	"fmt"

	"github.com/jmCodeCraft/go-network/model"

	"time"
)

func main() {
	// Classic generators:
	// CompleteGraph(numberOfNodes int)
	// LadderGraph(nodesInSinglePath int)
	// CircularLadderGraph(nodesInSinglePath int)
	// WheelGraph(numberOfNodes int)
	// TuranGraph(numberOfNodes int, numberOfPartitions int)
	// TrivialGraph()
	// NullGraph()
	// TadpoleGraph(cycleSize int, pathSize int)
	// StarGraph(numberOfNodes int)
	// PathGraph(numberOfNodes int)
	// LollipopGraph(completeGraphSize int, pathGraphSize int)
	// CycleGraph(numberOfNodes int)
	// CirculantGraph(numberOfNodes int, offset int)

	// Random generators:
	// FastGNPRandomGraph(numberOfNodes int, probabilityForEdgeCreation float64)
	// DenseGNMRandomGraph(numberOfNodes int, numberOfEdges int)
	// BarabasiAlbertRandomGraph(numberOfNodes int, numberOfEdges int) (g *UndirectedGraph)
	// WattsStrogatzRandomGraph(numberOfNodes int, nearestNeighboursCount int, edgeRewiringProbability float32)

	// ####################
	// In the comments above we provide all implemented graph generators
	// Below are the tests to trun for this demo
	// ####################

	// For the first test, please run a few of the classic graph generators and look at the output.
	// We recommend keeping the graphs fairly small for this test.
	// Are the functions intuitive? Are the outputs interpretable?
	var run_test__classic_graph_generator = true // Set this to true to run this test
	if run_test__classic_graph_generator {
		var classic_graph = model.CycleGraph(5) // Edit this line

		fmt.Printf("Nodes: {")
		for node := range classic_graph.Nodes {
			fmt.Printf("%d, ", node)
		}
		fmt.Printf("}\nEdges: ")
		fmt.Println(classic_graph.Edges)
	}

	// For the second test, please run a few of the random graph generators and look at the output.
	// We recommend keeping the graphs fairly small for this test.
	// If everything is going well then the generated graphs should be random each run.
	// Are the outputs indeed random? Can you find a way to make some of them deterministic?
	var run_test_random_graph_generator = false // Set this to true to run this test
	if run_test_random_graph_generator {
		var random_graph = model.FastGNPRandomGraph(5, 0.1) // Edit this line

		fmt.Printf("Nodes: {")
		for node := range random_graph.Nodes {
			fmt.Printf("%d, ", node)
		}
		fmt.Printf("}\nEdges: ")
		fmt.Println(random_graph.Edges)
	}

	// For the third test, please generate some BIG graphs.
	// By implementing these generators in Go, we can generate them at fairly high speeds
	// How long does it take to run each generator? How big did you make the graphs?
	var run_test_generator_speed = false // Set this to true to run this test
	if run_test_generator_speed {
		start := time.Now()
		model.FastGNPRandomGraph(10000, 0.1) // Edit this line
		elapsed := time.Since(start)
		fmt.Printf("Time taken: %s\n", elapsed)
	}

	// TODO describe the graph sampler test
	var run_test_sampler = false // Set this to true to run this test
	if run_test_sampler {
		// TODO implement this test
	}
}
