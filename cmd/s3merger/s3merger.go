//--------------------------------------------------------------------------
// Copyright 2018 Infinite Devices GmbH
// www.infinimesh.io
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//--------------------------------------------------------------------------

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
	log1 "github.com/infinimesh/infinimesh/pkg/log"
	"github.com/infinimesh/infinimesh/pkg/s3merger"
	"github.com/infinimesh/infinimesh/pkg/s3merger/s3mergerpb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	sourceTopicReported = "shadow.reported-state.full"
	sourceTopicDesired  = "shadow.desired-state.full"
	consumerGroup       = "s3-persister"
)

var (
	broker         string
	s3Host         string
	s3MergerClient s3mergerpb.S3ReposClient
)

func init() {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)
	viper.SetDefault("KAFKA_HOST", "localhost:9092")
	viper.SetDefault("KAFKA_CONSUMER_GROUP", "s3-persister")
	viper.SetDefault("S3_Host", "localhost:8084")
	viper.AutomaticEnv()
	broker = viper.GetString("KAFKA_HOST")
	s3Host = viper.GetString("S3_Host")

	//consumerGroup = viper.GetString("KAFKA_CONSUMER_GROUP")
}

func main() {
	log1, err := log1.NewProdOrDev()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = log1.Sync()
	}()

	conn, err := grpc.Dial(s3Host, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("unable to connect to S3_Host")
	}
	defer conn.Close()

	s3MergerClient = s3mergerpb.NewS3ReposClient(conn)

	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Return.Errors = false
	config.Version = sarama.V2_0_0_0
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10

	client, err := sarama.NewClient([]string{broker}, config)
	if err != nil {
		panic(err)
	}

	group, err := sarama.NewConsumerGroupFromClient(consumerGroup, client)
	if err != nil {
		panic(err)
	}
	log1.Info("s3merger consumer group created")

	handler := &s3merger.Consumer{
		Log:                 log1.Named("S3 Connector Controller"),
		ConsumerGroup:       consumerGroup,
		SourceTopicDesired:  sourceTopicDesired,
		SourceTopicReported: sourceTopicReported,
		S3Client:            s3MergerClient,
	}

	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGINT)

	done := make(chan bool, 1)

	go func() {
	outer:
		for {

			err = group.Consume(context.Background(), []string{sourceTopicDesired, sourceTopicReported}, handler)
			log1.Info("messages reading")
			if err != nil {
				panic(err)
			}

			select {
			case <-done:
				break outer
			default:
			}

		}

	}()

	<-c
	done <- true
	err = group.Close()
	if err != nil {
		panic(err)
	}
}
