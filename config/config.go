package config

type Config struct {
	IsMaintenance bool   `json:"isMaintenance"`
	UpdateLink    string `json:"updateLink"`
	Port          int    `json:"port"`
	Major         int    `json:"major"`
	Minor         int    `json:"minor"`
	Build         int    `json:"build"`
	ServerName    string `json:"serverName"`
}

var Default = Config{
	IsMaintenance: false,
	UpdateLink:    "https://t.me/vbcbrawl",
	Port:          1488,
	Major:         19,
	Minor:         0,
	Build:         111,
	ServerName:    "Sigma Brawl",
}