package manejador

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func (m *Manejador) Descargar() {
	os.RemoveAll(m.RutaCarpetaTemporal)

	carpetaTemporalError := os.Mkdir(m.RutaCarpetaTemporal, 0777)
	if carpetaTemporalError != nil {
		log.Fatalf("error creando la carpeta temp necesaria para los archivos temporales: %v\n", carpetaTemporalError)
	}

	comando := exec.Command(m.RutaYoutube, "-f", "ba", "-x", "--audio-format", "mp3", fmt.Sprintf("https://www.youtube.com/watch?v=%v", m.Identificador), "-o", m.RutaAudio)

	comandoError := comando.Run()
	if comandoError != nil {
		log.Fatalf("error descargando el vídeo: %v\n", comandoError)
	}

	// Descargar portada
	cliente := http.Client{}
	pet, petError := http.NewRequest("GET", m.InfoDisco.Portada, nil)
	if petError != nil {
		log.Printf("no se ha podido descargar la portada, pero continuará el proceso de descarga: %v\n", petError)
		return
	}
	resp, respError := cliente.Do(pet)
	if respError != nil {
		log.Printf("no se ha podido descargar la portada, pero continuará el proceso de descarga: %v\n", respError)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		// Intentar alternativa
		urlAlternativa := fmt.Sprintf("https://i.ytimg.com/vi/%v/hqdefault.jpg", m.Identificador)
		pet2, pet2Error := http.NewRequest("GET", urlAlternativa, nil)
		if pet2Error != nil {
			log.Printf("no se ha podido descargar la portada, pero continuará el proceso de descarga: status code incorrecto -> %v. Prueba a descargar manualmente esta alternativa: %v\n", resp.Status, urlAlternativa)
			return
		}
		resp2, resp2Error := cliente.Do(pet2)
		if resp2Error != nil {
			log.Printf("no se ha podido descargar la portada, pero continuará el proceso de descarga: status code incorrecto -> %v. Prueba a descargar manualmente esta alternativa: %v\n", resp.Status, urlAlternativa)
			return
		}
		defer resp2.Body.Close()
		if resp2.StatusCode != 200 {
			log.Printf("no se ha podido descargar la portada, pero continuará el proceso de descarga: status code incorrecto -> %v, status code alternativa: %v\n", resp.Status, resp2.Status)
			return
		}
		resp = resp2

	}

	var nombrePortada string

	for clave, valor := range resp.Header {
		if strings.Contains(clave, "ontent") && strings.Contains(clave, "ype") && !strings.Contains(clave, "tions") {
			infoContentType := strings.Split(valor[0], "/")
			if len(infoContentType) != 2 {
				log.Printf("no se ha podido descargar la portada, pero continuará el proceso de descarga: Content-Type no procesable -> %v", valor[0])
				return
			}
			if infoContentType[0] != "image" {
				log.Printf("no se ha podido descargar la portada, pero continuará el proceso de descarga: el Content-Type no corresponde a image -> %v", valor[0])
				return
			}
			nombrePortada = fmt.Sprintf("cover.%v", infoContentType[1])
		}
	}
	m.PortadaNombre = nombrePortada
	rutaPortada := filepath.Join(m.RutaCarpetaTemporal, nombrePortada)

	archivoPortada, archivoPortadaError := os.Create(rutaPortada)
	if archivoPortadaError != nil {
		log.Printf("no se ha podido descargar la portada, pero continuará el proceso de descarga: error creando el archivo para la portada -> %v", archivoPortadaError)
		return
	}
	defer archivoPortada.Close()

	_, copiaError := io.Copy(archivoPortada, resp.Body)
	if copiaError != nil {
		log.Printf("no se ha podido descargar la portada, pero continuará el proceso de descarga: error copiando portada -> %v", copiaError)
	}

	m.PortadaDescargada = true
}
