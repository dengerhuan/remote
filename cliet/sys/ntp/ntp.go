package ntp

import (
	"client/protoc"
	"encoding/binary"
	"time"
)

var buf [64]byte

func RequestNetTimeStamp() []byte {

	buf[protoc.TypeIndex] = 0
	buf[protoc.DomainIndex] = 1
	buf[protoc.CodecIndex] = 1

	binary.BigEndian.PutUint32(buf[protoc.LenIndex:protoc.LenIndex+4], 32)
	binary.BigEndian.PutUint64(buf[20:28], uint64(time.Now().UnixNano()))

	return buf[:52]
}
