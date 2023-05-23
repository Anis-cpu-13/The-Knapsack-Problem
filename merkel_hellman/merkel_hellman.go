package merkel_hellman

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math"
	"math/big"
)

/* Définission de type de données des clés publique et privées */
type PublicKey struct {
	M []*big.Int
}

type PrivateKey struct {
	R    []*big.Int
	A, B []*big.Int
}

/* Fonction qui convertie une string en binaire */

func StringToBinary(s string) ([]byte, error) {
	var b []byte

	for _, c := range s {
		if c > 255 {
			return nil, fmt.Errorf("Invalid character '%c' in string", c)
		}
		b = append(b, byte(c))
	}

	binaryMessage := make([]byte, len(b)*8)
	for i := range b {
		for j := 0; j < 8; j++ {
			bit := (b[i] >> uint(7-j)) & 1
			binaryMessage[i*8+j] = byte(bit)
		}
	}

	return binaryMessage, nil
}

func BinaryToString(bits []byte) (string, error) {
	// Vérifier que la longueur de la séquence de bits est un multiple de 8
	if len(bits)%8 != 0 {
		return "", errors.New("Sequence length must be a multiple of 8")
	}

	// Convertir la séquence de bits en une chaîne de caractères
	var bytes []byte
	for i := 0; i < len(bits); i += 8 {
		b := byte(0)
		for j := 0; j < 8; j++ {
			b <<= 1
			if bits[i+j] == byte(1) {
				b |= 1
			}
		}
		bytes = append(bytes, b)
	}

	return string(bytes), nil
}

/* Génèrer un nombre aléatoir compris entre un min et un max */
func generateRandomNumber(min *big.Int, max *big.Int) (num *big.Int) {
	diff := big.NewInt(0)
	if num, err := rand.Int(rand.Reader, diff.Sub(max, min)); err != nil {
		panic(err)
	} else {
		num.Add(num, min)
		return num
	}
}

/* fonction qui génère une suite supercroissante */
func GenerateSuperIncreasingSequence(n int) (r []*big.Int, err error) {
	if n < 2 {
		return nil, errors.New("Length of super increasing sequence should be greater than 1")
	}

	aux := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)
	twoExpN := big.NewInt(0)

	r = make([]*big.Int, n)
	// Set first element of sequence to a value >= 2^n
	twoExpN.Exp(two, big.NewInt(int64(n)), nil)
	offset := generateRandomNumber(one, aux.Sqrt(twoExpN))
	r[0] = offset.Add(offset, twoExpN)

	for i := 1; i < n; i++ {
		offset = generateRandomNumber(one, aux.Sqrt(r[i-1]))
		// Next element of the sequence must be at least 2 times larger than the previous one
		aux.Mul(two, r[i-1])
		r[i] = offset.Add(offset, aux)
	}

	return r, nil
}

/* Fonction qui génère deux nombre premier */
func GenerateCoprimes(n *big.Int) (a *big.Int, b *big.Int, err error) {
	gcd := big.NewInt(0)
	one := big.NewInt(1)
	bitLen := n.BitLen()

	b, err = rand.Prime(rand.Reader, bitLen+1)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to generate prime number: %v", err)
	}

	tries := 10 // nombre d'essais pour trouver un nombre premier aléatoire
	for i := tries; i > 0; i-- {
		a = generateRandomNumber(one, b)

		// Vérifier que A et B sont premiers entre eux
		if gcd.GCD(nil, nil, a, b).Cmp(one) == 0 {
			return a, b, nil
		}
	}

	return nil, nil, fmt.Errorf("Failed to generate coprime number after %d attempts", tries)
}

// Generates a sequence m where each element m[i] is calculated as r[i]*a mod b.
func MulMod(r []*big.Int, a *big.Int, b *big.Int) (m []*big.Int) {
	mul := big.NewInt(0)
	m = make([]*big.Int, len(r))

	for i, ri := range r {
		mul.Mul(a, ri)
		m[i] = big.NewInt(0).Mod(mul, b)
	}

	return m
}

/* Fonction qui génère les paramétres des clefs publics et privées */
func GenerateKeyParameters(r []*big.Int, iterations int) (a []*big.Int, b []*big.Int, m []*big.Int, err error) {
	m = r
	a = make([]*big.Int, iterations)
	b = make([]*big.Int, iterations)

	for i := 0; i < iterations; i++ {
		sum := big.NewInt(0)
		for _, mi := range m {
			sum.Add(sum, mi)
		}

		ai, bi, err := GenerateCoprimes(sum)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("Failed to generate key parameters: %v", err)
		}

		a[i] = ai
		b[i] = bi
		m = MulMod(m, ai, bi)
	}

	return a, b, m, nil
}

/* Fonction qui génrer les clefs publics et privées */
func GenerateKeys(byteSize uint, iterations int) (privKey *PrivateKey, pubKey *PublicKey, err error) {
	if byteSize < 2 {
		return nil, nil, errors.New("Length of public key should be greater than 1 byte")
	}
	if iterations < 1 {
		return nil, nil, errors.New("Number of iterations must be greater than 0")
	}

	bitSize := 8 * byteSize
	n := int(math.Ceil(math.Sqrt(float64(bitSize) / 2.0)))

	r, err := GenerateSuperIncreasingSequence(n)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to generate super increasing sequence: %v", err)
	}

	a, b, m, err := GenerateKeyParameters(r, iterations)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to generate pub/priv key pair: %v", err)
	}

	privKey = &PrivateKey{r, a, b}
	pubKey = &PublicKey{m}

	return privKey, pubKey, nil
}

/* Fonction qui encrypte le message */
func Encrypt(pubKey *PublicKey, message string) (c *big.Int, err error) {
	bits, err := StringToBinary(message)
	if err != nil {
		return nil, err
	}

	if len(pubKey.M) < len(bits) {
		return nil, errors.New("Public key length is not sufficient for the message")
	}

	length := len(bits)
	c = big.NewInt(0)
	for i := 0; i < length; i++ {
		if bits[i] == 1 {
			c.Add(c, pubKey.M[i])
		}
	}

	return c, nil
}

/* Fonction qui décrypte le message */
func Decrypt(privKey *PrivateKey, c *big.Int) (message string, err error) {
	s := new(big.Int).Set(c)
	length := len(privKey.R)

	// Inverser les opérations effectuées dans MulMod pour chaque itération
	for i := len(privKey.A) - 1; i >= 0; i-- {
		aiInv := big.NewInt(0)
		aiInv.ModInverse(privKey.A[i], privKey.B[i])
		s.Mul(s, aiInv)
		s.Mod(s, privKey.B[i])
	}

	// Décoder le message
	bits := make([]byte, length)
	for i := length - 1; i >= 0; i-- {
		ri := privKey.R[i]

		if s.Cmp(ri) >= 0 {
			bits[i] = 1
			s.Sub(s, ri)
		} else {
			bits[i] = 0
		}
	}

	// Ajuster la taille de la séquence de bits pour qu'elle soit un multiple de 8
	bitPadding := 8 - (len(bits) % 8)
	if bitPadding < 8 {
		bits = append(bits, make([]byte, bitPadding)...)
	}

	message, err = BinaryToString(bits)
	if err != nil {
		return "", err
	}

	return message, nil
}
