package ch4

func PC32(arr [32]byte, arr2 [32]byte) (count int) {
	for i := range arr {
		// count += int((arr[i] >> (j * 8) & 1) ^ (arr2[i] >> (j * 8) & 1))
		count += int((arr[i] ^ arr2[i]) & 1)
	}
	return
}
