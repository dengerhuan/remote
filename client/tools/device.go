package tools

import (
	"client/crypto"
)
//
//func GetMAC() {
//
//	fmt.Println("sas")
//
//	//in := getFirstNetworkCard()
//	//fmt.Println(in.HardwareAddr.String())
//
//	//did := getHardDiskInfo()
//	//if did != 0 {
//	//	fmt.Println(did)
//	//}
//
//	did := GenDeviceIdByHardDisk()
//
//	fmt.Println(did)
//}

func GenDeviceIdByHardDisk() string {
	//deviceId := getHardDiskInfo()
	//if deviceId == 0 {
	//	panic("system error could not get hark disk info")
	//}

	str := crypto.GetMd5("nokia")
	return str
}
//
//func getFirstNetworkCard() net.Interface {
//
//	interfaces, err := net.Interfaces()
//	if err != nil {
//		panic("Poor soul,here is what you got: " + err.Error())
//	}
//	if len(interfaces) == 0 {
//		panic("Poor soul,here is what you got: no network card found ")
//	}
//	return interfaces[0]
//
//}

func IntToString(orig []int8) string {
	ret := make([]byte, len(orig))
	size := -1
	for i, o := range orig {
		if o == 0 {
			size = i
			break
		}
		ret[i] = byte(o)
	}
	if size == -1 {
		size = len(orig)
	}

	return string(ret[0:size])
}
