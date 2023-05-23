package reserch_exhastive

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"../common"
)

/* Knapsack résout le problème du sac à dos en utilisant une recherche exhaustive et retourne la meilleure valeur et les objets qui peuvent être emportés dans le sac.*/

func Knapsack(objects []common.Objects, capacity int) (int, []common.Objects) {
	bestValue := 0
	bestSubset := make([]common.Objects, 0)

	// Générer tous les sous-ensembles possibles et trouver celui avec la meilleure valeur
	GenerateSubsets(objects, capacity, 0, make([]common.Objects, 0), &bestValue, &bestSubset)

	return bestValue, bestSubset
}

/* GenerateSubsets génère tous les sous-ensembles possibles d'objets et met à jour la meilleure valeur et le meilleur sous-ensemble. */

func GenerateSubsets(objects []common.Objects, capacity, index int, subset []common.Objects, bestValue *int, bestSubset *[]common.Objects) {
	if index == len(objects) {
		// Calculer la valeur du sous-ensemble généré
		subsetValue := ComputeSubsetValue(subset)
		if subsetValue > *bestValue {
			// Mettre à jour la meilleure valeur et le meilleur sous-ensemble
			*bestValue = subsetValue
			*bestSubset = make([]common.Objects, len(subset))
			copy(*bestSubset, subset)
		}
		return
	}

	if SubsetWeight(subset)+objects[index].Weight <= capacity {
		subset = append(subset, objects[index])
		GenerateSubsets(objects, capacity, index+1, subset, bestValue, bestSubset)
		subset = subset[:len(subset)-1]
	}

	GenerateSubsets(objects, capacity, index+1, subset, bestValue, bestSubset)
}

/* SubsetWeight calcule le poids total d'un sous-ensemble d'objets. */

func SubsetWeight(subset []common.Objects) int {
	weight := 0
	for _, obj := range subset {
		weight += obj.Weight
	}
	return weight
}

func ComputeSubsetValue(subset []common.Objects) int {
	value := 0
	for _, obj := range subset {
		value += obj.Value
	}
	return value
}

func LoadJSONData(filename string) ([]common.Objects, error) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var objects []common.Objects
	err = json.Unmarshal(fileBytes, &objects)
	if err != nil {
		return nil, err
	}

	return objects, nil
}

func PrintNewBag(objects []common.Objects) {
	fmt.Println("Les éléments qui peuvent être emportés dans le sac :")
	for _, obj := range objects {
		fmt.Printf("L'élément de poids %d et de valeur %d\n", obj.Weight, obj.Value)
	}
}

func MeasureExecutionTime(fn func(), fnName string) time.Duration {
	startTime := time.Now()
	fn()
	elapsedTime := time.Since(startTime)
	fmt.Printf("Temps d'exécution pour %s : %d µs\n", fnName, elapsedTime.Microseconds())
	return elapsedTime
}
