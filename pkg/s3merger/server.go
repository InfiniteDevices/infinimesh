package s3merger

import (
	"bytes"
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/infinimesh/infinimesh/pkg/s3merger/s3mergerpb"
	"go.uber.org/zap"
)

type Server struct {
	Log *zap.Logger
}

var (
	srv *s3.S3
)

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SetDeviceState(context context.Context, request *s3mergerpb.SaveDeviceStateRequest) (response *s3mergerpb.SaveDeviceStateResponse, err error) {
	log := s.Log.Named("Set Device Status Controller")

	//to open the s3 session
	srv = s3.New(session.New())

	log.Info("Function Invoked", zap.String("Device", request.Bucket))
	_, err = srv.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(request.Bucket),
	})
	if err != nil {
		log.Fatal("failed to create bucket")
	}
	_, err = srv.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(request.Bucket),
		Key:    aws.String(request.Bucket + "" + time.Now().String()),
		Body:   bytes.NewReader(request.DeviceState),
	})
	if err != nil {
		log.Fatal("failed to write a record to bucket")
	} else {
		log.Info("successfully wrote record to s3")
	}
	response = &s3mergerpb.SaveDeviceStateResponse{
		Status: true,
	}
	return response, nil
}
