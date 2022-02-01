package collect

import (
	"encoding/json"
	"time"

	"github.com/baez90/windpark-challenge/sitesim"
)

var _ json.Marshaler = (*ParksSnapshot)(nil)

type ParksSnapshot struct {
	Timestamp time.Time
	Parks     map[int]ParkSnapshot
}

func (ps *ParksSnapshot) Ingest(p []sitesim.Site) {
	if ps.Parks == nil {
		ps.Parks = make(map[int]ParkSnapshot)
	}

	for idx := range p {
		parkVal := p[idx]
		var parkSnap ParkSnapshot
		if p, ok := ps.Parks[parkVal.ID]; ok {
			parkSnap = p
		} else {
			parkSnap = ParkSnapshot{
				ID:   parkVal.ID,
				Name: parkVal.Name,
			}
		}

		parkSnap.Ingest(parkVal.Turbines)
		ps.Parks[parkVal.ID] = parkSnap
	}
}

func (ps ParksSnapshot) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Timestamp time.Time      `json:"Timestamp"`
		Parks     []ParkSnapshot `json:"Parks"`
	}{
		Timestamp: ps.Timestamp,
		Parks:     make([]ParkSnapshot, 0, len(ps.Parks)),
	}

	for _, v := range ps.Parks {
		tmp.Parks = append(tmp.Parks, v)
	}

	return json.Marshal(tmp)
}
