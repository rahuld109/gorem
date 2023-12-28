package main

import (
	"os"
	"path/filepath"
	"testing"
)

func BenchmarkConvertFile(b *testing.B) {
	tempDir := b.TempDir()
	cssFilePath := filepath.Join(tempDir, "benchmark.css")

	cssContent := `
		body {
			font-size: 16px;
			margin: 20px;
		}
	`
	err := os.WriteFile(cssFilePath, []byte(cssContent), 0644)
	if err != nil {
		b.Fatalf("Error writing benchmark CSS file: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err = convertFile(cssFilePath)
		if err != nil {
			b.Fatalf("Error converting benchmark file: %v", err)
		}
	}

	b.StopTimer()
}
