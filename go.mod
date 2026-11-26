module github.com/d2lang/brotli

go 1.22

// This fork carries upstream's tags (v1.0.0 to v1.3.0), whose go.mod files
// declare github.com/andybalholm/brotli, so they cannot be used under this
// module path. v1.3.1 only exists to publish this retraction; use v0.x.
retract [v1.0.0, v1.3.1]
