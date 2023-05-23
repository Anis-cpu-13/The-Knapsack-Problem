package create_data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type Objects struct {
	Weight int `json:"weight"`
	Value  int `json:"value"`
}

func GenerateData(numExamples int) {
	// Génère les exemples aléatoires
	examples := GenerateRandomExamples(numExamples)

	// Encode les exemples en JSON avec une indentation pour une meilleure lisibilité
	jsonData, err := json.MarshalIndent(examples, "", "\t")
	if err != nil {
		fmt.Println("Erreur lors de l'encodage JSON :", err)
		return
	}

	// Écrit les données encodées dans un fichier
	err = ioutil.WriteFile("data.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier :", err)
		return
	}

	fmt.Println("Données stockées avec succès dans le fichier data.json")
}

func GenerateRandomExamples(numExamples int) []Objects {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	var examples []Objects
	for i := 0; i < numExamples; i++ {
		weight := random.Intn(100)
		value := random.Intn(100)

		example := Objects{
			Weight: weight,
			Value:  value,
		}

		examples = append(examples, example)
	}

	return examples
}
