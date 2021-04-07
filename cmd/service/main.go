package main

import (
	"github.com/aldrinleal/sls-go-template/util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/shurcooL/go-goon"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"net/http"
)

var (
	ginLambda *ginadapter.GinLambda
	engine    *gin.Engine
)

func init() {
	log.Infof("Initializing")

	engine = gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	ginLambda = ginadapter.New(engine)
}

func main() {
	if util.IsRunningOnLambda() {
		lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
			log.Infof("req: %s", goon.Sdump(req))

			return ginLambda.ProxyWithContext(ctx, req)
		})
	} else {
		log.Fatalf("Oops", http.ListenAndServe(":"+util.EnvIf("PORT", "8000"), engine))
	}
}
