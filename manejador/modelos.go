package manejador

type Manejador struct {
	Identificador       string
	RutaRaiz            string
	RutaCarpetaTemporal string
	RutaAudio           string
	RutaYoutube         string
	RutaFfmpeg          string
	InfoDisco           InfoDisco
	PortadaDescargada   bool
	PortadaNombre       string
}

type InfoDisco struct {
	Disco             string
	Grupo             string
	Año               string
	Portada           string
	NombreNormalizado string
	Canciones         []Cancion
}
type Cancion struct {
	Inicio            int
	Titulo            string
	Numero            string
	Track             int
	NombreNormalizado string
}
