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
package s3merger

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/infinimesh/infinimesh/pkg/s3merger/s3mergerpb"
	"github.com/infinimesh/infinimesh/pkg/shadow"
	"go.uber.org/zap"
)

type Consumer struct {
	Log                 *zap.Logger
	S3Client            s3mergerpb.S3ReposClient
	SourceTopicReported string
	SourceTopicDesired  string
	ConsumerGroup       string
	S3Session           *s3.S3
}

//setting up the consumer for kafka and open the session for s3
func (h *Consumer) Setup(s sarama.ConsumerGroupSession) error {
	fmt.Println("Rebalance, assigned partitions:", s.Claims())
	return nil
}

func (h *Consumer) Cleanup(s sarama.ConsumerGroupSession) error {
	return nil
}

//consuming data from kafka topics and writing to s3
func (h *Consumer) ConsumeClaim(s sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {

		fmt.Println("got msg", string(message.Value))

		var stateFromKafka shadow.DeviceStateMessage
		if err := json.Unmarshal(message.Value, &stateFromKafka); err != nil {
			fmt.Println("Failed to deserialize message with offset ", message.Offset)
			continue
		}

		j, _ := json.Marshal(stateFromKafka.State)
		/*
			record := []string{
				string(message.Key),
				string(stateFromKafka.Version),
				string(j),
			}*/
		stat, err := h.S3Client.SetDeviceState(context.Background(), &s3mergerpb.SaveDeviceStateRequest{
			Bucket:      string(message.Key),
			Key:         string(message.Key),
			DeviceState: j})

		// FIXME ignore errors for now
		if err != nil {
			fmt.Println("Failed to persist message with offset due to error", message.Offset, err)
		}
		if stat.Status == false {
			fmt.Println("Failed to persist message ", message.Offset, stat)
		}

		s.MarkMessage(message, "")
	}
	return nil
}
