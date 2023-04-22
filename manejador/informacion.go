package manejador

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func (m *Manejador) ObtenerInformacion() {
	rutaArchivoInfo := filepath.Join(m.RutaRaiz, "info.txt")
	archivoInfo, archivoInfoError := os.Open(rutaArchivoInfo)
	if archivoInfoError != nil {
		log.Fatalf("no se ha podido abrir el archivo info.txt %v\n", archivoInfoError)
	}
	defer archivoInfo.Close()

	escaner := bufio.NewScanner(archivoInfo)
	escaner.Split(bufio.ScanLines)

	var infoDisco InfoDisco
	contador := 0
	numeroCancion := 1
	for escaner.Scan() {
		linea := strings.TrimSpace(escaner.Text())
		contador++
		if contador == 1 {
			infoDisco.Grupo = linea
			continue
		}
		if contador == 2 {
			infoDisco.Disco = linea
		}
		if contador == 3 {
			infoDisco.Año = linea
		}
		if contador == 4 {
			if linea == "" {
				infoDisco.Portada = fmt.Sprintf("https://img.youtube.com/vi/%v/maxresdefault.jpg", m.Identificador)
				continue
			}
			infoDisco.Portada = linea
		}
		if contador > 4 {
			var cancion Cancion
			cancionInicio, cancionTitulo, cancionError := obtenerTiempoTitulo(linea)
			if cancionError != nil {
				log.Fatalf("error procesando la información de la linea %d: %v -> %v\n", contador, linea, cancionError)
			}
			cancion.Inicio = cancionInicio
			cancion.Titulo = cancionTitulo
			cancion.Numero = fmt.Sprintf("%02d", numeroCancion)
			cancion.Track = numeroCancion
			infoDisco.Canciones = append(infoDisco.Canciones, cancion)
			numeroCancion++
		}
	}

	nombreCompleto := fmt.Sprintf("%v %v %v", infoDisco.Grupo, infoDisco.Año, infoDisco.Disco)
	infoDisco.NombreNormalizado = normalizarTexto(nombreCompleto)

	m.InfoDisco = infoDisco
}
