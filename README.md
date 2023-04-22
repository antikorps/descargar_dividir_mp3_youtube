# Descargar y dividir mp3 de YouTube
## Descripción
Este script está ideado para descargar discos de YouTube subidos como un único vídeo.\
Busca automatizar el proceso de descarga del vídeo, extracción del mp3 y división en canciones.\
Para la extracción de las canciones se espera que exista el archivo **info.txt** en el mismo directorio que el ejecutable.\
Este archivo debe tener la siguiente estructurada y siguiente marcado:
```text
NOMBRE DEL GRUPO
TÍTULO DEL DISCO
AÑO DE PUBLICACIÓN
PORTADA
00:00:00 Título de la canción 1
00:03:20 Título de la canción 2
00:06:40 Título de la canción 3
```
- Se debe respeta el orden de la línea y el marcado temporal.
- La portada debe ser una URL, si se deja en blanco utilizará el poster de YouTube.

## Uso GNU/Linux (amd64)
### Preparación
Descargar el archivo comprimido de la carpeta bin. Junto al archivo **descargar_dividir_mp3_youtube** aparecerá la carpeta **herramientas** con los 2 ejecutables utilizados en el script: **ffmpeg** (utilizado para dividir e incorporar los metadatos) y **yt-dlp_linux** para descargar de YouTube. Esta estructura de archivos y carpetas debe mantenerse siempre. Asegurar que todos los archivos disponen de permisos de ejecución:
```bash
chmod +x descargar_dividir_mp3_youtube && chmod +x herramientas/ffmpeg && chmod +x yt-dlp_linux
```
### Ejecución
Simplemente debe lanzarse el ejecutable con el identificador de YouTube como argumento. Este identificador es un código de 11 caracteres:
```bash
./descargar_dividir_mp3_youtube TnSAT3OVoKg
```
### Resultados
Mientras se está ejecutando se creará una carpeta llamada **temp** donde se irán almacenando los archivos temporales. Una vez acabado el script se moverán la carpeta resultante al directorio **discos**

### Uso avanzando
El script admite 2 flags -ffmpeg y -youtube. Estos flags pueden usarse para pasar la ruta de ambos ejecutables. Puede ser útil si ya si disponen de los mismos. O incluso para usarse en Windows.

## Uso otros sistemas operativos o arquitecturas
No he preparado versiones para otros sistemas operativos o arquitecturas, aunque tendría que poderse realizar fácilmente. Simplemente bastaría con compilar el programa y buscar los ejecutables correspondientes al sistema operativo/arquitectura en el que se van a usar. Estos ejecutables pueden pasarse a través de los flags indicados en el punto anterior (uso avanzando).


