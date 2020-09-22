package main

import (
	"client/drivemanager"
	"client/rpc"
	"github.com/go-netty/go-netty"
	"log"
	"runtime"
	"time"
)

func MockStat(channel netty.Channel) {
	c := &rpc.Context{channel}

	//log.Println(stat)
	//
	//message, _ := json.Marshal(stat)
	ticket := time.NewTicker(time.Second * 1)
	i := 0

	l := len(mock)
	go func() {

		for {
			select {
			case <-ticket.C:

				i++

				poi := mock[i%l]
				if !drivemanager.RdState {
					break
				}

				wtime := drivemanager.GetTime()
				carStat := map[string]interface{}{"carId": drivemanager.CarId,
					"adminState": true,
					"rdState":    drivemanager.RdState,
					"speed":      17,
					"delay":      57,
					"distance":   17 * wtime / 3600,
					"time":       wtime,
					"longitude":  poi[0],
					"latitude":   poi[1],
					"direction":  53.3,
					"timeStamp":  time.Now().UnixNano() / 1e6}

				stat := map[string]interface{}{
					"status": []map[string]interface{}{
						carStat},
					"time": time.Now().UnixNano() / 1e6}

				c.RenderJson(c.MsgHead(1, 0), stat)

			case <-channel.Context().Done():
				log.Println("channel done")
				ticket.Stop()
				runtime.Goexit()
			}
		}

	}()

}

func MockCmd(channel netty.Channel) {
	c := &rpc.Context{channel}

	ticket := time.NewTicker(time.Second / 1)

	go func() {

		for {
			select {
			case <-ticket.C:

				if !drivemanager.RdState {
					break
				}

				s := H{"orderId": drivemanager.OrderId,
					"timeStamp": time.Now().UnixNano() / 1e6}
				s["driveLogOrigin"] = H{"SteeringWheel": drivemanager.FCmd.Wheel,
					"ThrottleAngle": drivemanager.FCmd.Gas,
					"Brake":         drivemanager.FCmd.Brake,
					"Gear":          drivemanager.FCmd.Gear,
					"Enable":        drivemanager.FCmd.Flag}
				s["driveLogOptFirst"] = H{"SteeringWheel": drivemanager.FCmd.Wheel,
					"ThrottleAngle": drivemanager.FCmd.Gas,
					"Brake":         drivemanager.FCmd.Brake,
					"Gear":          drivemanager.FCmd.Gear,
					"Enable":        drivemanager.FCmd.Flag}
				s["driveLogOptSecond"] = H{"SteeringWheel": drivemanager.FCmd.Wheel,
					"ThrottleAngle": drivemanager.FCmd.Gas,
					"Brake":         drivemanager.FCmd.Brake,
					"Gear":          drivemanager.FCmd.Gear,
					"Enable":        drivemanager.FCmd.Flag}
				c.RenderJson(c.MsgHead(1, 1), s)
			case <-channel.Context().Done():
				log.Println("channel done")
				ticket.Stop()
				runtime.Goexit()
			}
		}

	}()

}

var mock = [][]float64{
	{106.712209, 26.594248},
	{106.712111, 26.594211},
	{106.712543, 26.590641},
	{106.710987, 26.590687},
	{106.710274, 26.590516},
	{106.709966, 26.590518},
	{106.709793, 26.590513},
	{106.709741, 26.590505},
	{106.709706, 26.590497},
	{106.709688, 26.590497},
	{106.709709, 26.59054},
	{106.709708, 26.590539},
	{106.708642, 26.590372},
	{106.708617, 26.590372},
	{106.708617, 26.590372},
	{106.708413, 26.585014},
	{106.708413, 26.585014},
	{106.708728, 26.580542},
	{106.708466, 26.579375},
	{106.708461, 26.579335},
	{106.708457, 26.579297},
	{106.708096, 26.579157},
	{106.70793, 26.579176},
	{106.707838, 26.579177},
	{106.705962, 26.580007},
	{106.695202, 26.580181},
	{106.694806, 26.57991},
	{106.694779, 26.579892},
	{106.694765, 26.579881},
	{106.687205, 26.574189},
	{106.686461, 26.573919},
	{106.686414, 26.573905},
	{106.686377, 26.573898},
	{106.686356, 26.573901},
	{106.683839, 26.571346},
	{106.683741, 26.571264},
	{106.682505, 26.570289},
	{106.669384, 26.572032},
	{106.669384, 26.572032},
	{106.669384, 26.572032},
	{106.669384, 26.572032},
	{106.669384, 26.572032},
	{106.669384, 26.572032},
	{106.667872, 26.571683},
	{106.667316, 26.573113},
	{106.669024, 26.569604},
	{106.66972, 26.56855},
	{106.669776, 26.568484},
	{106.666482, 26.563008},
	{106.666403, 26.562561},
	{106.666388, 26.562508},
	{106.666753, 26.561872},
	{106.666739, 26.561823},
	{106.666732, 26.561767},
	{106.670071, 26.568048},
	{106.670547, 26.568108},
	{106.670682, 26.567973},
	{106.671034, 26.567537},
	{106.671078, 26.567501},
	{106.671078, 26.567501},
	{106.671088, 26.56748},
	{106.671092, 26.567468},
	{106.671086, 26.567467},
	{106.671077, 26.567469},
	{106.671072, 26.567468},
	{106.671066, 26.567466},
	{106.672974, 26.568573},
	{106.678942, 26.568292},
	{106.679231, 26.568256},
	{106.678206, 26.568512},
	{106.677908, 26.568528},
	{106.677561, 26.568538},
	{106.67737, 26.568584},
	{106.677253, 26.568606},
	{106.672255, 26.568125},
	{106.672184, 26.568059},
	{106.671836, 26.567725},
	{106.67179, 26.567687},
	{106.671716, 26.567634},
	{106.671707, 26.56763},
	{106.6717, 26.567625},
	{106.671694, 26.56763},
	{106.671682, 26.567629},
	{106.670028, 26.565333},
	{106.669929, 26.565095},
	{106.668612, 26.562095},
	{106.668262, 26.561047},
	{106.668232, 26.560947},
	{106.668163, 26.560759},
	{106.667714, 26.560438},
	{106.667598, 26.560439},
	{106.667018, 26.560455},
	{106.666884, 26.560477},
	{106.666764, 26.560507},
	{106.669673, 26.56798},
	{106.667444, 26.56589},
	{106.6674, 26.565805},
	{106.667186, 26.565601},
	{106.666343, 26.561497},
}
