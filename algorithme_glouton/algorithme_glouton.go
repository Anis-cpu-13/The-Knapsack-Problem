package algorithme_glouton

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"../common"
)

func LoadJSONData(filename string) ([]common.Objects, error) {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data []common.Objects
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Knapsack(data []common.Objects, capacity int) ([]common.Objects, int, int) {
	var weight int
	var results []common.Objects
	for _, obj := range data {
		if weight+obj.Weight <= capacity {
			weight += obj.Weight
			results = append(results, obj)
		} else {
			// Si ajouter l'objet dépasse la capacité, nous arrêtons la boucle
			break
		}
	}
	return results, weight, capacity - weight
}

func PrintNewBag(data []common.Objects) {
	fmt.Println("Les éléments qui peuvent être emportés dans le sac :")
	for _, obj := range data {
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
