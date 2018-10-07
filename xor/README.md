### XOR Cypher

Here I implement an XOR cypher [1] for encryption and decryption. This cypher can be used to provide theoretically perfect security if they key length is the same as the plaintext length [2], however, with shorter keys it is extremely weak. I also implement an attack that breaks the cypher, even with "long" key lengths of >20 characters.

1: https://en.wikipedia.org/wiki/XOR_cipher
2: https://en.wikipedia.org/wiki/One-time_pad
