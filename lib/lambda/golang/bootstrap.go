package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.SQSEvent) (int, error) {
	return 0, nil
}

func main() {
	lambda.Start(handler)
}
