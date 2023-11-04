package service

type Msgservice interface {
	Sendmsg(msg string) bool
}
