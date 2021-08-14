package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	proto "github.com/golang/protobuf/proto"
	model "github.com/rnzsgh/serverless-processing-poc-ingest/protob/model"
)

func handler(ctx context.Context, kinesisEvent events.KinesisEvent) error {

	for _, record := range kinesisEvent.Records {

		kinesisRecord := record.Kinesis

		data := kinesisRecord.Data

		if len(data) == 0 {
			continue
		}

		group := &model.Group{}
		err := proto.Unmarshal(data, group)
		if err != nil {
			return fmt.Errorf("Failed to unmarshal event: %v", err)
		}

		fmt.Println("record handled")
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
