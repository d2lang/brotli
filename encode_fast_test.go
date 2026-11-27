package brotli

import (
	"bytes"
	"io"
	"math/rand"
	"testing"
)

func TestEncodeFast(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	for _, n := range []int{0, 1, 16, 300, 5000, 1 << 16, 1<<17 - 1, 1 << 17, 1<<18 - 1, 1 << 18, 1<<18 + 1, 3<<18 + 97} {
		for _, random := range []bool{false, true} {
			input := make([]byte, n)
			for i := range input {
				if random {
					input[i] = byte(r.Intn(256))
				} else {
					input[i] = byte((i / 17) % 13)
				}
			}
			encoded := EncodeFast(input)
			decoded, err := io.ReadAll(NewReader(bytes.NewReader(encoded)))
			if err != nil || !bytes.Equal(input, decoded) {
				t.Fatalf("length %d random %v: round trip failed: %v", n, random, err)
			}
			var reference bytes.Buffer
			w := NewWriterOptions(&reference, WriterOptions{Quality: 1, LGWin: 18})
			if _, err := w.Write(input); err != nil {
				t.Fatal(err)
			}
			if err := w.Close(); err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(reference.Bytes(), encoded) {
				t.Fatalf("length %d random %v: got %d compressed bytes, want %d", n, random, len(encoded), reference.Len())
			}
		}
	}
}
