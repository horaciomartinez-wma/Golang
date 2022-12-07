package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
   /*
   jsonData := []byte(`{
		"name": "morpheus",
		"job": "leader"
	}`)
  */

  jsonData, err := os.ReadFile("socios.json")
  if err != nil {log.Fatal(err)}

  return &events.APIGatewayProxyResponse{
    StatusCode:        200,
    //Body:              "Hello, World!",
    Body:              string(jsonData),
  }, nil
}

func main() {
  lambda.Start(handler)
}
