package helpers

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

func RunCopyBaseEntity(cmd *cobra.Command, args []string) {
	// Step 1: Dynamically resolve the base schema directory
	baseDir := getBaseDir()
	baseSchemaDir := filepath.Join(baseDir, "ent", "schema")

	// Step 2: Dynamically determine the target project schema directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		return
	}
	projectSchemaDir := filepath.Join(currentDir, "ent", "schema")

	// Step 3: List all available entities in the base schema directory
	entities, err := listEntities(baseSchemaDir)
	if err != nil {
		fmt.Printf("Error listing base schemas: %v\n", err)
		return
	}

	fmt.Println("Available entities to copy:")
	for _, entity := range entities {
		fmt.Printf(" - %s\n", entity)
	}

	// Step 4: Prompt user input
	fmt.Print("Enter the entities to copy (comma-separated): ")
	var input string
	fmt.Scanln(&input)
	selectedEntities := strings.Split(input, ",")

	// Step 5: Check and copy entities
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

func getBaseDir() string {
	// Use the directory of this source file during development
	_, sourceFile, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filepath.Dir(sourceFile))
	return baseDir
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

func getModuleName() (string, error) {
	// Step 2: Dynamically determine the target project schema directory
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current working directory: %v", err)
	}
	goFile := filepath.Join(currentDir, "go.mod")

	file, err := os.Open(goFile)
	if err != nil {
		return "", fmt.Errorf("error opening go.mod: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading go.mod: %v\n", err)
	}

	return "", fmt.Errorf("module name not found in go.mod")
}

func getBaseModuleName() (string, error) {
	baseDir := getBaseDir()
	baseGoFile := filepath.Join(baseDir, "go.mod")

	file, err := os.Open(baseGoFile)
	if err != nil {
		return "", fmt.Errorf("error opening go.mod: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading go.mod: %v\n", err)
	}

	return "", fmt.Errorf("module name not found in go.mod")
}
