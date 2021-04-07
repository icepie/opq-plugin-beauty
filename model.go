package main

type Friendpic struct {
	FileMd5  string `json:"FileMd5"`
	FileSize int64  `json:"FileSize"`
	Path     string `json:"Path"`
	Url      string `json:"Url"`
}

type FriendPicContent struct {
	Content   interface{} `json:"Content"`
	Friendpic []Friendpic `json:"FriendPic"`
	Tips      string      `json:"Tips"`
}

type GroupPic struct {
	FileId       int64  `json:"FileId"`
	FileMd5      string `json:"FileMd5"`
	FileSize     int64  `json:"FileSize"`
	ForwordBuf   string `json:"ForwordBuf"`
	ForwordField int64  `json:"ForwordField"`
	Url          string `json:"Url"`
}

type GroupPicContent struct {
	Content  interface{} `json:"Content"`
	GroupPic []GroupPic  `json:"GroupPic"`
	Tips     string      `json:"Tips"`
}
