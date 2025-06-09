package main

import (
	"leaguies_backend/router"

	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
)

var adapter *chiadapter.ChiLambdaV2

func init() {
	r := router.NewRouter()
	adapter = chiadapter.NewV2(r)
}

func main() {
	lambda.Start(adapter.ProxyWithContextV2)
}
