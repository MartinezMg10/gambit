package main

import (
	"context"
	"os"
	"strings"

	"github.com/MartinezMG10/gambit/awsgo"
	"github.com/MartinezMG10/gambit/bd"
	"github.com/MartinezMG10/gambit/handlers"
	"github.com/aws/aws-lambda-go/events"

	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {

	awsgo.InicializoAws()

	if !ValidoParametros() {
		panic("Error en los parametros debe enviar 'SecretName','UrlPrefix' ")
	}

	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	headers := request.Headers

	bd.ReadSecret()

	status, message := handlers.Manejadores(path, method, body, headers, request)

	headersResp := map[string]string{
		"Content-type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headersResp,
	}

	return res, nil

}

func ValidoParametros() bool {
	_, traeParametro := os.LookupEnv("SecretName")
	if traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("UrlPrefix")
	if traeParametro {
		return traeParametro
	}

	return traeParametro
}
