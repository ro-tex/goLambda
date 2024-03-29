package main

import (
	"bytes"
	"context"
	"encoding/json"

	"goLambda/lib"

	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (lib.Response, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": "Go Serverless v1.0! Your function executed successfully!",
	})
	if err != nil {
		return lib.Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := lib.Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
