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
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Shopify/sarama"
	log1 "github.com/infinimesh/infinimesh/pkg/log"
	"github.com/infinimesh/infinimesh/pkg/s3merger"
	"github.com/infinimesh/infinimesh/pkg/s3merger/s3mergerpb"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port string
)

func init() {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)
	viper.SetDefault("PORT", "8091")
	viper.AutomaticEnv()
	port = viper.GetString("PORT")

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

	log1.Info("starting s3connector")

	server := s3merger.NewServer()
	server.Log = log1.Named("s3merger server")

	// gRPC client initialization
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal("Failed to listen", port, zap.Error(err))
	}
	log1.Info("tcp listner started at port")

	s := grpc.NewServer()
	s3mergerpb.RegisterS3ReposServer(s, server)
	reflection.Register(s)
	log1.Info("s3merger server registered")

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve gRPC", zap.Error(err))
	}
}
