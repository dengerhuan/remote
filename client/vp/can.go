package vp

import (
	. "client/eventbus"
	"client/tools"
	"encoding/binary"
	"fmt"
	"net"
)

var (
	steeringBuf [13]byte
	throttleBuf [13]byte
	brakeBuf    [13]byte
	gearBuf     [13]byte

	steeringModulus = 6.97180851

	manualRecoverFlag = false

	enable byte = 0
	wheel  uint16
	gas    byte
	brake  byte
	gear   byte
)

func init() {

	GlobalBus.Subscribe("cmd", func(message interface{}) {
		msg := message.([]byte)

		enable = msg[25]
		wheel = binary.BigEndian.Uint16(msg[20:22])
		gas = msg[22]
		brake = msg[23]
		gear = msg[24]

		//[0 0 0 0 0   0 1 2 0 0    0 1 0 0 0   6 0 0 0 0  6 73 43 231 2 0   0 0 0 0 0 0   22 66 102 226 109 93 18 64]
		fmt.Println(msg)
	})
}

func VpInstall() {
	// Can卡固定IP
	canAddr := net.UDPAddr{
		IP:   net.IPv4(192, 168, 56, 101),
		Port: 8002,
	}

	can, err := net.DialUDP("udp", nil, &canAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(can)

	go writeToCan(can)

	go read(can)
}

func writeToCan(conn *net.UDPConn) {
	//ticket := time.NewTicker(time.Millisecond * 20)

	ticket := tools.NewInterval(500)
	for {
		select {
		case <-ticket:

			conn.Write(goSteering(wheel))
			conn.Write(goThrottle(uint16(gas)))
			conn.Write(goBrake(uint16(brake)))
			conn.Write(goGear(gear))
			//conn.Write(drivemanager.DoSteering())
			////conn.Write(drivemanager.DoBrake())
			//conn.Write(drivemanager.DoThrottle())
			//conn.Write(drivemanager.DoGear())
		}
	}
}

func goSteering(steering uint16) []byte {
	slice := steeringBuf[:]

	slice[0] = 0x08

	slice[4] = 0x64

	_s := float64(steering) / steeringModulus

	s := int16(_s)
	_steering := s - 4700

	//	6.97180851
	// +- 4700 65535
	// 1110110110100100

	slice[5] = byte(_steering >> 8)
	slice[6] = byte(_steering)
	//binary.LittleEndian.PutUint16(slice[5:7], _steering) // + left - right

	slice[7] = byte(enable)
	return slice

}

func goBrake(brake uint16) []byte {
	brake = (255 - brake) * 257
	slice := brakeBuf[:]

	slice[0] = 0x08

	slice[4] = 0x60
	binary.LittleEndian.PutUint16(slice[5:7], brake)
	slice[7] = 1
	slice[8] = byte(enable)
	return slice

}

func goGear(gear byte) []byte {

	switch gear {

	case 0:
		gear = 3
	case 1, 4, 16:
		//gear=p/
		fmt.Println("1 = Park （can’t be used）")
	case 2, 8, 32:
		gear = 4
	case 64:
		gear = 2

	}
	// 0  1  2  4  8 16 32 64
	//0 = None--
	//1 = Park （can’t be used）
	//2 = Reverse //
	//3 = Neutral //
	//4 = Drive //
	//5 = Sport--

	slice := gearBuf[:]

	slice[0] = 0x08

	slice[4] = 0x66

	slice[5] = gear
	slice[6] = byte(enable)
	return slice
}

func goThrottle(throttle uint16) []byte {

	throttle = (255 - throttle) * 257
	// lg 255 0
	// 65535 100%
	slice := brakeBuf[:]

	slice[0] = 0x08

	slice[4] = 0x62
	binary.LittleEndian.PutUint16(slice[5:7], throttle)
	/*	//slice[7] = 1*/
	slice[8] = byte(enable)
	return slice
}

// read can state
func read(conn *net.UDPConn) {

	buf := make([]byte, 13)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Print(err)
			continue
		}
		// 读取can 状
		fmt.Println(buf)
	}
}
