package main

import (
	"fmt"
	"gene-algo/internal/core"
	"gene-algo/internal/helpers"
	"gene-algo/pkg/pokemon"
	"math/rand"
	"time"
)

const (
	NORMAL = iota
	FIRE
	WATER
	ELECTRIC
	GRASS
	ICE
	FIGHTING
	POISON
	GROUND
	FLYING
	PSYCHIC
	BUG
	ROCK
	GHOST
	DRAGON
	DARK
	STEEL
	FAIRY
)

func main() {
	fmt.Println("Genetic Algorithm")
	var geneSize int = 6
	var populationSize int = 500
	var generations int = 1000
	var survivorSize int = 20
	var mutationRate float32 = 0.08

	best_solution := core.Chromossome{}

	start := time.Now()

	typeChart := pokemon.MakeTypeChart()
	p := MakePopulation(populationSize, geneSize, survivorSize, mutationRate)

	for j := range generations {
		// fmt.Printf(" === Generation %d === \n", j)

		for i := range p.Population {
			gene := &p.Population[i]
			gene.Fitness = pokemon.CalculateFitness(*gene, typeChart, false)
		}

		p.Population = helpers.SortList(p.Population)
		UpdateBestSolution(p.Population[0], &best_solution)

		if j+1 == generations {
			break
		}

		p.Population = MakeNewGeneration(p, typeChart)
	}

	elapsed := time.Since(start)
	p.Population = helpers.SortList(p.Population)
	p.Print()
	ShowFinalSolution(best_solution, typeChart)
	pokemon.CalculateFitness(best_solution, typeChart, true)
	fmt.Println("Execution time:", elapsed)
}

func UpdateBestSolution(current core.Chromossome, best *core.Chromossome) {
	if current.Fitness > best.Fitness {
		*best = Clone(current)
	}
}

func ShowFinalSolution(g core.Chromossome, t pokemon.TypeChart) {
	fmt.Println("Final Solution:")
	for _, value := range g.Chromossome {
		typeName, err := helpers.GetType(value)
		if err != nil {
			panic(err)
		}
		fmt.Print(typeName + " ")
	}
	fmt.Println()
	fmt.Printf("Fitness: %.2f\n", g.Fitness)
}

func Fitness(g core.Chromossome, t pokemon.TypeChart) float32 {
	var sum float32
	atk := make([]float32, 18)
	def := make([]float32, 18)

	unique := VerifyUnique(g.Chromossome)

	var atkSum float32
	var defSum float32

	for i, v := range g.Chromossome {
		atk[i] = t.Evaluate(v, i)

		switch atk[i] {
		case 0.5:
			atk[i] = -1
		case 0:
			atk[i] = -2
		}

		atkSum += atk[i]

		def[i] = t.Evaluate(i, v)

		switch def[i] {
		case 0.5:
			def[i] = +1
		case 0:
			def[i] = +2
		}

		defSum += def[i]
		sum += float32(v)
	}

	result := (atkSum + defSum) * unique

	return result
}

func MakeNewGeneration(p core.Population, t pokemon.TypeChart) []core.Chromossome {
	survivors := Selection(p)
	newPop := make([]core.Chromossome, 0)

	newSurvivorNum := p.PopulationSize - len(survivors)

	for range newSurvivorNum {
		n1 := rand.Intn(len(survivors))
		n2 := rand.Intn(len(survivors))

		c1 := Crossover(survivors[n1], survivors[n2])
		c1.Fitness = pokemon.CalculateFitness(c1, t, false)
		newPop = append(newPop, c1)
	}

	newPop = append(newPop, survivors...)

	for i := range newPop {
		// newPop[i].Print(i)
		// fmt.Println()
		newPop[i].Mutate()
		// newPop[i].Print(i)
		// fmt.Println()
	}

	return newPop
}

func Selection(p core.Population) []core.Chromossome {
	// p.population = SortList(p.population)
	return p.Population[0:p.SurvivorSize]
}

func Crossover(p1 core.Chromossome, p2 core.Chromossome) core.Chromossome {
	c1 := Clone(p1)
	c2 := Clone(p2)

	idx := rand.Intn(6)

	for i := range p1.Chromossome {
		if i >= idx {
			c1.Chromossome[i] = p2.Chromossome[i]
			c2.Chromossome[i] = p1.Chromossome[i]
		}
	}
	return c1
}

func VerifyUnique(gene []int) float32 {
	collision := make(map[int]bool)
	var multiplier float32 = 1.0
	for _, v := range gene {
		_, ok := collision[v]
		if !ok {
			collision[v] = true
		} else {
			multiplier *= 0.5
		}
	}

	if multiplier == 1.0 {
		return 1.5
	}

	return multiplier
}

func MakeGene(size int, mutationRate float32) core.Chromossome {
	gene := make([]int, size)

	for i := range gene {
		gene[i] = rand.Intn(18)
	}

	return core.Chromossome{
		Chromossome:     gene,
		ChromossomeSize: size,
		Fitness:         0.00,
		MutationRate:    mutationRate,
	}
}

func Clone(g1 core.Chromossome) core.Chromossome {
	newGene := make([]int, g1.ChromossomeSize)
	copy(newGene, g1.Chromossome)

	return core.Chromossome{
		Chromossome:     newGene,
		ChromossomeSize: g1.ChromossomeSize,
		Fitness:         g1.Fitness,
		MutationRate:    g1.MutationRate,
	}
}

func MakePopulation(populationSize int, geneSize int, survivorSize int, mutationRate float32) core.Population {
	population := make([]core.Chromossome, populationSize)
	for i := range population {
		population[i] = MakeGene(geneSize, mutationRate)
	}
	return core.Population{
		Population:     population,
		PopulationSize: populationSize,
		SurvivorSize:   survivorSize,
	}
}
