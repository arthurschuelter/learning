package core

import (
	"fmt"
	"math/rand"
)

type GeneticAlgorithm struct {
	population     *Population
	problem        Problem
	populationSize int
	mutationRate   float32
}

type Problem interface {
	GenerateNewChromossome() Chromossome
	CalculateFitness(c Chromossome) float32
}

type Chromossome struct {
	Chromossome     []int
	ChromossomeSize int
	Fitness         float32
	MutationRate    float32
}

type Population struct {
	Population     []Chromossome
	PopulationSize int
	SurvivorSize   int
}

func (p *Population) Print() {
	fmt.Println("Population size:", p.PopulationSize)
	for i, g := range p.Population {
		fmt.Print(i, " ")
		g.Print()
		fmt.Printf(" Fit: %.2f\n", g.Fitness)
	}
}

func (g *Chromossome) Print() {
	fmt.Printf("[")
	for _, v := range g.Chromossome {
		fmt.Printf("%2d, ", v)
	}
	fmt.Printf("]")
}

func (g *Chromossome) Mutate() {
	rate := g.MutationRate

	for i := range g.Chromossome {
		chance := rand.Float32()
		if chance < rate {
			// fmt.Print(i, " Mutei ", g.Chromossome[i], " -> ")
			g.Chromossome[i] = rand.Intn(18)
			// fmt.Println(g.Chromossome[i])
		}
	}
}
