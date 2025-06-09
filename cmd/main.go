package main

import (
	"leaguies_backend/internal/config"
	"leaguies_backend/internal/db"
	"leaguies_backend/router"

	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
)

var adapter *chiadapter.ChiLambdaV2

func init() {
	config.LoadEnv()
	db.Connect()
	r := router.NewRouter()
	adapter = chiadapter.NewV2(r)
}

func main() {
	lambda.Start(adapter.ProxyWithContextV2)
}
