/*
Copyright 2013 The Perkeep Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package hashutil contains misc hashing functions lacking homes elsewhere.
package hashutil // import "perkeep.org/internal/hashutil"

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

// SHA256Prefix computes the SHA-256 digest of data and returns
// its first twenty lowercase hex digits.
func SHA256Prefix(data []byte) string {
	h := sha256.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil))[:20]
}

// SHA1Prefix computes the SHA-1 digest of data and returns
// its first twenty lowercase hex digits.
func SHA1Prefix(data []byte) string {
	h := sha1.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil))[:20]
}
