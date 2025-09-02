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
	// Below are the tests to run for this demo
	// ####################

	// >>> TEST 1 >>>
	// For the first test, please run a few of the classic graph generators and look at the output.
	// We recommend keeping the graphs fairly small for this test.
	// Are the functions intuitive? Are the outputs interpretable?

	var run_test__classic_graph_generator = true // Set this to true to run this test
	if run_test__classic_graph_generator {
		var classic_graph = model.CycleGraph(5) // Edit this line

		fmt.Println("### Running test 1 ###")
		fmt.Printf("Nodes: {")
		for node := range classic_graph.Nodes {
			fmt.Printf("%d, ", node)
		}
		fmt.Printf("}\nEdges: ")
		fmt.Println(classic_graph.Edges)
		fmt.Println()
	}

	// >>> TEST 2 >>>
	// For the second test, please run a few of the random graph generators and look at the output.
	// We recommend keeping the graphs fairly small for this test.
	// If everything is going well then the generated graphs should be random each run.
	// Are the outputs indeed random? Can you find a way to make some of them deterministic?

	var run_test_random_graph_generator = false // Set this to true to run this test
	if run_test_random_graph_generator {
		var random_graph = model.FastGNPRandomGraph(5, 0.1) // Edit this line

		fmt.Println("### Running test 2 ###")
		fmt.Printf("Nodes: {")
		for node := range random_graph.Nodes {
			fmt.Printf("%d, ", node)
		}
		fmt.Printf("}\nEdges: ")
		fmt.Println(random_graph.Edges)
		fmt.Println()
	}

	// >>> TEST 3 >>>
	// For the third test, please generate some BIG graphs.
	// By implementing these generators in Go, we can generate them at fairly high speeds
	// How long does it take to run each generator? How big did you make the graphs?

	var run_test_generator_speed = false // Set this to true to run this test
	if run_test_generator_speed {
		fmt.Println("### Running test 3 ###")
		start := time.Now()
		model.FastGNPRandomGraph(10000, 0.1) // Edit this line
		elapsed := time.Since(start)
		fmt.Printf("Time taken: %s\n", elapsed)
		fmt.Println()
	}

	// ####################
	// In the comments below we provide all implemented graph sampling methods
	// Below that are the tests for sampling
	// ####################

	// Deletion graph sampling classes:
	// DeletionRandomNodeSampling
	// DeletionRandomNodeNeighbourSampling
	// DeletionInclusiveRandomNodeNeighbourSampling
	// DeletionRandomDegreeNodeSampling
	// DeletionRandomEdgeSampling
	// DeletionRandomNodeEdgeSampling
	// DeletionHybridSampling
	// DeletionRandomWalkSampling
	// DeletionRandomWalkWithJumpSampling
	// DeletionRandomWalkWithRestartSampling

	// Preservation graph sampling classes:
	// PreservationRandomNodeSampling
	// PreservationRandomNodeNeighbourSampling
	// PreservationInclusiveRandomNodeNeighbourSampling
	// PreservationRandomDegreeNodeSampling
	// PreservationRandomEdgeSampling
	// PreservationRandomNodeEdgeSampling
	// PreservationRandomWalkSampling
	// PreservationRandomWalkWithJumpSampling
	// PreservationRandomWalkWithRestartSampling
	// PreservationTopKEdgeSampling

	// >>> TEST 4 >>>
	// As the name suggests, deletion graph samplers work by deleting parts of the input graph, yielding a smaller output graph.
	// Please try out some of the deletion samplers (above) in the experiment below.
	// Are the sampled graphs indeed smaller than the input graphs?

	var run_test_deletion_sampler = false // Set this to true to run this test
	if run_test_deletion_sampler {
		fmt.Println("### Running test 4 ###")
		var graph = model.CompleteGraph(8) // Edit this line to change the original graph that is created
		fmt.Println("Original Graph:")
		fmt.Printf("Nodes: {")
		for node := range graph.Nodes {
			fmt.Printf("%d, ", node)
		}
		fmt.Printf("}\nEdges: ")
		fmt.Println(graph.Edges)
		fmt.Println()

		var howManyToDelete = 3                        // Edit this line to change how many "things" (e.g. node or edges, depending on the sampler) the sampler will delete.
		sampler := &model.DeletionRandomNodeSampling{} // Edit this line to change the sampler. You can just copy-paste a deletion method from above
		sampler.SamplingStage(graph, howManyToDelete)
		fmt.Println("Sampled Graph:")
		fmt.Printf("Nodes: {")
		for node := range graph.Nodes {
			fmt.Printf("%d, ", node)
		}
		fmt.Printf("}\nEdges: ")
		fmt.Println(graph.Edges)
		fmt.Println()
	}

	// >>> TEST 5 >>>
	// Preservation graph samplers work by selecting parts to keep from the input graph, yielding a smaller output graph.
	// Please try out some of the deletion samplers (above) in the experiment below.
	// Are the sampled graphs indeed smaller than the input graphs?

	var run_test_preservation_sampler = false // Set this to true to run this test
	if run_test_preservation_sampler {
		fmt.Println("### Running test 5 ###")
		var graph = model.CompleteGraph(8) // Edit this line to change the original graph that is created
		fmt.Println("Original Graph:")
		fmt.Printf("Nodes: {")
		for node := range graph.Nodes {
			fmt.Printf("%d, ", node)
		}
		fmt.Printf("}\nEdges: ")
		fmt.Println(graph.Edges)
		fmt.Println()

		var sampledGraphSizeRatio = float32(0.3)           // Edit this line to change how much smaller the sample should be, compared to the original graph
		sampler := &model.PreservationRandomNodeSampling{} // Edit this line to change the sampler
		var sampled_graph, _ = sampler.Sample(*graph, sampledGraphSizeRatio)
		fmt.Println("Sampled Graph:")
		fmt.Printf("Nodes: {")
		for node := range sampled_graph.Nodes {
			fmt.Printf("%d, ", node)
		}
		fmt.Printf("}\nEdges: ")
		fmt.Println(sampled_graph.Edges)
		fmt.Println()
	}

	// That was the last test for the Go-Network!
	// Thank you for completing this part of the demo!
}
