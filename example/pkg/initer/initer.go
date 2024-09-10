package initer

import (
	db "prueba/pkg/dbconnector"
	freader "prueba/pkg/freader"

	initer "github.com/raulbondarchuk/initer"
)

func IniciateApp() {
	initer.RegisterInitFunc(db.SetDBConnector, 998)
	initer.RegisterInitFunc(freader.IniciateFReader, 999)

	// Ejecución init() de diferentes paquetes en orden
	initer.Initialize()
}
