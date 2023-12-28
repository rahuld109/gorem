package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestConvertFileWithFractionalPixels(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Create a sample CSS file with fractional pixel values
	cssContent := `
		body {
			font-size: 16.88px;
			margin: 13.33px;
		}
	`
	cssFilePath := filepath.Join(tempDir, "test_fractional_pixels.css")
	err := os.WriteFile(cssFilePath, []byte(cssContent), 0644)
	if err != nil {
		t.Fatalf("Error writing sample CSS file: %v", err)
	}

	// Run the convertFile function on the test CSS file
	err = convertFile(cssFilePath)
	if err != nil {
		t.Fatalf("Error converting file: %v", err)
	}

	// Read the modified content of the CSS file
	modifiedContent, err := os.ReadFile(cssFilePath)
	if err != nil {
		t.Fatalf("Error reading modified content: %v", err)
	}

	// Check if the content contains the expected rem values for fractional pixels
	expectedRemValues := []string{"1.05rem", "0.83rem"}
	for _, expectedRemValue := range expectedRemValues {
		if !strings.Contains(string(modifiedContent), expectedRemValue) {
			t.Errorf("Expected content to contain %s, got %s", expectedRemValue, string(modifiedContent))
		}
	}
}

func TestConvertEmptyFile(t *testing.T) {
	tempDir := t.TempDir()
	cssFilePath := filepath.Join(tempDir, "empty.css")

	err := os.WriteFile(cssFilePath, []byte(""), 0644)
	if err != nil {
		t.Fatalf("Error writing empty CSS file: %v", err)
	}

	err = convertFile(cssFilePath)
	if err != nil {
		t.Fatalf("Error converting empty file: %v", err)
	}

	// Read the modified content of the CSS file
	modifiedContent, err := os.ReadFile(cssFilePath)
	if err != nil {
		t.Fatalf("Error reading modified content: %v", err)
	}

	// Check if the content is still empty
	if string(modifiedContent) != "" {
		t.Errorf("Expected content to be empty, got %s", string(modifiedContent))
	}
}

func TestConvertFileNoPixelValues(t *testing.T) {
	tempDir := t.TempDir()
	cssFilePath := filepath.Join(tempDir, "no_pixel_values.css")

	cssContent := `
		body {
			font-family: 'Arial', sans-serif;
			color: #333;
		}
	`
	err := os.WriteFile(cssFilePath, []byte(cssContent), 0644)
	if err != nil {
		t.Fatalf("Error writing CSS file without pixel values: %v", err)
	}

	err = convertFile(cssFilePath)
	if err != nil {
		t.Fatalf("Error converting file without pixel values: %v", err)
	}

	// Read the modified content of the CSS file
	modifiedContent, err := os.ReadFile(cssFilePath)
	if err != nil {
		t.Fatalf("Error reading modified content: %v", err)
	}

	// Check if the content remains unchanged
	if string(modifiedContent) != cssContent {
		t.Errorf("Expected content to be unchanged, got %s", string(modifiedContent))
	}
}

func TestConvertFile(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Create a sample CSS file in the temporary directory
	cssContent := `
		body {
			font-size: 16px;
			margin: 20px;
		}
	`
	cssFilePath := filepath.Join(tempDir, "test.css")
	err := os.WriteFile(cssFilePath, []byte(cssContent), 0644)
	if err != nil {
		t.Fatalf("Error writing sample CSS file: %v", err)
	}

	// Run the convertFile function on the test CSS file
	err = convertFile(cssFilePath)
	if err != nil {
		t.Fatalf("Error converting file: %v", err)
	}

	// Read the modified content of the CSS file
	modifiedContent, err := os.ReadFile(cssFilePath)
	if err != nil {
		t.Fatalf("Error reading modified content: %v", err)
	}

	// Check if the content contains the expected rem value
	expectedRemValue := "1.00rem"
	if !strings.Contains(string(modifiedContent), expectedRemValue) {
		t.Errorf("Expected content to contain %s, got %s", expectedRemValue, string(modifiedContent))
	}
}
