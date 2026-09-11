package config

type Config struct {
	IsMaintenance bool   `json:"isMaintenance"`
	UpdateLink    string `json:"updateLink"`
	Port          int    `json:"port"`
	Major         int    `json:"major"`
	Minor         int    `json:"minor"`
	Build         int    `json:"build"`
	ServerName    string `json:"serverName"`

	Database DatabaseConfig `json:"database"` 
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

var Default = Config{
	IsMaintenance: false,
	UpdateLink:    "https://t.me/vbcbrawl",
	Port:          9339,
	Major:         19,
	Minor:         0,
	Build:         111,
	ServerName:    "Sigma Brawl",
	Database: DatabaseConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "v19go",
		Password: "v19go",
		DBName:   "v19go",
	},
}