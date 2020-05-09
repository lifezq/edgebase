package types

type HeartBeat struct {
	Client string `json:"client"`
	Time   int64  `json:"time"`
	Alive  bool   `json:"alive"`
}

type HeartBeatService struct {
	Service string `json:"service"`
	Time    int64  `json:"time"`
	Alive   bool   `json:"alive"`
}
