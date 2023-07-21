package loges

import (
	"github.com/olivere/elastic"
	log "github.com/sirupsen/logrus"
	"github.com/smartystreets/goconvey/convey"
	"gopkg.in/sohlich/elogrus.v3"
	"os"
	"strings"
	"testing"
	"time"
)

func TestLogES(t *testing.T) {
	convey.Convey("TestLogES", t, func() {
		// 1. get client
		client, err := elastic.NewClient(elastic.SetURL("http://192.168.9.176:9211", "http://192.168.9.176:9212",
			"http://192.168.9.176:9213", "http://192.168.9.176:9214"), elastic.SetSniff(false))
		if err != nil {
			log.Panic(err)
		}
		// 2. get hostname
		hostname, err := os.Hostname()
		if err != nil {
			log.Panic(err)
		}
		// 3. build index
		index := strings.Builder{}
		index.WriteString("testlog")
		index.WriteString("-")
		index.WriteString(time.Now().Format("20060102"))
		// 4. new hook
		// NewAsyncElasticHook creates new hook with asynchronous log
		// client - ElasticSearch client using gopkg.in/olivere/elastic.v5
		// host - host of system
		// level - log level
		// index - name of the index in ElasticSearch
		hook, err := elogrus.NewElasticHook(client, hostname, log.InfoLevel, index.String())
		//hook, err := elogrus.NewAsyncElasticHook(client, hostname, log.InfoLevel, index.String())
		if err != nil {
			log.Panic(err)
		}
		// 5. add hook
		log.AddHook(hook)
		// 6. use
		log.Info("test loges.")
	})
}
