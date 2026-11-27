# d2lang/brotli

A trimmed fork of [andybalholm/brotli](https://github.com/andybalholm/brotli),
a pure-Go brotli compressor and decompressor translated from the reference
implementation (https://github.com/google/brotli) with the `c2go` tool at
https://github.com/andybalholm/c2go. All credit for the code goes to its
upstream authors; see LICENSE.

This fork exists for [D2](https://github.com/d2lang/d2), which compiles to
WebAssembly. Upstream's `http.go` imports `net/http`, and importing a package
links its initialization (and much of `crypto/x509`, `encoding/asn1` and
friends) into the WebAssembly module even though nothing calls it. This fork
keeps only what D2 uses:

- `NewReader`, `Reader` (decoder)
- `EncodeFast` (one-shot quality 1 encoder with an 18-bit window)
- `NewWriter`, `NewWriterLevel`, `NewWriterOptions`, `WriterOptions`, `Writer`
  and the compression level constants (the encoder translated from C)

Removed from upstream v1.3.0:

- `http.go` (`HTTPCompressor`, `HTTPCompressorWithLevel`)
- the `matchfinder` package, `NewWriterV2`, `Encoder` and `FastEncoder`
- the `flate` package and the `corpus-test` command
- `NewReaderOptions` and `ReaderOptions` (large-window decoding)
- the tests of the removed code, and with them the test dependency on
  `github.com/xyproto/randomstring`
- `SECURITY.md` (report vulnerabilities in the code itself upstream)

The remaining code is unchanged from upstream apart from deleting unexported
helpers that nothing calls any more. If you are not D2, use upstream.
