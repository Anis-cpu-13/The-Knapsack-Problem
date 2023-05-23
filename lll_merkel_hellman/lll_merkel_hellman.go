package lll_merkel_hellman

/*en cours */
/* Attention: Activation de l'attaque contre Merkle-Hellman */

import (
	"math/big"

	"../algo_reduc_reseau"
	"../merkel_hellman"
)

func CryptanalyseMerkleHellman(ciphertext *big.Int, pubKey *merkel_hellman.PublicKey, privKey *merkel_hellman.PrivateKey) (plaintext string, err error) {
	//  Construire le réseau à partir de la clé publique
	matrix := algo_reduc_reseau.CreateMatrix(len(pubKey.M), len(pubKey.M))

	for i := 0; i < len(matrix); i++ { // Utiliser len(matrix) au lieu de len(pubKey.M)
		matrix[i][i] = pubKey.M[i]
	}

	//  Appliquer la réduction LLL à ce réseau
	reducedMatrix := algo_reduc_reseau.LLL(matrix, big.NewRat(3, 4), 1000)

	//  La clé privée peut être approximée à partir de la matrice réduite
	approxPrivateKey := make([]*big.Int, len(reducedMatrix))
	for i, vector := range reducedMatrix {
		approxPrivateKey[i] = vector[i]
	}

	//  Utiliser la clé privée approximative pour déchiffrer le texte chiffré
	approxPrivKey := &merkel_hellman.PrivateKey{
		R: approxPrivateKey,
		A: privKey.A,
		B: privKey.B,
	}
	plaintext, err = merkel_hellman.Decrypt(approxPrivKey, ciphertext)

	return plaintext, err
}
