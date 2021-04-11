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
