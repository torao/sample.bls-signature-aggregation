package main

import (
	"fmt"
	"math/rand"
	"github.com/herumi/bls-eth-go-binary/bls"
	"time"
)

func main() {
	err := bls.Init(bls.BLS12_381)
	if err != nil {
		panic(fmt.Sprintf("ERROR: %s", err))
	}

	// generate BLS keys and signatures
	message := []byte("hello, world")
	keyCount := 10
	sKeys := make([]bls.SecretKey, keyCount)
	pKeys := make([]bls.PublicKey, keyCount)
	signs := make([]bls.Sign, keyCount)
	for i := 0; i < keyCount; i++ {
		sKeys[i] = bls.SecretKey{}
		sKeys[i].SetByCSPRNG()
		pKeys[i] = *sKeys[i].GetPublicKey()
		signs[i] = *sKeys[i].SignByte(message)
		success := signs[i].VerifyByte(&pKeys[i], message)
		fmt.Printf("private key[%d] := %s... (size=%d)\n", i, sKeys[i].SerializeToHexStr()[:24], len(sKeys[i].Serialize()))
		fmt.Printf("public key [%d] := %s... (size=%d)\n", i, pKeys[i].SerializeToHexStr()[:24], len(pKeys[i].Serialize()))
		fmt.Printf("signature  [%d] := %s... (size=%d, verify=%v)\n", i, signs[i].SerializeToHexStr()[:24], len(signs[i].Serialize()), success)
	}

	// generate aggregated signature
	aggrSign := bls.Sign{}
	aggrSign.Aggregate(signs)
	fmt.Printf("aggregated signature := %s... (size=%d)\n", aggrSign.SerializeToHexStr()[:24], len(aggrSign.Serialize()))

	// verify aggregated signature by all public keys
	// shuffle public keys to show that it doesn't depend on their order.
	shuffledPKeys := make([]bls.PublicKey, keyCount)
	copy(shuffledPKeys, pKeys)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shuffledPKeys), func(i, j int) { shuffledPKeys[i], shuffledPKeys[j] = shuffledPKeys[j], shuffledPKeys[i] })
	success := aggrSign.FastAggregateVerify(shuffledPKeys, message)
	fmt.Printf("verify aggregated signature by all public keys: %v\n", success)

	// verify aggregated signature by aggregated public key
	aggrPKey := bls.PublicKey{}
	for i:=0; i<keyCount; i++ {
		aggrPKey.Add(&shuffledPKeys[i])
	}
	fmt.Printf("aggregated public key := %s... (size=%d)\n", aggrPKey.SerializeToHexStr()[:24], len(aggrPKey.Serialize()))
	success = aggrSign.VerifyByte(&aggrPKey, message)
	fmt.Printf("verify aggregated signature by aggregated public key: %v\n", success)
}
