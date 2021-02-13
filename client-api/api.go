package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handlers struct {
	DataEndpoint   string
	NestFleaMarket bool
}

type LootData struct {
	ItemData
	FleaMarketData
}

type ItemData struct {
	Name           string  `json:"name,omitempty"`
	Type           string  `json:"type,omitempty"`
	Weight         float32 `json:"weight,omitempty"`
	GridSize       string  `json:"grid_size,omitempty"`
	LootExperience int32   `json:"loot_experience,omitempty"`
}

type FleaMarketData struct {
	Rarity          string  `json:"rarity,omitempty"`
	AveragePrice24H float32 `json:"average_price_24h,omitempty"`
	AveragePrice7D  float32 `json:"average_price_7d,omitempty"`
}

type NestedLootData struct {
	Item ItemData       `json:"item"`
	Flea FleaMarketData `json:"flea"`
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

	var lootData LootData
	err = json.NewDecoder(resp.Body).Decode(&lootData)
	if err != nil {
		http.Error(w, fmt.Sprintf("bad loot json %+v", err), http.StatusInternalServerError)
		return
	}

	if h.NestFleaMarket {
		nested := NestedLootData{
			Item: ItemData{
				Name:           lootData.Name,
				Type:           lootData.Type,
				Weight:         lootData.Weight,
				GridSize:       lootData.GridSize,
				LootExperience: lootData.LootExperience,
			},
			Flea: FleaMarketData{
				Rarity:          lootData.Rarity,
				AveragePrice24H: lootData.AveragePrice24H,
				AveragePrice7D:  lootData.AveragePrice7D,
			},
		}
		bytes, err := json.MarshalIndent(&nested, "", "\t")
		if err != nil {
			http.Error(w, fmt.Sprintf("resp json: %+v", err), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(bytes))
		return
	}

	err = json.NewEncoder(w).Encode(lootData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
