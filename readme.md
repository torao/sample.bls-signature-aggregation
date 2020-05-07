# BLS Signature Aggregation Sample

This is a repository for creating sample code to perform "BLS signature aggregation" using
[herumi/bls-eth-go-binary](https://github.com/herumi/bls-eth-go-binary).

* **Size**: private key = 32B, public key = 48B, signature = 96B
* It's not depend on the order of public key in signing or verification.

See [main.go](main/main.go) for the specific code. The following is the result of the execution.

<details><summary>▶ $ go test -bench . -benchmem</summary>
<p>

```
$ go test -bench . -benchmem
goos: darwin
goarch: amd64
BenchmarkBLS/Key-Pair_Generation-8                 16455             73486 ns/op             176 B/op          2 allocs/op
BenchmarkBLS/Sign-8                                 2144            552673 ns/op             336 B/op          3 allocs/op
BenchmarkBLS/Verify-8                                830           1386782 ns/op              48 B/op          2 allocs/op
BenchmarkBLS/Signatures_Aggregation[1]-8        10985703               113 ns/op             288 B/op          1 allocs/op
BenchmarkBLS/Signatures_Aggregation[2]-8          312457              3441 ns/op             288 B/op          1 allocs/op
BenchmarkBLS/Signatures_Aggregation[4]-8           62047             20311 ns/op             288 B/op          1 allocs/op
BenchmarkBLS/Signatures_Aggregation[10]-8           7959            145796 ns/op             288 B/op          1 allocs/op
BenchmarkBLS/Signatures_Aggregation[21]-8           1725            681825 ns/op             288 B/op          1 allocs/op
BenchmarkBLS/Signatures_Aggregation[46]-8            352           3459057 ns/op             288 B/op          1 allocs/op
BenchmarkBLS/Signatures_Aggregation[100]-8            70          16587340 ns/op             288 B/op          1 allocs/op
BenchmarkBLS/Aggregated_Signature_Verification[1]-8                  862           1394435 ns/op              48 B/op          2 allocs/op
BenchmarkBLS/Aggregated_Signature_Verification[2]-8                  834           1393686 ns/op              48 B/op          2 allocs/op
BenchmarkBLS/Aggregated_Signature_Verification[4]-8                  836           1396647 ns/op              48 B/op          2 allocs/op
BenchmarkBLS/Aggregated_Signature_Verification[10]-8                 858           1436367 ns/op              48 B/op          2 allocs/op
BenchmarkBLS/Aggregated_Signature_Verification[21]-8                 816           1413289 ns/op              48 B/op          2 allocs/op
BenchmarkBLS/Aggregated_Signature_Verification[46]-8                 840           1436811 ns/op              48 B/op          2 allocs/op
BenchmarkBLS/Aggregated_Signature_Verification[100]-8                820           1453223 ns/op              48 B/op          2 allocs/op
BenchmarkBLS/Public_Keys_Aggregation[1]-8                       10384542                97.0 ns/op           144 B/op          1 allocs/op
BenchmarkBLS/Public_Keys_Aggregation[2]-8                        1663330               727 ns/op             144 B/op          1 allocs/op
BenchmarkBLS/Public_Keys_Aggregation[4]-8                         582181              2024 ns/op             144 B/op          1 allocs/op
BenchmarkBLS/Public_Keys_Aggregation[10]-8                        204166              5822 ns/op             144 B/op          1 allocs/op
BenchmarkBLS/Public_Keys_Aggregation[21]-8                         93244             13091 ns/op             144 B/op          1 allocs/op
BenchmarkBLS/Public_Keys_Aggregation[46]-8                         39918             29182 ns/op             144 B/op          1 allocs/op
BenchmarkBLS/Public_Keys_Aggregation[100]-8                        18784             64513 ns/op             144 B/op          1 allocs/op
BenchmarkBLS/Aggregated_Signature_Verification[1]_by_Aggregated_Public_Key-8                 861           1402853 ns/op              48 B/op          2 allocs/op
BenchmarkBLS/Aggregated_Signature_Verification[2]_by_Aggregated_Public_Key-8                 855           1426795 ns/op              48 B/op          2 allocs/op
BenchmarkBLS/Aggregated_Signature_Verification[4]_by_Aggregated_Public_Key-8                 686           1662080 ns/op              48 B/op          2 allocs/op
BenchmarkBLS/Aggregated_Signature_Verification[10]_by_Aggregated_Public_Key-8                790           1493050 ns/op              48 B/op          2 allocs/op
BenchmarkBLS/Aggregated_Signature_Verification[21]_by_Aggregated_Public_Key-8                822           1500175 ns/op              48 B/op          2 allocs/op
BenchmarkBLS/Aggregated_Signature_Verification[46]_by_Aggregated_Public_Key-8                774           1508392 ns/op              48 B/op          2 allocs/op
BenchmarkBLS/Aggregated_Signature_Verification[100]_by_Aggregated_Public_Key-8               846           1421861 ns/op              48 B/op          2 allocs/op
[BLS] private key: 32 bytes, public key: 48 bytes, signature: 96 bytes
BenchmarkEd25519/Key-Pair_Generation-8                                                     25320             46336 ns/op             128 B/op          3 allocs/op
BenchmarkEd25519/Sign-8                                                                    25587             47036 ns/op             512 B/op          6 allocs/op
BenchmarkEd25519/Verify-8                                                                   9686            124300 ns/op             288 B/op          2 allocs/op
[Ed25519] private key: 64 bytes, public key: 32 bytes, signature: 64 bytes
BenchmarkECDSA/Key-Pair_Generation-8                                                       80736             14711 ns/op             608 B/op         12 allocs/op
BenchmarkECDSA/Sign-8                                                                      51530             23724 ns/op            2673 B/op         32 allocs/op
BenchmarkECDSA/Verify-8                                                                    17085             69779 ns/op             880 B/op         16 allocs/op
[ECDSA] private key: 96 bytes, public key: 64 bytes, signature: 64 bytes
PASS
ok      github.com/herumi/bls-eth-go-binary     51.882s

```

</p>
</details>

In aggregated signature and public key, I generated 10 key-pairs, created 10 signatures for one message, and aggregated them into a single signature. It's able to verify the aggregated signature using 10 public keys or one public key that aggregates them.

The order in which signatures are aggregated and the order of public keys has no effect on verification.

## Key and Signature Size Comparison

| Data | BLS | Ed25519 | ECDSA (p256) |
|:-----|-----:|----:|----:|
| Private Key | 32 bytes | 64 bytes | 96 bytes |
| Public Key | 48 bytes | 32 bytes | 64 bytes |
| Signature | 96 bytes | 64 bytes | 64 bytes |
| Aggregated Signature | 96 bytes | - | - |
| Aggregated Public Key | 48 bytes | - | - |

The table above shows that aggregated public key and aggregated signature have the same size as regular public key and signature.

## Comparison with other elliptic-curve cryptography

Follows are the execution times for 1) key-pair generation, 2) signing with a single private key, and 3) verifying a single signature with a single public key by [bls_test.go](https://github.com/torao/sample.bls-signature-aggregation/blob/master/bls_test.go):

| | BLS | Ed25519 | ECDSA (p256) |
|:----------------|-------------:|----------:|----------:|
| Key-Pair Generation | 73,486ns            | 46,336ns | 14,711ns |
| Sign                   | 552,673ns     | 47,036ns           | 23,724ns          |
| Verify                 | 1,386,782ns | 124,300ns            | 69,779ns          |

BLS signature is faster than ed25519 and ECDSA in generating key-pair but takes longer to sign and verify.

## Execution times relative to the number of keys and signatures

The time for signatures and public keys aggregation increases in proportion to their number. However, the verification for the aggregated signature is completed in constant time.

![4528009632284672](https://user-images.githubusercontent.com/836654/80785563-af097800-8bbb-11ea-9568-e3584dd6cb8d.png)
