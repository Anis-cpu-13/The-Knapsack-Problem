package algo_reduc_reseau

import (
	"fmt"
	"math/big"
)

type Vector []*big.Int
type Matrix []Vector

/* Fonction pour crée un vecteur de taille n */
func CreateVector(n int) Vector {
	v := make(Vector, n)

	for i := 0; i < n; i++ {
		v[i] = big.NewInt(0)
	}

	return v
}

/* Fonction pour crée une nouvelle matrice de taille n*m */
func CreateMatrix(m int, n int) Matrix {
	M := make(Matrix, m)

	for i := 0; i < m; i++ {
		M[i] = CreateVector(n)
	}

	return M
}

/* Fonction qui soustrait deux vecteurs */
func VectorSub(a, b Vector) Vector {
	if len(a) != len(b) {
		panic("Vectors must be the same size to be subtracted")
	}

	result := CreateVector(len(a))
	for i := range a {
		result[i].Sub(a[i], b[i])
	}

	return result
}

/* Fonction qui calcule le produit scalaire deux vecteur */
func DotProduct(a, b Vector) *big.Int {
	if len(a) != len(b) {
		panic("The vectors must have the same size for the dot product")
	}

	result := big.NewInt(0)
	tmp := big.NewInt(0)

	for i := range a {
		tmp.Mul(a[i], b[i])
		result.Add(result, tmp)
	}
	return result
}

/* Fonction qui calcule le produit d'un vecteur par un scalaire */
func MulVecToScal(v Vector, scalar *big.Int) Vector {

	result := CreateVector(len(v))

	for i := range v {
		result[i].Mul(v[i], scalar)
	}

	return result
}

/* Cette fonction prend un vecteur v et un nombre rationnel r en tant que paramètres d'entrée. */
func MulVecToRat(v Vector, r *big.Rat) Vector {
	result := make(Vector, len(v))

	for i, value := range v {
		temp := new(big.Rat).Mul(r, new(big.Rat).SetInt(value))
		result[i] = temp.Num()
	}

	return result
}

/* Fonction pour effectuer l'orthogonalisation de Gram-Schmidt */
func GramSchmidtOrthogonalization(B Matrix) Matrix {
	m := len(B)
	U := make(Matrix, m)

	for i := 0; i < m; i++ {
		U[i] = make(Vector, len(B[i]))
		copy(U[i], B[i])

		for j := 0; j < i; j++ {
			proj := new(big.Rat).SetInt(DotProduct(U[j], B[i]))
			proj.Quo(proj, new(big.Rat).SetInt(DotProduct(U[j], U[j])))
			U[i] = VectorSub(U[i], MulVecToRat(U[j], proj))
		}
	}

	return U
}

func LLL(B Matrix, delta *big.Rat, MaxIterations int) Matrix {
	k := 1
	m := len(B)
	iter := 0

	U := GramSchmidtOrthogonalization(B)

	for k < m && iter < MaxIterations {
		iter++

		for j := k - 1; j >= 0; j-- {
			q := new(big.Rat).SetInt(DotProduct(U[j], B[k]))
			q.Quo(q, new(big.Rat).SetInt(DotProduct(U[j], U[j])))
			roundQ := new(big.Int).Set(q.Num())
			roundQ.Div(roundQ, q.Denom())
			B[k] = VectorSub(B[k], MulVecToScal(B[j], roundQ))
			U = GramSchmidtOrthogonalization(B)
		}

		d := new(big.Rat).Mul(delta, new(big.Rat).SetInt(DotProduct(U[k-1], U[k-1])))

		v := new(big.Rat).SetInt(DotProduct(U[k], U[k]))
		v.Add(v, new(big.Rat).SetInt(DotProduct(U[k-1], U[k])))

		if d.Cmp(v) < 0 {
			B[k], B[k-1] = B[k-1], B[k]
			k = Max(k-1, 1)
			U = GramSchmidtOrthogonalization(B)
		} else {
			k++
		}
	}

	return B
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/* Fonction pour générer des réseaux Lagarias-Odlyzko */
func GenerateLagariasOdlyzkoNetwork(n int) Matrix {
	M := CreateMatrix(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				M[i][j].SetInt64(1)
			} else {
				M[i][j].SetInt64(int64(i * j % (n + 1)))
			}
		}
	}
	return M
}

/* Fonction pour générer des réseaux Joux-Stern */
func GenerateJouxSternNetwork(n int) Matrix {
	M := CreateMatrix(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				M[i][j].SetInt64(1)
			} else {
				M[i][j].SetInt64(int64(i * j % (n * (n - 1) / 2)))
			}
		}
	}
	return M
}

// Fonction pour imprimer une matrice
func PrintMatrix(M Matrix) {
	for _, row := range M {
		for _, elem := range row {
			fmt.Printf("%4s ", elem.String())
		}
		fmt.Println()
	}
}

func DotProductVec(v1, v2 []*big.Int) *big.Int {
	sum := big.NewInt(0)
	for i := 0; i < len(v1); i++ {
		sum.Add(sum, new(big.Int).Mul(v1[i], v2[i]))
	}
	return sum
}

func VectorLength(v []*big.Int) *big.Float {
	sum := big.NewInt(0)
	for _, val := range v {
		sum.Add(sum, new(big.Int).Mul(val, val))
	}
	return new(big.Float).Sqrt(new(big.Float).SetInt(sum))
}

func AreResultsCorrect(initial, reduced Matrix) bool {
	n := len(initial)

	// Vérifier que les déterminants sont égaux
	detInitial := big.NewInt(1)
	detReduced := big.NewInt(1)
	for i := 0; i < n; i++ {
		detInitial.Mul(detInitial, initial[i][i])
		detReduced.Mul(detReduced, reduced[i][i])
	}
	if detInitial.Cmp(detReduced) != 0 {
		return false
	}

	// Vérifier que les vecteurs de base réduits sont plus courts
	for i := 0; i < n; i++ {
		if VectorLength(reduced[i]).Cmp(VectorLength(initial[i])) > 0 {
			return false
		}
	}

	// Vérifier que les vecteurs de base réduits sont plus orthogonaux
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dotInitial := new(big.Float).SetInt(DotProductVec(initial[i], initial[j]))
			dotReduced := new(big.Float).SetInt(DotProductVec(reduced[i], reduced[j]))
			lenInitialI := VectorLength(initial[i])
			lenInitialJ := VectorLength(initial[j])
			lenReducedI := VectorLength(reduced[i])
			lenReducedJ := VectorLength(reduced[j])

			cosThetaInitial := new(big.Float).Quo(dotInitial, new(big.Float).Mul(lenInitialI, lenInitialJ))
			cosThetaReduced := new(big.Float).Quo(dotReduced, new(big.Float).Mul(lenReducedI, lenReducedJ))

			absCosThetaInitial, _ := new(big.Float).Abs(cosThetaInitial).Float64()
			absCosThetaReduced, _ := new(big.Float).Abs(cosThetaReduced).Float64()

			if absCosThetaReduced > absCosThetaInitial {
				return false
			}
		}
	}

	return true
}

func EfficiencyScore(reduced Matrix) float64 {
	n := len(reduced)
	sumVectorLength := big.NewFloat(0)
	sumOrthogonality := big.NewFloat(0)

	for i := 0; i < n; i++ {
		sumVectorLength.Add(sumVectorLength, VectorLength(reduced[i]))

		for j := i + 1; j < n; j++ {
			dotReduced := new(big.Float).SetInt(DotProductVec(reduced[i], reduced[j]))
			lenReducedI := VectorLength(reduced[i])
			lenReducedJ := VectorLength(reduced[j])

			cosThetaReduced := new(big.Float).Quo(dotReduced, new(big.Float).Mul(lenReducedI, lenReducedJ))

			sumOrthogonality.Add(sumOrthogonality, new(big.Float).Abs(cosThetaReduced))
		}
	}

	efficiency, _ := new(big.Float).Quo(sumVectorLength, sumOrthogonality).Float64()
	return efficiency
}

func CompareNetworkEfficiency(reducedLO, reducedJS Matrix) {
	efficiencyLO := EfficiencyScore(reducedLO)
	efficiencyJS := EfficiencyScore(reducedJS)

	fmt.Printf("Efficacité du réseau Lagarias-Odlyzko : %f\n", efficiencyLO)
	fmt.Printf("Efficacité du réseau Joux-Stern : %f\n", efficiencyJS)

	if efficiencyLO > efficiencyJS {
		fmt.Println("Le réseau Lagarias-Odlyzko est plus efficace.")
	} else if efficiencyLO < efficiencyJS {
		fmt.Println("Le réseau Joux-Stern est plus efficace.")
	} else {
		fmt.Println("Les deux réseaux ont la même efficacité.")
	}
}
