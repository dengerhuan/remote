package tools

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"gw/crypto"
)

//
//func GetMAC() {
//
//	fmt.Println("sas")
//
//	in := getFirstNetworkCard()
//	fmt.Println(in.HardwareAddr.String())
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

//func CpuInfo() {
//	cpuInfos, err := cpu.Info()
//	if err != nil {
//		fmt.Printf("get cpu info failed, err:%v", err)
//	}
//
//	fmt.Println(len(cpuInfos))
//	for _, ci := range cpuInfos {
//		fmt.Println(ci)
//
//	}
//	// CPU使用率
//	//for {
//	//	percent, _ := cpu.Percent(time.Second, false)
//	//	fmt.Printf("cpu percent:%v\n", percent)
//	//}
//}
func GenDeviceIdByHardDisk() string {
	//deviceId := getHardDiskInfo()
	//if deviceId == 0 {
	//	panic("system error could not get hark disk info")
	//}

	str := crypto.GetMd5("nokiagw")

	//log.Println(str)
	return str
}

//
//// get hard disk info  get and return  hark disk type == 19(devfs )  id
//func getHardDiskInfo() int32 {
//	count, err := unix.Getfsstat(nil, unix.MNT_WAIT)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fs := make([]unix.Statfs_t, count)
//	if _, err = unix.Getfsstat(fs, unix.MNT_WAIT); err != nil {
//		fmt.Println(err)
//	}
//
//	for _, stat := range fs {
//		if stat.Type == 19 {
//			return stat.Fsid.Val[0]
//		}
//	}
//	return 0
//}
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
//
//func IntToString(orig []int8) string {
//	ret := make([]byte, len(orig))
//	size := -1
//	for i, o := range orig {
//		if o == 0 {
//			size = i
//			break
//		}
//		ret[i] = byte(o)
//	}
//	if size == -1 {
//		size = len(orig)
//	}
//
//	return string(ret[0:size])
//}

func Goid() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
