package mq

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/Shopify/sarama"
	"io/ioutil"
)

var cfg *MqConfig
var SyncProducer sarama.SyncProducer
//var AsyncProducer sarama.AsyncProducer

func init() {

	fmt.Print("init kafka producer, it may take a few seconds to init the connection\n")

	var err error

	cfg = &MqConfig{}
	LoadJsonConfig(cfg, "kafka.json")

	mqConfig := sarama.NewConfig()
	mqConfig.Net.SASL.Enable = true
	mqConfig.Net.SASL.User = cfg.Ak
	mqConfig.Net.SASL.Password = cfg.Password
	mqConfig.Net.SASL.Handshake = true

	mqConfig.Version = sarama.V0_10_2_1

	certBytes, err := ioutil.ReadFile(GetFullPath(cfg.CertFile))

	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		panic("kafka producer failed to parse root certificate")
	}

	mqConfig.Net.TLS.Config = &tls.Config{
		//Certificates:       []tls.Certificate{},
		RootCAs:            clientCertPool,
		InsecureSkipVerify: true,
	}

	mqConfig.Net.TLS.Enable = true
	mqConfig.Producer.Return.Successes = true

	if err = mqConfig.Validate(); err != nil {
		msg := fmt.Sprintf("Kafka producer config invalidate. config: %v. err: %v", *cfg, err)
		fmt.Println(msg)
		panic(msg)
	}

	SyncProducer, err = sarama.NewSyncProducer(cfg.Servers, mqConfig)
	if err != nil {
		msg := fmt.Sprintf("Kafak producer create fail. err: %v", err)
		fmt.Println(msg)
		panic(msg)
	}
	//mqConfig.Producer.Return.Successes = false
	//AsyncProducer, err = sarama.NewAsyncProducer(cfg.Servers, mqConfig)
	//if err != nil {
	//	msg := fmt.Sprintf("async Kafak producer create fail. err: %v", err)
	//	fmt.Println(msg)
	//	panic(msg)
	//}

	/**

	 */
}

func Produce(topic string, key string, content []byte) error {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.StringEncoder(key),
		Value:     sarama.ByteEncoder(content),
		//Timestamp: time.Now(),
	}

	_, _, err := SyncProducer.SendMessage(msg)
	if err != nil {
		msg := fmt.Sprintf("Send Error topic: %v. key: %v. content: %v", topic, key, content)
		fmt.Println(msg)
		return err
	}
	fmt.Printf("Send OK topic:%s key:%s value:%s\n", topic, key, content)

	return nil
}

//func produceAsync(topic string, key string, content string) error {
//
//	// on the same partition.
//	AsyncProducer.Input() <- &sarama.ProducerMessage{
//		Topic: "access_log",
//		Key:   sarama.StringEncoder(key),
//		Value: sarama.StringEncoder(content),
//	}
//	return nil
//}

//func main() {
//	//the key of the kafka messages
//	//do not set the same the key for all messages, it may cause partition im-balance
//	key := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
//	value := "this is a kafka message!"
//	produce(cfg.Topics[0], key, value)
//}
