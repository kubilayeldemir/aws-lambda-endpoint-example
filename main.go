package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, request events.LambdaFunctionURLRequest) (events.APIGatewayProxyResponse, error) {
	//If you are gonna use the functionURL instead of api gateway, use events.LambdaFunctionURLRequest as request model.
	//There are some differences. For example: FunctionURL do not pass HTTPMethod variable since it's a URL.
	if request.Body == "" {
		name, _ := request.QueryStringParameters["name"]
		message := fmt.Sprintf("Hello %s!!!", name)
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       message,
		}, nil
	}

	var myEvent MyEvent
	// Parse the JSON body from the event
	err := json.Unmarshal([]byte(request.Body), &myEvent)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	message := fmt.Sprintf("Hello %s!!!!", myEvent.Name)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       message,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
