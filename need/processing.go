package need

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Out_ndjson(files []string, path string) {

	findslash := strings.FieldsFunc(path, func(r rune) bool { //funcion de separacion
		return r == '/'
	}) //generacion de slice para el path
	var PathName string

	PathName = findslash[len(findslash)-1]
	PathName = strings.TrimPrefix(PathName, "_") //elimina el "_"

	//var cont int64 = 0
	for _, name_file := range files {
		filepath.Join(path, name_file)
		fmt.Println()

	}

	return
}

// Separa las cadenas
func section() {

}

func WavingZinc() {
	file, err := os.ReadFile(os.Getenv("Dataset_") + ".ndjson")
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:4080/api/_bulk", bytes.NewBuffer(file))
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth("admin", "Complexpass#123")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
