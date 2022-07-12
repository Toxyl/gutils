package gutils

import "fmt"

func UintSliceToStringSlice(parts []uint) []string {
	out := []string{}
	for _, p := range parts {
		out = append(out, fmt.Sprint(p))
	}
	return out
}

func Uint8SliceToStringSlice(parts []uint8) []string {
	out := []string{}
	for _, p := range parts {
		out = append(out, fmt.Sprint(p))
	}
	return out
}

func Uint16SliceToStringSlice(parts []uint16) []string {
	out := []string{}
	for _, p := range parts {
		out = append(out, fmt.Sprint(p))
	}
	return out
}

func Uint32SliceToStringSlice(parts []uint32) []string {
	out := []string{}
	for _, p := range parts {
		out = append(out, fmt.Sprint(p))
	}
	return out
}

func Uint64SliceToStringSlice(parts []uint64) []string {
	out := []string{}
	for _, p := range parts {
		out = append(out, fmt.Sprint(p))
	}
	return out
}

// StringSliceDifference returns the elements in `a` that aren't in `b`.
// from https://stackoverflow.com/a/45428032/3337885
func StringSliceDifference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func ChunkByteSlice(slice []byte, chunkSize int) [][]byte {
	var chunks [][]byte
	for {
		if len(slice) == 0 {
			break
		}

		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}
