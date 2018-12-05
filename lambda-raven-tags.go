package lambdaraventags

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambdacontext"
	raven "github.com/getsentry/raven-go"
)

func init() {
	raven.SetTagsContext(map[string]string{
		"log_group_name":   lambdacontext.LogGroupName,
		"log_stream_name":  lambdacontext.LogStreamName,
		"function_name":    lambdacontext.FunctionName,
		"memory_limit":     fmt.Sprintf("%d", lambdacontext.MemoryLimitInMB),
		"function_version": lambdacontext.FunctionVersion,
	})
}
