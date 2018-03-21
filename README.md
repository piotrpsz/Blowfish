# Blowfish

Implementation of the Blowfish encryption algorithm (the creator of the algorithm is Bruce Scheiner).
It must be clear that the code is not tuned for speed - main goal is explanation how works the algorithm.

To test the correctness of the operation, run the following program:
# Example: How to test

```Go
func main () {
    fmt.Println ()

    L: = uint32 (1)
    R: = uint32 (2)
    bf: = blowfish.New ([] byte ("TESTKEY"))

    bf.Encrypt (& L, & R)
    fmt.Printf ("% 08x,% 08x \ n", L, R)

    if (L == 0xdf333fd2 && R == 0x30a71bb4) {
        fmt.Println ("Test encryption OK.")
    } else {
        fmt.Println ("Test encryption failed.")
    }

    bf.Decrypt (& L, & R)
    if (L == 1 && R == 2) {
        fmt.Println ("Test decryption OK.")
    } else {
        fmt.Println ("Test decryption failed.");
    }
}
```
