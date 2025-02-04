package discovery

import "encoding/json"

type EndpointInfo struct {
	Ip       string                 `json:"ip"`
	Port     string                 `json:"port"`
	MetaData map[string]interface{} `json:"meta"`
}

func (ed *EndpointInfo) Marshal() string {
	data, err := json.Marshal(ed)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func UnmarshalEndpointInfo(data []byte) (*EndpointInfo, error) {
	ed := &EndpointInfo{}
	err := json.Unmarshal(data, ed)
	if err != nil {
		return nil, err
	}
	return ed, nil
}
