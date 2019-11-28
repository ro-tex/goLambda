package main

import (
	"fmt"

	. "goLambda/lib"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(req Request) (Response, error) {

	godotenv.Load() // load .env variables
	if !HasValidAuth(&req) {

	}

	fmt.Println("Request body: ", req.Body)
	fmt.Println("Request method: ", req.HTTPMethod)
	fmt.Println("Request path+param: ", req.Path, req.PathParameters, req.QueryStringParameters)
	// Request path+param: /v0/docs/123 map[id:123] map[foo:bar]
	// https://gvcix41zob.execute-api.eu-west-1.amazonaws.com/dev/v0/docs/123?foo=bar

	return Response{
		StatusCode: 200,
		Body:       req.Body,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
