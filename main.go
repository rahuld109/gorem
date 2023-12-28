package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

const (
	pixelBase = 16.0
	filePerm  = 0644
	version   = "1.0.0"
)

var rootCmd = &cobra.Command{
	Use:     "gorem",
	Short:   "Converts pixel values to rem in CSS files",
	Run:     wrapConvertCSSFiles,
	Version: version,
}

func wrapConvertCSSFiles(cmd *cobra.Command, args []string) {
	err := convertCSSFiles(cmd, args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func convertCSSFiles(_ *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: gorem <directory>")
	}

	directory := args[0]
	files, err := getCSSFiles(directory)
	if err != nil {
		return fmt.Errorf("error getting CSS files: %v", err)
	}

	var wg sync.WaitGroup
	var successCount int32

	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			err := convertFile(file)
			if err != nil {
				fmt.Printf("Error converting file %s: %v\n", file, err)
			} else {
				atomic.AddInt32(&successCount, 1)
			}
		}(file)
	}

	wg.Wait()

	fmt.Printf("Conversion complete. Successfully converted %d files.\n", successCount)

	return nil
}

func getCSSFiles(directory string) ([]string, error) {
	var files []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".css" {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func convertFile(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`(\d+(\.\d+)?)px`)
	convertedContent := re.ReplaceAllStringFunc(string(content), func(match string) string {
		pxValue, _ := strconv.ParseFloat(strings.TrimSuffix(match, "px"), 64)
		remValue := float64(pxValue) / pixelBase

		return fmt.Sprintf("%.2frem", remValue)
	})

	err = os.WriteFile(filename, []byte(convertedContent), filePerm)
	if err != nil {
		return err
	}

	fmt.Printf("Converted file: %s\n", filename)
	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
