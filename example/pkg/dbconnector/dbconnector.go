package dbconnector

import (
	"fmt"

	dblib "github.com/raulbondarchuk/dbconnector-go"
	"gorm.io/gorm"
)

var dbManager DBManager

// Getter de la lista de los conexiones bdd
func GetDBConnector() *DBManager { return &dbManager }

type DBManager struct {
	DBName *gorm.DB
}

// SetDBConnector establece la conexión con las bases de datos
func SetDBConnector() {
	registry := dblib.GetInstanceMlt()

	// Configuración de la conexión con los bdds
	registry.AddDBManager("database")

	// Añadimos los bdds a nuestra lista
	dbManager = DBManager{
		DBName: dblib.GetInstanceMlt().GetDBManager("database").GetDB(),
	}

	fmt.Println("**** **** La conexión con base de datos establecida correctamente **** ****")
}
