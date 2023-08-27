package main

import (
	"back/need"
	"fmt"
	"os"
	"path/filepath"
)

type FileList struct {
	Files []string
	Dirs  []string
}

func main() {

	//Se busca el path de la base de datos
	databaseName := os.Args[1]
	Path, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	// path de la base de datos
	databasePath := filepath.Join(Path, databaseName)
	os.Setenv("Dataset_", databaseName)    // Dataset_ se vuelve variable de entorno con el valor de databaseName
	files, err := os.ReadDir(databasePath) // lee el contenido de un directorio en el sistema de archivos y devuelve una lista de objetos en files
	if err != nil {
		panic(err)
	}
	fileList := FileList{} // Crear una instancia de la estructura FileList

	for _, file := range files {
		if file.IsDir() {
			fileList.Dirs = append(fileList.Dirs, file.Name())
		} else {
			fileList.Files = append(fileList.Files, file.Name())
		}
	}
	fmt.Println(fileList.Dirs)
	///FALTA PROBAR
	if len(fileList.Files) >= 1 {
		need.Out_ndjson(fileList.Files, Path)
	}

	/*for _, dir := range fileList.Dirs {
		need.Browser_dirs(dir, Path)
	}*/

	// Aquí puedes continuar utilizando las listas fileList.Dirs y fileList.Files según tus necesidades
}
