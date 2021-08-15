package main

import (
	"time"
)

type User struct {
	Tagname  string
	Messages []*Message
	RankLvl  int
	RankPts  int
}

func NewUser(tagname string) *User {
	return &User{tagname, nil, 0, 0}
}

type Message struct {
	ID        int
	Data      time.Time
	UpVotes   int
	DownVotes int
}

func NewMessage(id int) *Message {
	return &Message{id, time.Now().UTC(), 0, 0}
}
