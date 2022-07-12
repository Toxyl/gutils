package gutils

import (
	"fmt"
	"strings"
)

func BytesToHexView(b []byte, bytesPerBlock, blocksPerLine int) []string {
	lines := []string{}
	bytesPerLine := bytesPerBlock * blocksPerLine
	charsPerBlock := bytesPerBlock * 2
	chunks := ChunkByteSlice(b, bytesPerLine)
	for _, chunk := range chunks {
		row := ""
		chars := ""
		blocks := ChunkByteSlice(chunk, bytesPerBlock)
		padBlocks := blocksPerLine - len(blocks)
		for _, block := range blocks {
			rb := fmt.Sprintf("%X", block)
			padChars := charsPerBlock - len(rb)
			if padChars > 0 {
				rb += strings.Repeat(" ", padChars)
			}
			row += rb + " "
			charsBlock := string(block)
			chars += ReplaceNonPrintable(ReplaceNonASCII(charsBlock, "."), ".")
		}
		if padBlocks > 0 {
			row += strings.Repeat(" ", (charsPerBlock+1)*padBlocks)
		}
		row += " " + chars
		lines = append(lines, row)
	}
	return lines
}

func StringToHexView(s string, bytesPerBlock, blocksPerLine int) []string {
	return BytesToHexView([]byte(s), bytesPerBlock, blocksPerLine)
}
