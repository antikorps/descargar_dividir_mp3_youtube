package main

import (
	"descargar_dividir_mp3_youtube/manejador"
	"flag"
	"log"
	"os"
)

func main() {
	argumentos := os.Args

	ffmpeg := flag.String("ffmpeg", "ffmpeg", "nombre del archivo ffmpeg en la carpeta herramientas")
	youtube := flag.String("youtube", "yt-dlp_linux", "nombre del archivo youtubedr en la carpeta herramientas")

	flag.Parse()

	m := manejador.CrearManejador(argumentos, ffmpeg, youtube)
	m.ObtenerInformacion()
	m.Descargar()
	m.Dividir()
	m.Mover()

	log.Println("FIN: ejecución correcta")
	os.Exit(0)
}
