package context

// Error struct
type Error struct {
	Code int64  `json:"errcode"`
	Msg  string `json:"errmsg"`
}
