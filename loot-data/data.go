package lootdata

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DefaultMux(fullDataEnabled bool) *http.ServeMux {
	handlers := Handlers{FullDataEnabled: fullDataEnabled}
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handlers.LootData))
	return mux
}

type LootData struct {
	Name            string  `json:"name,omitempty"`
	Type            string  `json:"type,omitempty"`
	Weight          float32 `json:"weight,omitempty"`
	GridSize        string  `json:"grid_size,omitempty"`
	LootExperience  int32   `json:"loot_experience,omitempty"`
	Rarity          string  `json:"rarity,omitempty"`
	AveragePrice24H float32 `json:"average_price_24h,omitempty"`
	AveragePrice7D  float32 `json:"average_price_7d,omitempty"`
}

type Handlers struct {
	FullDataEnabled bool
}

func (h Handlers) LootData(w http.ResponseWriter, r *http.Request) {
	lootData := LootData{
		Name:           "LEDx",
		Type:           "Medical Equipment",
		Weight:         0.23,
		GridSize:       "2x1",
		LootExperience: 50,
	}
	if h.FullDataEnabled {
		lootData.Rarity = "legendary"
		lootData.AveragePrice24H = 710000.23
		lootData.AveragePrice7D = 805000.79
	}

	bytes, err := json.Marshal(&lootData)
	if err != nil {
		http.Error(w, fmt.Sprintf("json: %+v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(bytes))
	w.WriteHeader(http.StatusOK)
}
