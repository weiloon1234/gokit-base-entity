package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func RunCopyBaseEntityHook(cmd *cobra.Command, args []string) {
	// Step 1: Dynamically resolve the base hook directory
	hookDir := getHookDir()
	baseHookDir := filepath.Join(hookDir, "ent", "hook")

	// Step 2: Dynamically determine the target project hook directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		return
	}
	projectHookDir := filepath.Join(currentDir, "ent", "hook")

	// Step 3: List all available entities in the base schema directory
	hooks, err := listHooks(baseHookDir)
	if err != nil {
		fmt.Printf("Error listing base schemas: %v\n", err)
		return
	}

	fmt.Println("Available entities to copy:")
	for _, hook := range hooks {
		fmt.Printf(" - %s\n", hook)
	}

	// Step 4: Prompt user input
	fmt.Print("Enter the hooks to copy (comma-separated): ")
	var input string
	fmt.Scanln(&input)
	selectedHooks := strings.Split(input, ",")

	var successHooks []string
	// Step 5: Check and copy entities
	for _, hook := range selectedHooks {
		hook = strings.TrimSpace(hook)
		if hook == "" {
			continue
		}

		baseFile := filepath.Join(baseHookDir, hook+".go")
		targetFile := filepath.Join(projectHookDir, hook+".go")

		// Check if the entity already exists in the project
		if fileExists(targetFile) {
			fmt.Printf("Hook %s already exists in the project. Skipping.\n", hook)
			continue
		}

		fromModuleName, ok1 := getBaseModuleName()
		toModuleName, ok2 := getModuleName()

		if ok1 != nil || ok2 != nil {
			fmt.Printf("Error getting module name: %v\n", err)
			continue
		}

		// Replace the project name in the base file
		baseContent, err := os.ReadFile(baseFile)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", hook, err)
			continue
		}

		if !fileExists(targetFile) {
			// Replace the project name in the file content
			updatedContent := []byte(strings.ReplaceAll(string(baseContent), fromModuleName, toModuleName))

			if err := os.WriteFile(baseFile, updatedContent, 0644); err != nil {
				fmt.Printf("Error writing %s: %v\n", hook, err)
				continue
			}

			fmt.Printf("Successfully copied %s to the project.\n", hook)
		} else {
			fmt.Printf("Hook %s already exists in the project. Skipping.\n", hook)
			continue
		}
	}

	if len(successHooks) > 0 {
		fmt.Printf("======================\n")
		for _, hook := range successHooks {
			// Generate a hint message based on the hook file name
			hookFuncName := cases.Title(language.Und).String(strings.ReplaceAll(hook, "_", ""))
			fmt.Printf(" - %s\n", hook)
			fmt.Printf("Hint: Remember to register the hook in your main.go file like this:\n")
			fmt.Printf("      dbClient.Use(hook.%s)\n", hookFuncName)
		}
		fmt.Printf("======================\n")
		fmt.Printf("Remember to rebuild entity after copying hooks:\n")
	} else {
		fmt.Printf("======================\n")
		fmt.Printf("No hooks copied.\n")
		fmt.Printf("======================\n")
	}
}

func getHookDir() string {
	// Use the directory of this source file during development
	_, sourceFile, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filepath.Dir(sourceFile))
	return baseDir
}

func listHooks(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var hooks []string
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".go") || strings.HasPrefix(file.Name(), "hook.go") {
			continue
		}
		hook := strings.TrimSuffix(file.Name(), ".go")
		hooks = append(hooks, hook)
	}
	return hooks, nil
}
