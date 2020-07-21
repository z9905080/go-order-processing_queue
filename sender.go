package go_order_processing_queue

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/nsqio/go-nsq"
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

func startSendOrder() {
	// init nsq config
	config := nsq.NewConfig()
	config.DialTimeout = 15 * time.Second

	nsqAddress := viper.GetString(FlagNSQDAddress) // retrieve value from viper

	// use the config to connect the nsqd.
	w, _ := nsq.NewProducer(nsqAddress+":4150", config)

	var wait sync.WaitGroup
	for i := 0; i < 10000; i++ {

		wait.Add(1)
		go func(ii int) {
			data := OrderData{
				UserName:   "A",
				OrderID:    "ac01",
				OrderType:  "A",
				OrderState: 1,
				Memo:       "Test",
				Money:      0.5 + float64(ii),
			}

			jsonBytes, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(data)
			//ch := make(chan *nsq.ProducerTransaction)
			err := w.Publish("order", jsonBytes)
			if err != nil {
				log.Panic("連線失敗。", err)
			}
			wait.Done()
		}(i)
	}
	wait.Wait()
	w.Stop()

}
