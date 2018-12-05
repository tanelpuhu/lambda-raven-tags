# lambda-raven-tags

example:

	package main

	import (
		"os/user"

		raven "github.com/getsentry/raven-go"
		_ "github.com/tanelpuhu/lambda-raven-tags"

		"github.com/aws/aws-lambda-go/lambda"
	)

	func hello() {
		_, err := user.Current()
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
		}
	}

	func main() {
		lambda.Start(hello)
	}

and you'll get tags like this automatically:

	function_name: hello_function
	function_version: 28
	log_group_name: /aws/lambda/hello_function
	log_stream_name: 2018/12/05/[28]53c9ee0305d641c098b54ee7f5a1c390
	memory_limit: 128
