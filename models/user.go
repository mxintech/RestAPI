package models

// User ....
type User struct {
	Name string `json:"name"`
	Meta Meta
}

// Meta ...
type Meta struct {
	Nickname string `json:"nickname"`
	Age      int    `json:"age"`
}
