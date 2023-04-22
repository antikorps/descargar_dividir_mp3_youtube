package manejador

import (
	"log"
	"os"
	"path/filepath"
)

func (m *Manejador) Mover() {
	rutaDiscos := filepath.Join(m.RutaRaiz, "discos")

	rutaDiscosExiste, rutaDiscosExisteError := existe(rutaDiscos)
	if rutaDiscosExisteError != nil {
		log.Fatalf("error comprobando la carpeta discos antes de mover el disco: %v\n", rutaDiscosExisteError)
	}
	if !rutaDiscosExiste {
		crearCarpetaDiscosError := os.Mkdir(rutaDiscos, 0777)
		if crearCarpetaDiscosError != nil {
			log.Fatalf("error creando la carpeta discos antes de mover el disco: %v\n", crearCarpetaDiscosError)
		}
	}

	rutaTempDisco := filepath.Join(m.RutaCarpetaTemporal, m.InfoDisco.NombreNormalizado)
	rutaDiscosDisco := filepath.Join(rutaDiscos, m.InfoDisco.NombreNormalizado)

	moverDiscoError := os.Rename(rutaTempDisco, rutaDiscosDisco)
	if moverDiscoError != nil {
		log.Fatalf("error moviendo la carpeta del directorio temporal al definitivo: %v\n", moverDiscoError)
	}

	borrarTempError := os.RemoveAll(m.RutaCarpetaTemporal)
	if borrarTempError != nil {
		log.Printf("error borrando la carpeta temporal: %v\n", borrarTempError)
	}

}
