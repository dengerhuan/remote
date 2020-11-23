package vp

import (
	. "client/eventbus"
	"client/tools"
	"encoding/binary"
	"fmt"
	"net"
	"time"
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
		//fmt.Println(msg)
	})

	//mock()
}

var mockBuf [40]byte

func mock() {

	ticket := time.NewTicker(time.Second / 50)

	go func() {
		var begin uint16 = 0
		var ss byte = 0
		for {

			select {

			case <-ticket.C:


				slice := mockBuf[:]

				//fmt.Println(begin)
				// mock steering

				binary.BigEndian.PutUint16(slice[20:22], begin)
				slice[25] = 1
				GlobalBus.Publish("cmd", slice)

				slice[22] = 255
				slice[23] = ss

				ss += 2
				// mock go
				//
				//binary.BigEndian.PutUint16(slice[20:22], 32768)
				//
				//if begin > 20000 {
				//	slice[22] = 200 // gas 200
				//	slice[23] = 255 //brake 0
				//} else {
				//	slice[22] = 255 // gas 0
				//	slice[23] = 200 // brake 200
				//}
				//
				//slice[24] = 2

				slice[25] = 1
				GlobalBus.Publish("cmd", slice)

				begin += 100
			}

		}

	}()

}

func VpInstall() {
	// Can卡固定IP
	canAddr := net.UDPAddr{
		IP:   net.IPv4(192, 168, 1, 10),
		Port: 8001,
	}

	can, err := net.DialUDP("udp", nil, &canAddr)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(can)

	go writeToCan(can)

	go read(can)
}

func writeToCan(conn *net.UDPConn) {
	//ticket := time.NewTicker(time.Millisecond * 20)

	ticket := tools.NewInterval(20)
	for {
		select {
		case <-ticket:

			conn.Write(goSteering(wheel))
			conn.Write(goThrottle(uint16(gas)))
			conn.Write(goBrake(uint16(brake)))
			conn.Write(goGear(gear))
		}
	}
}

func goSteering(steering uint16) []byte {

	if steering > 65535 {
		steering = 65535
	}

	if steering < 0 {
		steering = 0
	}
	slice := steeringBuf[:]

	slice[0] = 0x08

	slice[4] = 0x64

	_s := float64(steering) / steeringModulus

	s := int16(_s)

	var fs int16 = 4700 - s
	//_steering := 4700 - s

	//	6.97180851
	// +- 4700 65535
	// 1110110110100100

	//fmt.Println(fs)

	slice[5] = byte(fs)
	slice[6] = byte(fs >> 8)

	//binary.LittleEndian.PutUint16(slice[5:7], _steering) // + left - right

	//fmt.Println(slice)

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
		//fmt.Println(buf)
	}
}
