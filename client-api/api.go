package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func DefaultMux(dataEndpoint string) *http.ServeMux {
	handlers := Handlers{dataEndpoint: dataEndpoint}
	mux := http.NewServeMux()
	mux.Handle("/loot/description", http.HandlerFunc(handlers.LootData))
	return mux
}

type Handlers struct {
	dataEndpoint string
}

func (h Handlers) LootData(w http.ResponseWriter, r *http.Request) {
	// names, ok := r.URL.Query()["name"]
	// log.Printf("querying for description of %s\n", names)

	// if !ok || len(names[0]) < 1 {
	// 	http.Error(w, "name query parameter missing", http.StatusBadRequest)
	// 	return
	// }

	resp, err := http.Get(fmt.Sprintf("http://%s/", h.dataEndpoint))
	if err != nil {
		http.Error(w, fmt.Sprintf("data: %+v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "unexpected data resp code", http.StatusInternalServerError)
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("json: %+v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(bytes))
}
