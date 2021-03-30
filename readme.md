# BLS Signature Aggregation Sample

This is a repository for creating sample code to perform "BLS signature aggregation" using
[herumi/bls-go-binary](https://github.com/herumi/bls-go-binary).

Signature aggregation is the aggregation of multiple signatures, generated by multiple private keys, into a single signature. The aggregated signature can be verified using multiple public keys that are pairs of the private keys that generate the signature. In addition, BLS public keys can aggregate multiple public keys.

See [main.go](main/main.go) for the specific code. The following is the result of the execution.

<details><summary>▶ $ go test -bench . -benchmem</summary>
<p>

```
% go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/herumi/bls-eth-go-binary
BenchmarkBLS/Key-Pair_Generation-8   	    5884	    202213 ns/op	     320 B/op	       2 allocs/op
BenchmarkBLS/Sign-8                  	    3427	    387399 ns/op	     176 B/op	       2 allocs/op
BenchmarkBLS/Verify-8                	     751	   1644379 ns/op	      32 B/op	       1 allocs/op
[BLS] private key: 32 bytes, public key: 96 bytes, signature: 48 bytes
BenchmarkEd25519/Key-Pair_Generation-8         	   17025	     65943 ns/op	     128 B/op	       3 allocs/op
BenchmarkEd25519/Sign-8                        	   20230	     58491 ns/op	     448 B/op	       5 allocs/op
BenchmarkEd25519/Verify-8                      	    8234	    149904 ns/op	     288 B/op	       2 allocs/op
[Ed25519] private key: 64 bytes, public key: 32 bytes, signature: 64 bytes
BenchmarkECDSA/Key-Pair_Generation-8           	   64735	     17776 ns/op	     608 B/op	      12 allocs/op
BenchmarkECDSA/Sign-8                          	   40912	     29624 ns/op	    2673 B/op	      32 allocs/op
BenchmarkECDSA/Verify-8                        	   13507	     86202 ns/op	     880 B/op	      16 allocs/op
[ECDSA] private key: 95 bytes, public key: 63 bytes, signature: 64 bytes
BenchmarkBLSSignatureAggregation/Signatures_Aggregation[1]-8         	10897312	       115 ns/op	     144 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Signatures_Aggregation[2]-8         	 1433745	       830 ns/op	     144 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Signatures_Aggregation[4]-8         	  494150	      2383 ns/op	     144 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Signatures_Aggregation[10]-8        	  182233	      7164 ns/op	     144 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Signatures_Aggregation[21]-8        	   79200	     15652 ns/op	     144 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Signatures_Aggregation[46]-8        	   33891	     35048 ns/op	     144 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Signatures_Aggregation[100]-8       	   15133	     80607 ns/op	     144 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Aggregated_Signature_Verification[1]-8         	     784	   1526823 ns/op	      32 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Aggregated_Signature_Verification[2]-8         	     736	   1395230 ns/op	      32 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Aggregated_Signature_Verification[4]-8         	     878	   1452489 ns/op	      32 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Aggregated_Signature_Verification[10]-8        	     835	   1509504 ns/op	      32 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Aggregated_Signature_Verification[21]-8        	     834	   1494399 ns/op	      32 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Aggregated_Signature_Verification[46]-8        	     831	   1535658 ns/op	      32 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Aggregated_Signature_Verification[100]-8       	     708	   1625537 ns/op	      32 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Public_Keys_Aggregation[1]-8                   	 7220750	       156 ns/op	     288 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Public_Keys_Aggregation[2]-8                   	  546836	      2168 ns/op	     288 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Public_Keys_Aggregation[4]-8                   	  198849	      6242 ns/op	     288 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Public_Keys_Aggregation[10]-8                  	   62683	     18470 ns/op	     288 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Public_Keys_Aggregation[21]-8                  	   29636	     39782 ns/op	     288 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Public_Keys_Aggregation[46]-8                  	   13272	     93353 ns/op	     288 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Public_Keys_Aggregation[100]-8                 	    5840	    208423 ns/op	     288 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Aggregated_Signature_Verification[1]_by_Aggregated_Public_Key-8         	     856	   1324273 ns/op	      32 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Aggregated_Signature_Verification[2]_by_Aggregated_Public_Key-8         	     916	   1361653 ns/op	      32 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Aggregated_Signature_Verification[4]_by_Aggregated_Public_Key-8         	     903	   1389915 ns/op	      32 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Aggregated_Signature_Verification[10]_by_Aggregated_Public_Key-8        	     879	   1412366 ns/op	      32 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Aggregated_Signature_Verification[21]_by_Aggregated_Public_Key-8        	     862	   1372903 ns/op	      32 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Aggregated_Signature_Verification[46]_by_Aggregated_Public_Key-8        	     885	   1376454 ns/op	      32 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/Aggregated_Signature_Verification[100]_by_Aggregated_Public_Key-8       	     879	   1349361 ns/op	      32 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/AggregateSigsForDiffMsg[1]-8                                            	 9442695	       128 ns/op	     144 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/AggregateSigsForDiffMsg[2]-8                                            	 1378105	       873 ns/op	     144 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/AggregateSigsForDiffMsg[4]-8                                            	  505824	      2496 ns/op	     144 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/AggregateSigsForDiffMsg[10]-8                                           	  167084	      7789 ns/op	     144 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/AggregateSigsForDiffMsg[21]-8                                           	   73917	     16116 ns/op	     144 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/AggregateSigsForDiffMsg[46]-8                                           	   33160	     36805 ns/op	     144 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/AggregateSigsForDiffMsg[100]-8                                          	   14854	     82897 ns/op	     144 B/op	       1 allocs/op
BenchmarkBLSSignatureAggregation/VerifyAggrSigsForDiffMsg[1]-8                                           	     782	   1500140 ns/op	      64 B/op	       2 allocs/op
BenchmarkBLSSignatureAggregation/VerifyAggrSigsForDiffMsg[2]-8                                           	     597	   1980438 ns/op	      96 B/op	       2 allocs/op
BenchmarkBLSSignatureAggregation/VerifyAggrSigsForDiffMsg[4]-8                                           	     408	   2961081 ns/op	     160 B/op	       2 allocs/op
BenchmarkBLSSignatureAggregation/VerifyAggrSigsForDiffMsg[10]-8                                          	     206	   5643624 ns/op	     352 B/op	       2 allocs/op
BenchmarkBLSSignatureAggregation/VerifyAggrSigsForDiffMsg[21]-8                                          	     100	  11072527 ns/op	     736 B/op	       2 allocs/op
BenchmarkBLSSignatureAggregation/VerifyAggrSigsForDiffMsg[46]-8                                          	      51	  23527329 ns/op	    1568 B/op	       2 allocs/op
BenchmarkBLSSignatureAggregation/VerifyAggrSigsForDiffMsg[100]-8                                         	      22	  50250564 ns/op	    3232 B/op	       2 allocs/op
PASS
ok  	github.com/herumi/bls-eth-go-binary	74.757s
```

</p>
</details>

In aggregated signature and public key, I generated 10 key-pairs, created 10 signatures for one message, and aggregated them into a single signature. It's able to verify the aggregated signature using 10 public keys or one public key that aggregates them.

The order in which signatures are aggregated and the order of public keys has no effect on verification.

## Key and Signature Size Comparison

| Data | BLS | BLS (ETH) | Ed25519 | ECDSA (p256) |
|:-----|----:|----------:|--------:|-------------:|
| Private Key | 32 bytes | 32 bytes | 64 bytes | 96 bytes |
| Public Key | 96 bytes | 48 bytes | 32 bytes | 64 bytes |
| Signature | 48 bytes | 96 bytes | 64 bytes | 64 bytes |
| Aggregated Signature | 48 bytes | 96 bytes | - | - |
| Aggregated Public Key | 96 bytes | 48 bytes | - | - |

The table above shows that aggregated public key and aggregated signature have the same size as regular public key and signature.

## Comparison with other elliptic-curve cryptography

Follows are the execution times for 1) key-pair generation, 2) signing with a single private key, and 3) verifying a single signature with a single public key by [bls_test.go](https://github.com/torao/sample.bls-signature-aggregation/blob/master/bls_test.go):

|                     | BLS         | Ed25519   | ECDSA (p256) |
|:--------------------|------------:|----------:|-------------:|
| Key-Pair Generation | 202,213ns   | 65,943ns  | 17,776ns     |
| Sign                | 387,399ns   | 58,491ns  | 29,624ns     |
| Verify              | 1,644,379ns | 149,904ns | 86,202ns     |

We can see that BLS signatures are generally slower than ed25519 and ECDSA.

## Execution times relative to the number of keys and signatures

The time for signatures and public keys aggregation increases in proportion to their number. However, the verification for the aggregated signature is completed in constant time. In the case of signatures for different messages, the verification time is proportional to the number of signatures, even for aggregated signatures.

![2021-03-30 macOS i7 2.8GHz](https://user-images.githubusercontent.com/836654/112939492-88a2e280-9166-11eb-9a14-5b1336885d3b.png)
