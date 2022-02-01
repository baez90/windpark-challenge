package collect

import "encoding/json"

var _ json.Marshaler = (*StatValue)(nil)

type StatValue struct {
	sum   float64
	count int
}

func (s StatValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.sum / float64(s.count))
}

func (s *StatValue) Record(val float64) {
	s.sum += val
	s.count++
}

type Stats struct {
	WindSpeed         StatValue
	CurrentProduction StatValue
}

func (s *Stats) Record(windSpeed, currentProduction float64) {
	s.WindSpeed.Record(windSpeed)
	s.CurrentProduction.Record(currentProduction)
}

type TurbineSnapshot struct {
	ID    int    `json:"Id"`
	Name  string `json:"Name"`
	Stats Stats  `json:"Stats"`
}
