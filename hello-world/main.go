package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cfg := zap.Config{
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zap.DebugLevel),
	}
	logger, _ := cfg.Build()
	defer logger.Sync()
	logger.Debug("Received request")
	resp, err := http.Get(DefaultHTTPGetAddress)
	if err != nil {
		logger.Error("Error while getting ip address", zap.Error(err))
		return events.APIGatewayProxyResponse{}, err
	}

	if resp.StatusCode != 200 {
		logger.Error("Unexpected status code while getting ip address", zap.Int("status", resp.StatusCode))
		return events.APIGatewayProxyResponse{}, ErrNon200Response
	}

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error while reading response", zap.Error(err))
		return events.APIGatewayProxyResponse{}, err
	}

	if len(ip) == 0 {
		logger.Warn("No IP address found")
		return events.APIGatewayProxyResponse{}, ErrNoIP
	}

	logger.Info("Sending response")
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", string(ip)),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
