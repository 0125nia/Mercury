package discovery

import "encoding/json"

type Endpoint struct {
	Ip       string                 `json:"ip"`
	Port     string                 `json:"port"`
	MetaData map[string]interface{} `json:"meta"`
}

func (ed *Endpoint) Marshal() string {
	data, err := json.Marshal(ed)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func UnmarshalEndpoint(data []byte) (*Endpoint, error) {
	ed := &Endpoint{}
	err := json.Unmarshal(data, ed)
	if err != nil {
		return nil, err
	}
	return ed, nil
}
