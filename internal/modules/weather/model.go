package weather

type Data struct {
	Coord         Coord                `json:"coord"`
	Weather       []WeatherDescription `json:"weather"`
	Details       WeatherData          `json:"main"`
	Visibility    int64                `json:"visibility"`
	Wind          WindData             `json:"wind"`
	Cloud         CloudData            `json:"cloud"`
	TimestampUnix int64                `json:"dt"`
	CityID        int64                `json:"id"`
	Name          string               `json:"name"`
	Timezone      int64                `json:"timezone"`
}

type Coord struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

type WeatherDescription struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type WeatherData struct {
	Temperature float64 `json:"temp"`
	Pressure    int64   `json:"pressure"`
	Humidity    int64   `json:"humidity"`
	MinTemp     float64 `json:"temp_min"`
	MaxTemp     float64 `json:"temp_max"`
}

type WindData struct {
	Speed   float64 `json:"speed"` // Wind Speed
	Degrees int     `json:"deg"`   // Wind Direction
}

type CloudData struct {
	All int64 `json:"all"`
}

type SystemData struct {
	Country     string `json:"country"`
	SunriseUnix int64  `json:"sunrise"`
	SunsetUnix  int64  `json:"sunset"`
}
