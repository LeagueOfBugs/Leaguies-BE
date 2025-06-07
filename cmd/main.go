package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"leaguies_backend/router"
)

var adapter *chiadapter.ChiLambda

func init() {
	r := router.NewRouter()
	adapter = chiadapter.New(r)
}

func main() {
	lambda.Start(adapter.ProxyWithContext)
}
