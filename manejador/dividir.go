package manejador

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func (m *Manejador) Dividir() {
	// Duración del audio
	var duracionTotal int
	comando := exec.Command(m.RutaFfmpeg, "-i", m.RutaAudio)
	salida, _ := comando.CombinedOutput()

	escaner := bufio.NewScanner(bytes.NewReader(salida))
	escaner.Split(bufio.ScanLines)
	for escaner.Scan() {
		linea := strings.TrimSpace(escaner.Text())
		if strings.HasPrefix(linea, "Duration") {
			expRegHorasMinutosSegundos := regexp.MustCompile(`Duration:.([0-9]{0,2}):([0-9]{0,2}):([0-9]{0,2})`)
			gruposCapturados := expRegHorasMinutosSegundos.FindStringSubmatch(linea)
			if len(gruposCapturados) != 4 {
				log.Fatalf("no se ha detectado un formato de duración total válido: %v\n", strings.Join(gruposCapturados, "-->"))
			}
			calculoDuracion, duracionTotalError := hhmmssASegundos(gruposCapturados[1], gruposCapturados[2], gruposCapturados[3])
			if duracionTotalError != nil {
				log.Fatalln(duracionTotalError)
			}
			duracionTotal = calculoDuracion
		}
	}

	rutaCarpetaDisco := filepath.Join(m.RutaCarpetaTemporal, m.InfoDisco.NombreNormalizado)
	os.RemoveAll(rutaCarpetaDisco)
	errorCarpetaDisco := os.Mkdir(rutaCarpetaDisco, 0777)
	if errorCarpetaDisco != nil {
		log.Fatalf("error al crear al carpeta con el nombre del disco normalizado: %v\n", errorCarpetaDisco)
	}

	for indice, cancion := range m.InfoDisco.Canciones {
		var final int

		if indice == len(m.InfoDisco.Canciones)-1 {
			final = duracionTotal
		} else {
			final = m.InfoDisco.Canciones[indice+1].Inicio
		}

		nombreMp3 := fmt.Sprintf("%v_%v_%v.mp3", m.InfoDisco.NombreNormalizado, cancion.Numero, normalizarTexto(cancion.Titulo))

		rutaMp3 := filepath.Join(rutaCarpetaDisco, nombreMp3)
		// ffmpeg -i ENTRADA -acodec copy -ss INICIO -to FIN SALIDA

		metadataTitulo := fmt.Sprintf(`title="%v"`, escaparMetadata(cancion.Titulo))
		metadataArtist := fmt.Sprintf(`artist="%v"`, escaparMetadata(m.InfoDisco.Grupo))
		metadataAlbum := fmt.Sprintf(`album="%v"`, escaparMetadata(m.InfoDisco.Disco))
		metadataAño := fmt.Sprintf(`year="%v"`, escaparMetadata(m.InfoDisco.Año))
		metadataNumeroPista := fmt.Sprintf(`track="%v"`, escaparMetadata(fmt.Sprint(cancion.Track)))

		comando := exec.Command(m.RutaFfmpeg, "-i", m.RutaAudio, "-metadata", metadataTitulo, "-metadata", metadataArtist, "-metadata", metadataAlbum, "-metadata", metadataAño, "-metadata", metadataNumeroPista, "-acodec", "copy", "-ss", fmt.Sprint(cancion.Inicio), "-to", fmt.Sprint(final), rutaMp3)
		errorComando := comando.Run()
		if errorComando != nil {
			log.Fatalf("error intentando separar la canciones %v -> %v\n", cancion.Numero, errorComando)
		}
	}

	if m.PortadaDescargada {
		rutaPortada := filepath.Join(m.RutaCarpetaTemporal, m.PortadaNombre)
		rutaPortadaNueva := filepath.Join(rutaCarpetaDisco, m.PortadaNombre)
		errorPortada := os.Rename(rutaPortada, rutaPortadaNueva)
		if errorPortada != nil {
			log.Printf("error al mover la portada, se tendrá que descargar y hacer manualmente: %v", errorPortada)
		}
	}
}
