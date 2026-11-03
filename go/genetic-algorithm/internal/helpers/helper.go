package helpers

import (
	"fmt"
	"gene-algo/internal/core"
	"sort"
)

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

func SortList(list []core.Chromossome) []core.Chromossome {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Fitness > list[j].Fitness
	})
	return list
}
