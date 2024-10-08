package go_by_example

import (
	"crypto/sha256"
	"fmt"
)

// SHA256 hashes are frequently used to compute short identities for binary or text blobs.
// For example, TLS/SSL certificates use SHA256 to compute a certificate's signature.
// Go implements several hash functions in `crypto/*` packages.

func SHA256Hashes() {
	s := "sha256 this string"

	h := sha256.New()

	// `Write` expects bytes. If you have a string `s`, use `[]byte(s)` to coerce it to bytes.
	h.Write([]byte(s))

	// This gets the finalized hash result as a byte slice. 
	// The argument to `Sum` can be used to append to an existing byte slice: it usually isn't needed.
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	// You can compute other hashes using a similar pattern to the one shown above.
	// For example, to compute SHA512 hashes, import `crypto/sha512` and use `sha512.New()`.
	//
	// Note that if you need cryptographically secure hashes, you should carefully research hash strength.
}