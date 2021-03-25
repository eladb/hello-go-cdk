# Hello, CDK for Go!

This is a "Hello, CDK!" sample written in Go.

It defines an app with a single stack which contains an AWS Lambda function and
an S3 Bucket. When an object is created in the bucket, the function is called
through bucket notifications.

## Prerequisites

* [Go 1.6](https://golang.org/dl/)
* CDK CLI: `npm install -g aws-cdk`

## Usage

* `cdk synth`: print the CloudFormation template for this sample
* `cdk deploy`: deploys to your account

## License

Apache 2.0.
