package algo_prog_dynamique

import (
	"fmt"
	"math"
	"time"

	"../common"
)

/* Fonction qui implémente un algorithme de programmation dynamique */
func Knapsack(objects []common.Objects, capacity_max int) (int, []common.Objects) {
	n := len(objects)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, capacity_max+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= capacity_max; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if j < objects[i-1].Weight {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i-1][j-objects[i-1].Weight]+objects[i-1].Value)))
			}
		}
	}

	// Récupérer les objets sélectionnés
	selectedObjects := make([]common.Objects, 0)
	i, j := n, capacity_max
	for i > 0 && j > 0 {
		if dp[i][j] != dp[i-1][j] {
			selectedObjects = append(selectedObjects, objects[i-1])
			j -= objects[i-1].Weight
		}
		i--
	}

	return dp[n][capacity_max], selectedObjects
}

func PrintNewBag(objects []common.Objects) {
	fmt.Println("Les objets qui peuvent être emportés dans le sac :")
	for _, obj := range objects {
		fmt.Printf("L'objet de poids %d et de valeur %d\n", obj.Weight, obj.Value)
	}
}

func SubsetWeight(objects []common.Objects) int {
	weight := 0
	for _, obj := range objects {
		weight += obj.Weight
	}
	return weight
}

func MeasureExecutionTime(fn func(), fnName string) time.Duration {
	startTime := time.Now()
	fn()
	elapsedTime := time.Since(startTime)
	fmt.Printf("Temps d'exécution pour %s : %d µs\n", fnName, elapsedTime.Microseconds())
	return elapsedTime
}
