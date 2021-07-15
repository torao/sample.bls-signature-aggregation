package sample_bls_signature_aggregation

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/herumi/bls-go-binary/bls"
	"math"
	"testing"
)

var Message = []byte("hello, world")
var AggregateNumbers = []int{
	int(math.Pow(10, 1.0/3.0*0)),
	int(math.Pow(10, 1.0/3.0*1)),
	int(math.Pow(10, 1.0/3.0*2)),
	int(math.Pow(10, 1.0/3.0*3)),
	int(math.Pow(10, 1.0/3.0*4)),
	int(math.Pow(10, 1.0/3.0*5)),
	int(math.Pow(10, 1.0/3.0*6)),
}

func BenchmarkBLS(b *testing.B) {
	if err := bls.Init(bls.BLS12_381); err != nil {
		panic(err)
	}

	b.Run("Key-Pair Generation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			privateKey := bls.SecretKey{}
			privateKey.SetByCSPRNG()
			privateKey.GetPublicKey()
		}
	})

	privateKey := bls.SecretKey{}
	privateKey.SetByCSPRNG()
	b.Run("Sign", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			privateKey.SignByte(Message)
		}
	})

	signature := privateKey.SignByte(Message)
	publicKey := privateKey.GetPublicKey()
	b.Run("Verify", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			signature.VerifyByte(publicKey, Message)
		}
	})

	if !signature.VerifyByte(publicKey, Message) {
		panic("BLS verification failed!")
	}

	fmt.Printf("[BLS] private key: %d bytes, public key: %d bytes, signature: %d bytes\n",
		len(privateKey.Serialize()), len(publicKey.Serialize()), len(signature.Serialize()))
}

func BenchmarkEd25519(b *testing.B) {

	b.Run("Key-Pair Generation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ed25519.GenerateKey(rand.Reader)
		}
	})

	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic("Ed25519 key generation failed!")
	}
	b.Run("Sign", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ed25519.Sign(privateKey, Message)
		}
	})

	signature := ed25519.Sign(privateKey, Message)
	b.Run("Verify", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ed25519.Verify(publicKey, Message, signature)
		}
	})

	if !ed25519.Verify(publicKey, Message, signature) {
		panic("Ed25519 verification failed!")
	}

	fmt.Printf("[Ed25519] private key: %d bytes, public key: %d bytes, signature: %d bytes\n",
		len(privateKey), len(publicKey), len(signature))
}

func BenchmarkP256(b *testing.B) {

	p256 := elliptic.P256()
	seed := rand.Reader
	b.Run("Key-Pair Generation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			privateKey, _ := ecdsa.GenerateKey(p256, seed)
			privateKey.Public()
		}
	})

	privateKey, err := ecdsa.GenerateKey(p256, seed)
	if err != nil {
		panic("ECDSA key generation failed!")
	}
	b.Run("Sign", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ecdsa.Sign(seed, privateKey, Message)
		}
	})

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, Message)
	if err != nil {
		panic("ECDSA signature failed!")
	}
	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	b.Run("Verify", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ecdsa.Verify(publicKey, Message, r, s)
		}
	})

	if !ecdsa.Verify(publicKey, Message, r, s) {
		panic("ECDSA verification failed!")
	}

	fmt.Printf("[ECDSA (P-256)] private key: %d bytes, public key: %d bytes, signature: %d bytes\n",
		len(privateKey.D.Bytes())+len(privateKey.X.Bytes())+len(privateKey.Y.Bytes()),
		len(publicKey.Y.Bytes())+len(publicKey.X.Bytes()),
		len(signature))
}

func BenchmarkSecp256k1(b *testing.B) {
	secp256k1 := crypto.S256()
	seed := rand.Reader
	b.Run("Key-Pair Generation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			privateKey, _ := ecdsa.GenerateKey(secp256k1, seed)
			privateKey.Public()
		}
	})

	privateKey, err := ecdsa.GenerateKey(secp256k1, seed)
	if err != nil {
		panic("ECDSA key generation failed!")
	}
	b.Run("Sign", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ecdsa.Sign(seed, privateKey, Message)
		}
	})

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, Message)
	if err != nil {
		panic("ECDSA signature failed!")
	}
	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	b.Run("Verify", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ecdsa.Verify(publicKey, Message, r, s)
		}
	})

	if !ecdsa.Verify(publicKey, Message, r, s) {
		panic("ECDSA verification failed!")
	}

	fmt.Printf("[ECDSA (secp256k1)] private key: %d bytes, public key: %d bytes, signature: %d bytes\n",
		len(privateKey.D.Bytes())+len(privateKey.X.Bytes())+len(privateKey.Y.Bytes()),
		len(publicKey.Y.Bytes())+len(publicKey.X.Bytes()),
		len(signature))
}

func BenchmarkBLSSignatureAggregation(b *testing.B) {
	if err := bls.Init(bls.BLS12_381); err != nil {
		panic(err)
	}
	privateKey := bls.SecretKey{}
	privateKey.SetByCSPRNG()

	privateKeys := make([]bls.SecretKey, 100)
	publicKeys := make([]*bls.PublicKey, len(privateKeys))
	signatures := make([]bls.Sign, len(privateKeys))
	for i := 0; i < len(privateKeys); i++ {
		privateKeys[i] = bls.SecretKey{}
		privateKeys[i].SetByCSPRNG()
		publicKeys[i] = privateKeys[i].GetPublicKey()
		signatures[i] = *privateKeys[i].SignByte(Message)
	}

	for _, sigs := range AggregateNumbers {
		b.Run(fmt.Sprintf("Signatures Aggregation[%d]", sigs), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				aggregatedSignature := bls.Sign{}
				aggregatedSignature.Aggregate(signatures[:sigs])
			}
		})
	}

	aggregatedSignatures := make([]*bls.Sign, len(AggregateNumbers))
	for x, sigs := range AggregateNumbers {
		aggregatedSignatures[x] = &bls.Sign{}
		aggregatedSignatures[x].Aggregate(signatures[:sigs])
		publicKeyEntities := make([]bls.PublicKey, sigs)
		for i := 0; i < sigs; i++ {
			publicKeyEntities[i] = *publicKeys[i]
		}
		b.Run(fmt.Sprintf("Aggregated Signature Verification[%d]", sigs), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				aggregatedSignatures[x].FastAggregateVerify(publicKeyEntities[:sigs], Message)
			}
		})
		if !aggregatedSignatures[x].FastAggregateVerify(publicKeyEntities[:sigs], Message) {
			panic("BLS aggregated signature verification failed!")
		}
	}

	for _, keys := range AggregateNumbers {
		b.Run(fmt.Sprintf("Public Keys Aggregation[%d]", keys), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				aggregatedPublicKey := bls.PublicKey{}
				for j := 0; j < keys; j++ {
					aggregatedPublicKey.Add(publicKeys[j])
				}
			}
		})
	}

	aggregatedPublicKeys := make([]*bls.PublicKey, len(AggregateNumbers))
	for x, keys := range AggregateNumbers {
		aggregatedPublicKeys[x] = &bls.PublicKey{}
		for j := 0; j < keys; j++ {
			aggregatedPublicKeys[x].Add(publicKeys[j])
		}
		b.Run(fmt.Sprintf("Aggregated Signature Verification[%d] by Aggregated Public Key", keys), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				aggregatedSignatures[x].VerifyByte(aggregatedPublicKeys[x], Message)
			}
		})
		if !aggregatedSignatures[x].VerifyByte(aggregatedPublicKeys[x], Message) {
			panic("BLS aggregated signature verification by aggregated public key failed!")
		}
	}

	hashes := make([][]byte, len(privateKeys))
	for i := 0; i < len(privateKeys); i++ {
		hash := sha256.Sum256([]byte(fmt.Sprintf("hello, workd: %d", i)))
		hashes[i] = hash[:]
		signatures[i] = *privateKeys[i].SignHash(hashes[i])
	}

	for _, sigs := range AggregateNumbers {
		b.Run(fmt.Sprintf("AggregateSigsForDiffMsg[%d]", sigs), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				aggregatedSignature := bls.Sign{}
				for j := 0; j < sigs; j++ {
					aggregatedSignature.Add(&signatures[j])
				}
			}
		})
	}

	for _, sigs := range AggregateNumbers {
		aggregatedSignature := bls.Sign{}
		pubKeys := make([]bls.PublicKey, sigs)
		for j := 0; j < sigs; j++ {
			aggregatedSignature.Add(&signatures[j])
			pubKeys[j] = *publicKeys[j]
		}
		b.Run(fmt.Sprintf("VerifyAggrSigsForDiffMsg[%d]", sigs), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				aggregatedSignature.VerifyAggregateHashes(pubKeys, hashes[:sigs])
			}
		})
		if !aggregatedSignature.VerifyAggregateHashes(pubKeys, hashes[:sigs]) {
			panic("BLS aggregated signature verification failed")
		}
	}
}
