package collect

import (
	"encoding/json"

	"github.com/baez90/windpark-challenge/sitesim"
)

var _ json.Marshaler = (*ParkSnapshot)(nil)

type ParkSnapshot struct {
	ID       int
	Name     string
	Turbines map[int]TurbineSnapshot
}

func (p *ParkSnapshot) Record(turbineID int, windSpeed, currentProduction float64) {
	turb := p.Turbines[turbineID]
	turb.Stats.Record(windSpeed, currentProduction)
}

func (p *ParkSnapshot) Ingest(turbines []sitesim.Turbine) {
	if p.Turbines == nil {
		p.Turbines = make(map[int]TurbineSnapshot)
	}

	for idx := range turbines {
		turbineVal := turbines[idx]
		var turb TurbineSnapshot
		if t, ok := p.Turbines[turbineVal.ID]; ok {
			turb = t
		} else {
			turb = TurbineSnapshot{
				ID:    turbineVal.ID,
				Name:  turbineVal.Name,
				Stats: Stats{},
			}
		}

		turb.Stats.Record(turbineVal.WindSpeed, turbineVal.CurrentProduction)
		p.Turbines[turbineVal.ID] = turb
	}
}

func (p ParkSnapshot) MarshalJSON() ([]byte, error) {
	tmp := struct {
		ID       int               `json:"Id"`
		Name     string            `json:"Name"`
		Turbines []TurbineSnapshot `json:"Turbines"`
	}{
		ID:       p.ID,
		Name:     p.Name,
		Turbines: make([]TurbineSnapshot, 0, len(p.Turbines)),
	}

	for _, v := range p.Turbines {
		tmp.Turbines = append(tmp.Turbines, v)
	}

	return json.Marshal(tmp)
}
