package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Reqeust events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context, req Reqeust) (Response, error) {
	// var buf bytes.Buffer

	// body, err := json.Marshal(map[string]interface{}{
	// 	"message": "Okay",
	// })
	// if err != nil {
	// 	return Response{StatusCode: 404}, err
	// }
	// json.HTMLEscape(&buf, body)

	// var body string

	// if req.IsBase64Encoded {
	// 	body = "is"
	// } else {
	// 	body = "not is"
	// }

	resp := Response{
		StatusCode:      200,
		Body:            req.Body,
		IsBase64Encoded: true,
		Headers: map[string]string{
			"Content-Type":    "image/png",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
