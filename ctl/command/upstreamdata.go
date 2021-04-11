package command

// upstream target data struct
type upstream struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Create_at int64  `json:"created_at"`
}

type upstreams struct {
	Data []upstream `json:"data"`
}

type target struct {
	Id         string `json:"id"`
	Target     string `json:"target"`
	Weight     int    `json:"weight"`
	Upsteam    string
	Health     string  `json:"health"`
	Created_at float64 `json:"created_at"`
}

type targets struct {
	Data []target `json:"data"`
}
