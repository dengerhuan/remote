package main

//
//func Mock() {
//	//   RFC3339     = "2006-01-02T15:04:05Z07:00"
//	//
//	//timeForrmat:=	"2006-01-02 15:04:05.07:00"
//	//	// kafka队列 json格式消息(topic : rdstat )
//	//	{
//	//		“status”:[
//	//	{“carid”: “12345678901234567”, “adminState”:true, “rdState”:true,
//	//	“speed”:17.3, “delay”：57, “distance”:23.3, “time”:41,
//	//	“longitude”:43.3133, “latitude”: 123.323436, “direction”:53.3,
//	//	“timeStamp”: “2000-07-05 14:35:14.99” },
//	//		{“carid”: “12345678901234568”, “adminState”:true, “rdState”:false, “speed”:17.3, “delay”：0, “distance”:23.3, “time”:41, “longitude”:43.3123, “latitude”: 123.323343, “direction”:53.3, “timeStamp”: “2000-07-05 14:35:14.97”  }
//	//
//	//],
//	//	“time”: “2000-07-05 14:35:15.00”
//	//	}，
//
//	carStat := gin.H{"carId": "12345678901234567", "adminState": true, "rdState": true, "speed": 17, "delay": 57, "distance": 23.3,
//		"time": 40, "longitude": 43.3133, "latitude": 123.323436,
//		"direction": 53.3, "timeStamp": time.Now().UnixNano() / 1e6}
//
//	stat := gin.H{"status": []interface{}{carStat}, "time": time.Now().UnixNano() / 1e6}
//
//	log.Println(stat)
//
//	message, _ := json.Marshal(stat)
//
//	/*
//		// on the same partition.
//		AsyncProducer.Input() <- &sarama.ProducerMessage{
//			Topic: "access_log",
//			Key:   sarama.StringEncoder(key),
//			Value: sarama.StringEncoder(content),
//		}
//		return nil
//	*/
//	ticket := time.NewTicker(time.Second * 5)
//
//	go func() {
//
//		for {
//			select {
//			//case <-con.Context().Done():
//			//	fmt.Println("channel done")
//			//	ticket.Stop()
//			//	runtime.Goexit()
//			case <-ticket.C:
//				err := mq.Produce("rdstat", strconv.FormatInt(time.Now().UnixNano()/1e6, 10), message)
//
//				log.Println(message)
//
//				if err != nil {
//
//					log.Println(err)
//				}
//				//con.Write(ntp.RequestNetTimeStamp())
//
//				//mq.AsyncProducer.Input() <-
//				//	&sarama.ProducerMessage{
//				//		Topic: "rdstat",
//				//		Value: sarama.ByteEncoder(message)}
//				//
//				//log.Println("msdfdsd")
//			}
//		}
//
//	}()
//}
//
///**
//{
//“orderId”: “2731489313”,
//“timeStamp”: “2000-07-05 12:34:34.32”,
//“driveLogOrigin”: {
//     “SteeringWheel”: 134，
//     “ThrottleAngle”: 34,
//     “Brake”:13,
//     “Gear”:1,
//     “Enable”:1
//   },
//“driveLogOptFirst”: {
//     “SteeringWheel”: 134，
//     “ThrottleAngle”: 34,
//     “Brake”:13,
//     “Gear”:1,
//     “Enable”:1
//   },
//“driveLogOptSecond”: {
//     “SteeringWheel”: 134，
//     “ThrottleAngle”: 34,
//     “Brake”:13,
//     “Gear”:1,
//     “Enable”:1
//   }
//}
//
//*/
//func MockState() {
//
//	s := gin.H{"orderId": "2731489313", "timeStamp": time.Now().UnixNano() / 1e6}
//	s["driveLogOrigin"] = gin.H{"SteeringWheel": 134, "ThrottleAngle": 34, "Brake": 13, "Gear": 1, "Enable": 1}
//	s["driveLogOptFirst"] = gin.H{"SteeringWheel": 134, "ThrottleAngle": 34, "Brake": 13, "Gear": 1, "Enable": 1}
//	s["driveLogOptSecond"] = gin.H{"SteeringWheel": 134, "ThrottleAngle": 34, "Brake": 13, "Gear": 1, "Enable": 1}
//	log.Println(s)
//
//	ticket := time.NewTicker(time.Second * 5)
//	message, _ := json.Marshal(s)
//	go func() {
//
//		for {
//			select {
//			//case <-con.Context().Done():
//			//	fmt.Println("channel done")
//			//	ticket.Stop()
//			//	runtime.Goexit()
//			case <-ticket.C:
//				err := mq.Produce("rdlog", strconv.FormatInt(time.Now().UnixNano()/1e6, 10), message)
//
//				//log.Println(message)
//
//				if err != nil {
//
//					log.Println(err)
//				}
//				//con.Write(ntp.RequestNetTimeStamp())
//
//				//mq.AsyncProducer.Input() <-
//				//	&sarama.ProducerMessage{
//				//		Topic: "rdstat",
//				//		Value: sarama.ByteEncoder(message)}
//				//
//				//log.Println("msdfdsd")
//			}
//		}
//
//	}()
//
//}
