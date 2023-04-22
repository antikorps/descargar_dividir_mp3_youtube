package manejador

import (
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func existe(ruta string) (bool, error) {
	_, err := os.Stat(ruta)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func escaparMetadata(s string) string {
	comillasDobles := regexp.MustCompile(`"`)
	s = comillasDobles.ReplaceAllString(s, `\"`)
	return s
}

func normalizarTexto(texto string) string {
	transformado := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	textoTransformado, _, textoTransformadoError := transform.String(transformado, texto)
	if textoTransformadoError != nil {
		log.Fatalf("error normalizando %v -> %v\n", texto, textoTransformadoError)
	}
	// Minúsculas
	textoTransformado = strings.ToLower(textoTransformado)
	// Caracteres no deseados
	expRegNoDeseado := regexp.MustCompile(`[-!$%^&*()_+|~={}\[\]:";'<>?,.\/\s]`)
	textoTransformado = expRegNoDeseado.ReplaceAllString(textoTransformado, "_")
	// Corregir guion bajo
	expRegCorregirGuion := regexp.MustCompile(`_{2,}`)
	textoTransformado = expRegCorregirGuion.ReplaceAllString(textoTransformado, "_")
	// Sustituir dobles guiones
	expEspacios := regexp.MustCompile(`_+`)
	textoTransformado = expEspacios.ReplaceAllString(textoTransformado, "_")

	return textoTransformado
}

func hhmmssASegundos(lineaHoras, lineaMinutos, lineaSegundos string) (int, error) {
	var totalSegundos int

	horas, horasError := strconv.Atoi(lineaHoras)
	if horasError != nil {
		return totalSegundos, errors.New("no se ha podido obtener las horas numéricas del formato hh:mm:ss")
	}
	totalSegundos += horas * 3600

	minutos, minutosError := strconv.Atoi(lineaMinutos)
	if minutosError != nil {
		return totalSegundos, errors.New("no se ha podido obtener los minutos numéricos del formato hh:mm:ss")
	}
	totalSegundos += minutos * 60

	segundos, segundosError := strconv.Atoi(lineaSegundos)
	if segundosError != nil {
		return totalSegundos, errors.New("no se ha podido obtener los segundos numéricos del formato hh:mm:ss")
	}
	totalSegundos += segundos

	return totalSegundos, nil
}

func obtenerTiempoTitulo(linea string) (int, string, error) {
	expTiempoTitulo := regexp.MustCompile(`([0-9]{2}):([0-9]{2}):([0-9]{2})\s(.*)`)
	gruposCapturados := expTiempoTitulo.FindStringSubmatch(linea)
	if len(gruposCapturados) != 5 {
		return 0, "", errors.New("no se ha detectado un formato hh:mm:ss título válido")
	}
	lineaHoras, lineaMinutos, lineaSegundos, lineaTitulo := gruposCapturados[1], gruposCapturados[2], gruposCapturados[3], gruposCapturados[4]

	segundosCancion, segundosCancionError := hhmmssASegundos(lineaHoras, lineaMinutos, lineaSegundos)
	if segundosCancionError != nil {
		log.Fatalf(segundosCancionError.Error())
	}

	return segundosCancion, lineaTitulo, nil
}
