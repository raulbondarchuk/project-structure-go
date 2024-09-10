package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

func main() {
	// Salida del directorio de trabajo actual
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	fmt.Println("Current working directory:", wd)

	// Configuración Viper para leer build.toml
	viper.SetConfigName("builder")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./pkg/builder")

	// Log de la ruta del archivo de configuración
	fmt.Println("Looking for config file in:", filepath.Join(wd, "configs"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	outputDir := viper.GetString("build.output_dir")
	linuxOutput := viper.GetString("build.linux_output")
	windowsOutput := viper.GetString("build.windows_output")
	sourceFile := viper.GetString("build.source_file")

	// Crear directorio de salida con timestamp
	timestamp := time.Now().Format("2006-01-02-15-04-05")
	outputDir = filepath.Join(outputDir, "builds", "build-"+timestamp)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Error creating build directory: %v", err)
	}

	// build Linux
	linuxCmd := exec.Command("go", "build", "-o", filepath.Join(outputDir, linuxOutput), sourceFile)
	linuxCmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=amd64")
	if output, err := linuxCmd.CombinedOutput(); err != nil {
		log.Fatalf("Error building for Linux: %v\nOutput: %s", err, string(output))
	} else {
		fmt.Println("Successfully built for Linux:", filepath.Join(outputDir, linuxOutput))
	}

	// build Windows
	windowsCmd := exec.Command("go", "build", "-o", filepath.Join(outputDir, windowsOutput), sourceFile)
	windowsCmd.Env = append(os.Environ(), "GOOS=windows", "GOARCH=amd64")
	if output, err := windowsCmd.CombinedOutput(); err != nil {
		log.Fatalf("Error building for Windows: %v\nOutput: %s", err, string(output))
	} else {
		fmt.Println("Successfully built for Windows:", filepath.Join(outputDir, windowsOutput))
	}

	// Copiar archivo de configuración a la carpeta de salida
	configFilePath := filepath.Join(wd, "configs", "config.toml")
	destConfigFilePath := filepath.Join(outputDir, "config.toml")
	if err := copyFile(configFilePath, destConfigFilePath); err != nil {
		log.Fatalf("Error copying config file: %v", err)
	} else {
		fmt.Println("Successfully copied config file to:", destConfigFilePath)
	}
}

// copyFile sirve para copiar de src en dst
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, sourceFile); err != nil {
		return err
	}

	return nil
}
