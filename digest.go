package github

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

// Digest is a SHA1 hash.
type Digest [sha1.Size]byte

// IsZero returns true if the digest is '0000000000000000000000000000000000000000'.
func (d Digest) IsZero() bool {
	return d == Digest{
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00,
	}
}

// String implements the fmt.Stringer interface.
func (d Digest) String() string {
	return hex.EncodeToString(d[:])
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (d *Digest) UnmarshalText(text []byte) error {
	switch len(text) {
	case 0:
		// Special case, we leave d as nil.
	case hex.EncodedLen(sha1.Size):
		if _, err := hex.Decode(d[:], text); err != nil {
			return fmt.Errorf("hex.Decode for digest failed: %w", err)
		}
	default:
		return fmt.Errorf("invalid digest length: %d", len(text))
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface.
func (d Digest) MarshalText() ([]byte, error) {
	return []byte(hex.EncodeToString(d[:])), nil
}
