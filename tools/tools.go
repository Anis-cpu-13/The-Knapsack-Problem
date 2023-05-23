package tools

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"sort"

	"../algo_prog_dynamique"
	"../algorithme_glouton"
	"../common"
	"../merkel_hellman"
	"../reserch_exhastive"
)

// Change algorithme_glouton.Objects to common.Objects
func LoadDataFromFile(filename string) ([]common.Objects, error) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data []common.Objects
	err = json.Unmarshal(fileBytes, &data)
	if err != nil {
		return nil, err
	}

	// Trier les objets par rapport valeur/poids décroissant
	sort.Slice(data, func(i, j int) bool {
		return float64(data[i].Value)/float64(data[i].Weight) > float64(data[j].Value)/float64(data[j].Weight)
	})

	return data, nil
}

func SolveKnapsackWithGreedyAlgorithm(data []common.Objects, capacity int) string {
	knapsackExecutionTime := algorithme_glouton.MeasureExecutionTime(func() {
		selectedObjects, totalWeight, remainingWeight := algorithme_glouton.Knapsack(data, capacity)
		algorithme_glouton.PrintNewBag(selectedObjects)
		fmt.Printf("Le poids total du sac à dos est de %d\n", totalWeight)
		fmt.Printf("Le poids restant dans le sac à dos est de %d\n", remainingWeight)
	}, "knapsack")

	return fmt.Sprintf("Temps d'exécution total pour la résolution du problème du sac à dos : %s\n", knapsackExecutionTime)
}

func SolveKnapsackWithDynamicProgramming(filename string, capacity int) string {
	data2, err := LoadDataFromFile(filename)
	if err != nil {
		log.Fatal("Erreur lors de la lecture du fichier :", err)
	}

	knapsackExecutionTime2 := algo_prog_dynamique.MeasureExecutionTime(func() {
		_, selectedObjects := algo_prog_dynamique.Knapsack(data2, capacity)
		totalWeight := algo_prog_dynamique.SubsetWeight(selectedObjects)
		algo_prog_dynamique.PrintNewBag(selectedObjects)
		fmt.Printf("Le poids total du sac à dos est de %d\n", totalWeight)
	}, "knapsack")

	return fmt.Sprintf("Temps d'exécution total pour la résolution du problème du sac à dos avec l'algotihme programmation dynamique : %s\n", knapsackExecutionTime2)
}

func SolveKnapsackWithExhaustiveSearch(filename string, capacity int) string {
	data3, err := LoadDataFromFile(filename)
	if err != nil {
		log.Fatal("Erreur lors de la lecture du fichier :", err)
	}

	knapsackExecutionTime3 := reserch_exhastive.MeasureExecutionTime(func() {
		_, selectedObjects := reserch_exhastive.Knapsack(data3, capacity)
		totalWeight := reserch_exhastive.SubsetWeight(selectedObjects)
		reserch_exhastive.PrintNewBag(selectedObjects)
		fmt.Printf("Le poids total du sac à dos est de %d\n", totalWeight)
	}, "knapsack")

	return fmt.Sprintf("Temps d'exécution total pour la résolution du problème du sac à dos avec l'algorithme recherche exhaustive : %s\n", knapsackExecutionTime3)
}

func GenerateKeys() {
	privKey, pubKey, err := merkel_hellman.GenerateKeys(10000, 5)
	if err != nil {
		panic(err)
	}

	message := "Hello, world"
	c, err := merkel_hellman.Encrypt(pubKey, message)
	if err != nil {
		panic(err)
	}

	decrypted, err := merkel_hellman.Decrypt(privKey, c)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Message: %s\n", message)
	fmt.Printf("Chiffrement: %s\n", hex.EncodeToString(c.Bytes()))
	fmt.Printf("Déchiffrement: %s\n", decrypted)

	// Décrypter le message
	// plaintext, err := lll_merkel_hellman.CryptanalyseMerkleHellman(c, pubKey, privKey)
	// if err != nil {
	// 	fmt.Println("Erreur lors du décryptage du message:", err)
	// 	return
	// }
	//
	// fmt.Println("Message chiffré:", c)
	// fmt.Println("Message décrypté:", plaintext)
	// fmt.Println()
}

func PerformKnapsackBenchmark(filename string, capacity int) {
	// Charger les données depuis le fichier JSON
	data, err := LoadDataFromFile(filename)
	if err != nil {
		log.Fatal("Erreur lors de la lecture du fichier :", err)
	}

	// Résoudre le problème du sac à dos avec algorithme glouton
	fmt.Println("Résolution du problème du sac à dos avec l'algorithme glouton :")
	resultGreedy := SolveKnapsackWithGreedyAlgorithm(data, capacity)
	fmt.Println(resultGreedy)

	// Afficher la consommation mémoire
	printMemoryUsage()

	// Résoudre le problème du sac à dos avec algorithme programmation dynamique
	fmt.Println("Résolution du problème du sac à dos avec l'algorithme de programmation dynamique :")
	resultDynamic := SolveKnapsackWithDynamicProgramming(filename, capacity)
	fmt.Println(resultDynamic)

	// Afficher la consommation mémoire
	printMemoryUsage()

	// Résoudre le problème du sac à dos avec algorithme de recherche exhaustive
	fmt.Println("Résolution du problème du sac à dos avec l'algorithme de recherche exhaustive :")
	resultExhaustive := SolveKnapsackWithExhaustiveSearch(filename, capacity)
	fmt.Println(resultExhaustive)

	// Afficher la consommation mémoire finale
	printMemoryUsage()
}

func printMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Consommation mémoire : %d bytes\n", m.Sys)
}
