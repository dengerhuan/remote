package drivemanager

import "time"

var OrderId string

var RdState bool

var StartTime int64

var CarId string

func GetTime() int64 {
	return time.Now().UnixNano()/1e9 - StartTime
}

var SysTimeDiff int64


var HH int64