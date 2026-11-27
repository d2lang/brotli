package brotli

// EncodeFast compresses src into a complete Brotli stream at quality 1 with
// an 18-bit window. It only requires the two-pass fragment encoder, so programs
// that use this entry point can omit the other compression levels.
func EncodeFast(src []byte) []byte {
	const windowBits = 18
	const blockLimit = 1 << windowBits
	blockSize := min(uint(len(src)), uint(blockLimit))
	scratchSize := min(blockSize, uint(kCompressFragmentTwoPassBlockSize))
	commands := make([]uint32, scratchSize)
	literals := make([]byte, scratchSize)
	storage := make([]byte, 2*blockSize+503) // Same worst case as encoderCompressStreamFast.
	table := make([]int, hashTableSize(maxHashTableSize(fastTwoPassCompressionQuality), blockSize))
	var pending uint16
	var pendingBits byte
	encodeWindowBits(windowBits, false, &pending, &pendingBits)
	var dst []byte
	for {
		size := min(uint(len(src)), uint(blockLimit))
		last := size == uint(len(src))
		storage[0], storage[1] = byte(pending), byte(pending>>8)
		bits := uint(pendingBits)
		tableSize := hashTableSize(maxHashTableSize(fastTwoPassCompressionQuality), size)
		compressFragmentTwoPass(src, size, last, commands, literals, table, tableSize, &bits, storage)
		dst = append(dst, storage[:bits/8]...)
		if last { // The final fragment ends on a byte boundary.
			return dst
		}
		pending = uint16(storage[bits/8])
		pendingBits = byte(bits % 8)
		src = src[size:]
		clear(table)
	}
}
