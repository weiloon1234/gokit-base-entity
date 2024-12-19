package helpers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func RunCopyBaseEntity(cmd *cobra.Command, args []string) {
	// Dynamically resolve the base schema directory
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Error resolving executable path: %v\n", err)
		return
	}
	baseDir := filepath.Dir(filepath.Dir(executablePath))    // Resolve to package root
	baseSchemaDir := filepath.Join(baseDir, "ent", "schema") // Base schemas directory

	// Dynamically determine the target project schema directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		return
	}
	projectSchemaDir := filepath.Join(currentDir, "ent", "schema") // Project schema directory

	// Step 1: List all available entities in the base schema directory
	entities, err := listEntities(baseSchemaDir)
	if err != nil {
		fmt.Printf("Error listing base schemas: %v\n", err)
		return
	}

	fmt.Println("Available entities to copy:")
	for _, entity := range entities {
		fmt.Printf(" - %s\n", entity)
	}

	// Step 2: Prompt user input
	fmt.Print("Enter the entities to copy (comma-separated): ")
	var input string
	fmt.Scanln(&input)
	selectedEntities := strings.Split(input, ",")

	// Step 3: Check and copy entities
	for _, entity := range selectedEntities {
		entity = strings.TrimSpace(entity)
		if entity == "" {
			continue
		}

		baseFile := filepath.Join(baseSchemaDir, entity+".go")
		targetFile := filepath.Join(projectSchemaDir, entity+".go")

		// Check if the entity already exists in the project
		if fileExists(targetFile) {
			fmt.Printf("Entity %s already exists in the project. Skipping.\n", entity)
			continue
		}

		// Copy the schema file
		if err := copyFile(baseFile, targetFile); err != nil {
			fmt.Printf("Error copying %s: %v\n", entity, err)
		} else {
			fmt.Printf("Successfully copied %s to the project.\n", entity)
		}
	}
}

func listEntities(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var entities []string
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".go") {
			continue
		}
		entity := strings.TrimSuffix(file.Name(), ".go")
		entities = append(entities, entity)
	}
	return entities, nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}
