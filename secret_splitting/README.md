### Secret Splitting

Secret splitting is a process by which a secret may be split into `n` parts, can can only be recreated when all `n` parts are recombined.

This code implements the algorithm described in Bruce Schneier's book "Applied Cryptography":

To break a secret S into `n` parts
1. Generate `n-1` random sequences `{R_1, R_2... R_n-1}`the same length of the sequences
2. Exclusive or (XOR) the secret with each random sequence in turn such that `S' = S XOR R_1 XOR R_2 ... XOR R_n-1`
3. The secrets shares are `{R_1, R_2, ... R_n-1, S'}`
