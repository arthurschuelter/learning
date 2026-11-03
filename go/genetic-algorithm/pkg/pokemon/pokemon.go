package pokemon

import (
	"fmt"
	"gene-algo/internal/core"
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

type TypeChart struct {
	Table [18][18]float32
}

func (t *TypeChart) Evaluate(t1 int, t2 int) float32 {
	return t.Table[t1][t2]
}

func GetType(Type int) (string, error) {
	switch Type {
	case 0:
		return "NORMAL", nil
	case 1:
		return "FIRE", nil
	case 2:
		return "WATER", nil
	case 3:
		return "ELECTRIC", nil
	case 4:
		return "GRASS", nil
	case 5:
		return "ICE", nil
	case 6:
		return "FIGHTING", nil
	case 7:
		return "POISON", nil
	case 8:
		return "GROUND", nil
	case 9:
		return "FLYING", nil
	case 10:
		return "PSYCHIC", nil
	case 11:
		return "BUG", nil
	case 12:
		return "ROCK", nil
	case 13:
		return "GHOST", nil
	case 14:
		return "DRAGON", nil
	case 15:
		return "DARK", nil
	case 16:
		return "STEEL", nil
	case 17:
		return "FAIRY", nil
	default:
		return "", fmt.Errorf("Invalid type")
	}
}

func CalculateFitness(g core.Chromossome, t TypeChart, print bool) float32 {
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

func MakeTypeChart() TypeChart {
	matchupChart := [18][18]float32{
		NORMAL:   {1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 0.5, 0.0, 1.0, 1.0, 0.5, 1.0},
		FIRE:     {1.0, 0.5, 0.5, 1.0, 2.0, 2.0, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 0.5, 1.0, 0.5, 1.0, 2.0, 1.0},
		WATER:    {1.0, 2.0, 0.5, 1.0, 0.5, 1.0, 1.0, 1.0, 2.0, 1.0, 1.0, 1.0, 2.0, 1.0, 0.5, 1.0, 1.0, 1.0},
		ELECTRIC: {1.0, 1.0, 2.0, 0.5, 0.5, 1.0, 1.0, 1.0, 0.0, 2.0, 1.0, 1.0, 1.0, 1.0, 0.5, 1.0, 1.0, 1.0},
		GRASS:    {1.0, 0.5, 2.0, 1.0, 0.5, 1.0, 1.0, 0.5, 2.0, 0.5, 1.0, 0.5, 2.0, 1.0, 0.5, 1.0, 0.5, 1.0},
		ICE:      {1.0, 0.5, 0.5, 1.0, 2.0, 0.5, 1.0, 1.0, 2.0, 2.0, 1.0, 1.0, 1.0, 1.0, 2.0, 1.0, 0.5, 1.0},
		FIGHTING: {2.0, 1.0, 1.0, 1.0, 1.0, 2.0, 1.0, 0.5, 1.0, 0.5, 0.5, 0.5, 2.0, 0.0, 1.0, 2.0, 2.0, 0.5},
		POISON:   {1.0, 1.0, 1.0, 1.0, 2.0, 1.0, 1.0, 0.5, 0.5, 1.0, 1.0, 1.0, 0.5, 0.5, 1.0, 1.0, 0.0, 2.0},
		GROUND:   {1.0, 2.0, 1.0, 2.0, 0.5, 1.0, 1.0, 2.0, 1.0, 0.0, 1.0, 0.5, 2.0, 1.0, 1.0, 1.0, 2.0, 1.0},
		FLYING:   {1.0, 1.0, 1.0, 0.5, 2.0, 1.0, 2.0, 1.0, 1.0, 1.0, 1.0, 2.0, 0.5, 1.0, 1.0, 1.0, 0.5, 1.0},
		PSYCHIC:  {1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 2.0, 1.0, 1.0, 0.5, 1.0, 1.0, 1.0, 1.0, 0.0, 0.5, 1.0},
		BUG:      {1.0, 0.5, 1.0, 1.0, 2.0, 1.0, 0.5, 0.5, 1.0, 0.5, 2.0, 1.0, 1.0, 0.5, 1.0, 2.0, 0.5, 0.5},
		ROCK:     {1.0, 2.0, 1.0, 1.0, 1.0, 2.0, 0.5, 1.0, 0.5, 2.0, 1.0, 2.0, 1.0, 1.0, 1.0, 1.0, 0.5, 1.0},
		GHOST:    {0.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 1.0, 1.0, 2.0, 1.0, 0.5, 1.0, 1.0},
		DRAGON:   {1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 1.0, 0.5, 0.0},
		DARK:     {1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 0.5, 1.0, 1.0, 1.0, 2.0, 1.0, 1.0, 2.0, 1.0, 0.5, 1.0, 0.5},
		STEEL:    {1.0, 0.5, 0.5, 0.5, 1.0, 2.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 1.0, 1.0, 1.0, 0.5, 2.0},
		FAIRY:    {1.0, 0.5, 1.0, 1.0, 1.0, 1.0, 2.0, 0.5, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 2.0, 0.5, 1.0},
	}
	return TypeChart{
		Table: matchupChart,
	}
}
