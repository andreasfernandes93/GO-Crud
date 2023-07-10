package devTools

import (
	"log"
	"net/http"
)

func StartServer(port string) {
	log.Printf("[INICIANDO SERVIDOR LOCAL]")
	log.Printf("Conexão estabelecida na porta %s", port)
	log.Printf("Acesse a URL: http://localhost:%s para visualizar o conteúdo", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
