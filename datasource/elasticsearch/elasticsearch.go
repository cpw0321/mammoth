package elasticsearch

import (
	"fmt"
	"time"

	configs "github.com/cpw0321/mammoth/config"
	"github.com/cpw0321/mammoth/logger"

	"github.com/elastic/go-elasticsearch/v7"
)

// ESClient elasticsearch客户端
var ESClient *elasticsearch.Client

func InitES() {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses:     configs.Conf.Elasticsearch.Urls,
		RetryOnStatus: []int{502, 503, 504, 429},
		RetryBackoff:  func(i int) time.Duration { return time.Duration(i) * 100 * time.Millisecond },
		MaxRetries:    5,
	})
	if err != nil {
		logger.Log.Errorf("get elasticsearch client is err: %v", err)
		panic(fmt.Errorf("init elasticsearch client err: %v \n", err))
	}
	ESClient = client
}
