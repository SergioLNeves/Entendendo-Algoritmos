package Recursividade

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListarPasta(path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			ListarPasta(fullPath)
		} else {
			fmt.Println(fullPath)
		}
	}
}
