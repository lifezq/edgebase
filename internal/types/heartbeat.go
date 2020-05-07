package types

type HeartBeat struct {
	Client string `json:"client"`
	Time   int64  `json:"time"`
	Alive  bool   `json:"alive"`
}
