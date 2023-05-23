package main

import (
	"fmt"
	"math/big"

	"./algo_reduc_reseau"
	"./create_data"
	"./tools"
)

func main() {
	fmt.Println("=========== Début de l'exécution de The-Knapsack-Problem ===========")
	fmt.Println()

	create_data.GenerateData(100)

	filename := "data.json"
	capacity := 80

	fmt.Println("Exécution du benchmark du problème du sac à dos...")
	tools.PerformKnapsackBenchmark(filename, capacity)
	fmt.Println("Benchmark terminé.")
	fmt.Println()

	// Générer une paire de clés publiques et privées aléatoires
	fmt.Println("Génération des clés...")
	tools.GenerateKeys()
	fmt.Println("Génération des clés terminée.")
	fmt.Println()

	n := 10 // Vous pouvez changer la taille du réseau ici

	// Générer un réseau de Lagarias-Odlyzko
	fmt.Println("Génération du réseau de Lagarias-Odlyzko initial...")
	LONetwork := algo_reduc_reseau.GenerateLagariasOdlyzkoNetwork(n)
	algo_reduc_reseau.PrintMatrix(LONetwork)
	fmt.Println("Génération terminée.")
	fmt.Println()

	// Appliquer l'algorithme LLL au réseau de Lagarias-Odlyzko
	fmt.Println("Réduction du réseau de Lagarias-Odlyzko avec LLL...")
	LOReduced := algo_reduc_reseau.LLL(LONetwork, big.NewRat(3, 4), 1000)
	algo_reduc_reseau.PrintMatrix(LOReduced)
	fmt.Println("Réduction terminée.")
	fmt.Println()

	// Générer un réseau de Joux-Stern
	fmt.Println("Génération du réseau de Joux-Stern initial...")
	JSNetwork := algo_reduc_reseau.GenerateJouxSternNetwork(n)
	algo_reduc_reseau.PrintMatrix(JSNetwork)
	fmt.Println("Génération terminée.")
	fmt.Println()

	// Appliquer l'algorithme LLL au réseau de Joux-Stern
	fmt.Println("Réduction du réseau de Joux-Stern avec LLL...")
	JSReduced := algo_reduc_reseau.LLL(JSNetwork, big.NewRat(3, 4), 1000)
	algo_reduc_reseau.PrintMatrix(JSReduced)
	fmt.Println("Réduction terminée.")
	fmt.Println()

	// Vérifier si les résultats sont corrects
	fmt.Println("Vérification des résultats...")
	if algo_reduc_reseau.AreResultsCorrect(LONetwork, LOReduced) {
		fmt.Println("Les résultats sont corrects.")
	} else {
		fmt.Println("Les résultats sont incorrects.")
	}
	fmt.Println()
	fmt.Println("=========== Fin de l'exécution de The-Knapsack-Problem ===========")
}
