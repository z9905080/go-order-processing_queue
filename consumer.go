package go_order_processing_queue

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/nsqio/go-nsq"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// startProduceOrder to produce orders
func startProduceOrder() {

	// init nsq config
	config := nsq.NewConfig()

	// use default config to subscribe 'order' topic & subscribe 'order_channel' channel in this topic.
	q, _ := nsq.NewConsumer("order", "order_channel", config)

	// Set Log Level
	//q.SetLoggerLevel(nsq.LogLevelDebug)

	// add handler with receive new order message.
	q.AddConcurrentHandlers(nsq.HandlerFunc(handleFn), 10)

	nsqAddress := viper.GetString(FlagNSQDAddress) // retrieve value from viper
	// connect to NSQ cluster.
	err := q.ConnectToNSQLookupd(nsqAddress + ":4161")
	if err != nil {
		log.Panic("連線失敗。")
	}
	for {
		time.Sleep(5 * time.Second)
		lock.Lock()
		log.Println("count:", count)
		lock.Unlock()
	}
	gracefulShutdown()
	log.Println("End Of Receiver.")
}

// gracefulShutdown 等事情做完才會關閉
func gracefulShutdown() {

	ch := make(chan os.Signal, 10)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	sig := <-ch
	switch sig {
	case syscall.SIGINT, syscall.SIGTERM:
		signal.Stop(ch)
	}
	return
}

var count int
var lock sync.Mutex

func handleFn(message *nsq.Message) error {

	//log.Printf("receive an message：%v", message)

	orderStr := message.Body

	var orderData OrderData

	// if want to rollback this order ,can change order state
	if unMarshalErr := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(orderStr, &orderData); unMarshalErr != nil {
		log.Println("unMarshal Err ", unMarshalErr)
		return unMarshalErr
	}

	// 1 : processing
	if orderData.OrderState == 1 {
		lock.Lock()
		count++
		lock.Unlock()
		// produce this Order Data
		log.Printf("User: %v ,OrderID: %v , Add Money: %f", orderData.UserName, orderData.OrderID, orderData.Money)
		orderData.OrderState = 2
	}

	return nil
}
