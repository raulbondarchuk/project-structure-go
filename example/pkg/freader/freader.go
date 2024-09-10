package freader

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	fileConfigName string = "config"    // nombre del archivo de configuración sin la extensión
	fileConfigPath string = "./config/" // o cualquier otro directorio donde esté el archivo
	fileConfigType string = "toml"      // o "toml", "yaml", etc.

	nameWithoutExtension string
)

func GetViperValue(clave string) string {
	if nameWithoutExtension == "" {
		IniciateFReader()
	}
	return viper.GetString(clave)
}

func IniciateFReader() {
	_, errdir := os.Getwd()
	if errdir != nil {
		fmt.Println("Error al obtener el directorio de trabajo:", errdir)
		return
	}
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println("Error al obtener el ejecutable:", err)
		return
	}
	execName := filepath.Base(execPath)
	filename := execName
	extension := filepath.Ext(filename)
	nameWithoutExtension = filename[:len(filename)-len(extension)]

	viper.SetConfigName(fileConfigName)
	viper.SetConfigType(fileConfigType)
	viper.AddConfigPath(fileConfigPath)
	errvip := viper.ReadInConfig()
	if errvip != nil {
		log.Fatalf("Error al leer el archivo de configuración: %v", errvip)
	}
}

func GetExecName() string {
	return nameWithoutExtension
}
