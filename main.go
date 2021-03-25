package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3notifications"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

func newMyStack(scope constructs.Construct, id string) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, &awscdk.StackProps{})

	handler := awslambda.NewFunction(stack, jsii.String("MyFunction"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_NODEJS_14_X(),
		Code:    awslambda.AssetCode_Asset(jsii.String("lambda")),
		Handler: jsii.String("index.handler"),
	})

	bucket := awss3.NewBucket(stack, jsii.String("MyBucket"), &awss3.BucketProps{})
	bucket.AddObjectCreatedNotification(awss3notifications.NewLambdaDestination(handler))

	return stack
}

func main() {
	app := awscdk.NewApp(&awscdk.AppProps{})
	newMyStack(app, "MyStack")
	app.Synth(&awscdk.StageSynthesisOptions{})
}
