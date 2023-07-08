package model

type GeoInfo struct {
	PostalCode       string  `json:"postal_code"`
	HitCount         int     `json:"hit_count"`
	Address          string  `json:"address"`
	TokyoStaDistance float64 `json:"tokyo_sta_distance"`
}

type GeoApiResponse struct {
	Response struct {
		Location []Location `json:"location"`
	} `json:"response"`
}

type Location struct {
	City       string `json:"city"`
	CityKana   string `json:"city_kana"`
	Town       string `json:"town"`
	TownKana   string `json:"town_kana"`
	X          string `json:"x"`
	Y          string `json:"y"`
	Prefecture string `json:"prefecture"`
	Postal     string `json:"postal"`
}

type AccessLog struct {
	PostalCode   string `json:"postal_code"`
	RequestCount int    `json:"request_count"`
}
