package manejador

import (
	"log"
	"os"
	"path/filepath"
)

func esIdentificadorValido(identificador string) bool {
	return len([]rune(identificador)) == 11
}

func CrearManejador(args []string, ffmpeg, youtube *string) Manejador {
	if len(args) == 1 {
		log.Fatalln("no se ha proporcionado identificador de Youtube")
	}
	idValido := esIdentificadorValido(args[1])

	if !idValido {
		log.Fatalln("el identificador proporcionado de Youtube no parece válido")
	}

	ejecutable, ejecutableError := os.Executable()
	if ejecutableError != nil {
		log.Fatalf("no se ha podido recuperar la ruta del ejecutable %v\n", ejecutableError)
	}
	rutaRaiz := filepath.Dir(ejecutable)
	rutaFfmpeg := filepath.Join(rutaRaiz, "herramientas", *ffmpeg)
	rutaYoutube := filepath.Join(rutaRaiz, "herramientas", *youtube)
	rutaCarpetaTemporal := filepath.Join(rutaRaiz, "temp")
	rutaAudio := filepath.Join(rutaCarpetaTemporal, "youtube_audio.mp3")

	return Manejador{
		Identificador:       args[1],
		RutaRaiz:            rutaRaiz,
		RutaCarpetaTemporal: rutaCarpetaTemporal,
		RutaAudio:           rutaAudio,
		RutaFfmpeg:          rutaFfmpeg,
		RutaYoutube:         rutaYoutube,
	}
}
