package main

import (
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v3"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

// go test -bench=. -benchmem -cpu=8 -count=1 -benchtime=10s

func BenchmarkGolangLog(b *testing.B) {
	b.ResetTimer()
	b.Run("BenchmarkGolangLog", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			log.Println("benchmark golang log.")
		}
	})
}

func BenchmarkLogrus(b *testing.B) {
	b.ResetTimer()
	b.Run("BenchmarkLogrus", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			logrus.Info("benchmark logrus.")
		}
	})
}

func BenchmarkLoges(b *testing.B) {
	b.ResetTimer()
	b.Run("BenchmarkLoges", func(b *testing.B) {
		// 1. get client
		client, err := elastic.NewClient(elastic.SetURL("esUrl1", "esUrl2", "esUrl3", "esUrl4"), elastic.SetSniff(false))
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
		//hook, err := elogrus.NewElasticHook(client, hostname, logrus.InfoLevel, index.String())
		hook, err := elogrus.NewAsyncElasticHook(client, hostname, logrus.InfoLevel, index.String())
		if err != nil {
			log.Panic(err)
		}
		// 5. add hook
		logrus.AddHook(hook)
		// 6. use
		for n := 0; n < b.N; n++ {
			logrus.Info("benchmark.")
		}
	})
}
