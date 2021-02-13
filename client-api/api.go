package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Handlers struct {
	DataEndpoint string
}

func (h Handlers) LootData(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/", h.DataEndpoint), nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("data req: %+v", err), http.StatusInternalServerError)
	}
	req.Header.Set(DivertHeaderName, FromContext(r.Context()))

	resp, err := http.DefaultClient.Do(req)
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
