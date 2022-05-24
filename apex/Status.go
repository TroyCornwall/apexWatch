package apex

type Status struct {
	Istat struct {
		Hostname string `json:"hostname"`
		Software string `json:"software"`
		Hardware string `json:"hardware"`
		Serial   string `json:"serial"`
		Type     string `json:"type"`
		Extra    struct {
			Sdver string `json:"sdver"`
		} `json:"extra"`
		Timezone string `json:"timezone"`
		Date     int    `json:"date"`
		Feed     struct {
			Name   int `json:"name"`
			Active int `json:"active"`
		} `json:"feed"`
		Power struct {
			Failed   int `json:"failed"`
			Restored int `json:"restored"`
		} `json:"power"`
		Inputs  []Input `json:"inputs"`
		Outputs []struct {
			Status []string `json:"status"`
			Name   string   `json:"name"`
			Gid    string   `json:"gid"`
			Type   string   `json:"type"`
			ID     int      `json:"ID"`
			Did    string   `json:"did"`
		} `json:"outputs"`
		Link struct {
			LinkState int    `json:"linkState"`
			LinkKey   string `json:"linkKey"`
			Link      bool   `json:"link"`
		} `json:"link"`
	} `json:"istat"`
}
type Input struct {
	Did   string  `json:"did"`
	Type  string  `json:"type"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}
