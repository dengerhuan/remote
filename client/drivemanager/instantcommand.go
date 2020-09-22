package drivemanager

import (
	"bytes"
	"encoding/binary"
	"sync"
)

var (
	mu   sync.RWMutex
	FCmd Command
	SCmd Command
	TCmd Command
)

type Command struct {
	Wheel uint16
	Gas   byte
	Brake byte
	Gear  byte
	Flag  byte // 1 control 0 out control
}

func ReadCommand(msg []byte) {
	read := bytes.NewReader(msg)
	mu.RLock()
	defer mu.RUnlock()

	binary.Read(read, binary.LittleEndian, &FCmd.Wheel)
	binary.Read(read, binary.LittleEndian, &FCmd.Gas)
	binary.Read(read, binary.LittleEndian, &FCmd.Gear)
	binary.Read(read, binary.LittleEndian, &FCmd.Gear)
	binary.Read(read, binary.LittleEndian, &FCmd.Flag)
}
