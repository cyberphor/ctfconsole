package serve

import (
	"net/http"
)

func Console() {
	filePath := http.Dir("../templates/")
	fileServer := http.FileServer(filePath)
	http.Handle("/", fileServer)
	http.ListenAndServe(":8000", nil)
}
