package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func DefaultMux(dataEndpoint string) *http.ServeMux {
	handlers := Handlers{dataEndpoint: dataEndpoint}
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handlers.LootData))
	return mux
}

type Handlers struct {
	dataEndpoint string
}

func (h Handlers) LootData(w http.ResponseWriter, r *http.Request) {
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
