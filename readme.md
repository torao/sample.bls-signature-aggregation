# BLS Signature Aggregation Sample

This is a repository for creating sample code to perform "BLS signature aggregation" using
[herumi/bls-eth-go-binary](https://github.com/herumi/bls-eth-go-binary).

* **Size**: private key = 32B, public key = 48B, signature = 96B
* It's not depend on the order of public key in signing or verification.

See [main.go](main/main.go) for the specific code. The following is the result of the execution.

<details><summary>▶ $ go run ./main</summary>
<p>

```
$ go run ./main
private key[0] := 52baaa129be27188fcbd7334... (size=32)
public key [0] := 8a5cd776a5065269e586681f... (size=48)
signature  [0] := a1d4c1f860abfb1f59495c6c... (size=96, verify=true)
private key[1] := 11faabcff74ff1f802098b8a... (size=32)
public key [1] := 956037493463611e17c27a0f... (size=48)
signature  [1] := 95e7cbe739a6e7d4e519ab99... (size=96, verify=true)
private key[2] := 72ae23fa291c7a4b917b0b51... (size=32)
public key [2] := 8dd19344e6d2cd1ec57244e4... (size=48)
signature  [2] := 896b5e0a0334c8b96ae789b0... (size=96, verify=true)
private key[3] := 2b3826ca7c1d25237006db79... (size=32)
public key [3] := 83fb82ae5c682b8883ad7c7d... (size=48)
signature  [3] := a527a169b3bc0252d476a283... (size=96, verify=true)
private key[4] := 25fd7b4cd3d2738a7034cf41... (size=32)
public key [4] := 92460e922bf0ef0a4755b46e... (size=48)
signature  [4] := 99e7c21d298169829734ec99... (size=96, verify=true)
private key[5] := 3f1d9c4eb81e21f306c877fa... (size=32)
public key [5] := a7ecc66d6bb4123ef885e4af... (size=48)
signature  [5] := 874d952ed2954c78a528ae84... (size=96, verify=true)
private key[6] := 0575cce0cbfee75f3200f577... (size=32)
public key [6] := 858f5529373da0bf2ac70836... (size=48)
signature  [6] := aca1e26c209dba8d3fa51445... (size=96, verify=true)
private key[7] := 1c7d951329549ab48ca4f151... (size=32)
public key [7] := b413f70c0c398dff9deea31a... (size=48)
signature  [7] := b62bd051b99b4b80f908c343... (size=96, verify=true)
private key[8] := 080d646795fb929244aca531... (size=32)
public key [8] := b6eb49a944d94dcb3190dbdd... (size=48)
signature  [8] := a803bdf7c0035726d3862a1a... (size=96, verify=true)
private key[9] := 12ff25b419b1a99299077015... (size=32)
public key [9] := b70714b8f85d865e2c8c1590... (size=48)
signature  [9] := a0e228afc8194cb9fc40e553... (size=96, verify=true)
aggregated signature := 92c61cce4631726fb171f6b9... (size=96)
verify aggregated signature by all public keys: true
aggregated public key := b601df9d672f160dc17c65a2... (size=48)
verify aggregated signature by aggregated public key: true
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
| Key-Pair Generation | 851ns            | 46,535ns | 15,773ns |
| Sign                   | 568,114ns     | < 1ns           | < 1ns          |
| Verify                 | 1,394,315ns | < 1ns            | < 1ns          |

BLS signature is faster than ed25519 and ECDSA in generating key-pair but takes longer to sign and verify.

## Execution times relative to the number of keys and signatures

The time for signatures and public keys aggregation increases in proportion to their number. However, the verification for the aggregated signature is completed in constant time.

![4528009632284672](https://user-images.githubusercontent.com/836654/80785563-af097800-8bbb-11ea-9568-e3584dd6cb8d.png)
