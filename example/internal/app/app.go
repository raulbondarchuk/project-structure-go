package app

import (
	"fmt"
	"net/http"
	"os"
	freader "prueba/pkg/freader"
	initer "prueba/pkg/initer"

	logo "github.com/jaimenetel/golibs/plogo"
	nt "github.com/jaimenetel/ngolibs/nhttp"
)

func Run() {
	// Run the application
	fmt.Println("Executing the application...")
	os.Setenv("VARTOKEN", "64ece9a47243209e7f8739bde3ff17b4ea815c777fe0a4bdfadb889db9900340")
	initer.IniciateApp() // Ejecución init() de diferentes paquetes en orden

	// server http
	startServerHTTP()

	logo.LogoLiftel() // Logo Liftel

	fmt.Println("Application running:", "http://localhost:"+SERVER_PORT)

	defer finalization()
	select {}
}

var SERVER *nt.Nthttp
var SERVER_PORT = freader.GetViperValue("server.port")

func finalization() {
	fmt.Println("Finalización de la aplicación...")
}

func startServerHTTP() {
	SERVER = nt.Ntinstance()
	SERVER.Port = SERVER_PORT
	SERVER.Start()

	go func() {
		if err := http.ListenAndServe(":"+SERVER.Port, nil); err != nil {
			fmt.Println("Error al iniciar el servidor HTTP:", err)
		}
	}()

}
