package entities

import "log"


type Logger struct {
	Info  *log.Logger
	Error *log.Logger
}