// Package cloudstorage helps detect where an Ambient app is running and provides
// the correct storage plugin.
package cloudstorage

import (
	"os"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/ambient/lib/envdetect"
	"github.com/ambientkit/plugin/storage/awsbucketstorage"
	"github.com/ambientkit/plugin/storage/azureblobstorage"
	"github.com/ambientkit/plugin/storage/gcpbucketstorage"
	"github.com/ambientkit/plugin/storage/localstorage"
)

// StorageBasedOnCloud returns storage engine based on the environment it's
// running in.
func StorageBasedOnCloud(sitePath string, sessionPath string) ambient.StoragePlugin {
	// Select the storage engine for site and session information.
	var storage ambient.StoragePlugin
	if envdetect.RunningLocalDev() {
		storage = localstorage.New(sitePath, sessionPath)
	} else if RunningInGoogle() {
		storage = gcpbucketstorage.New(sitePath, sessionPath)
	} else if RunningInAWS() {
		storage = awsbucketstorage.New(sitePath, sessionPath)
	} else if RunningInAzureFunction() {
		storage = azureblobstorage.New(sitePath, sessionPath)
	} else {
		// Defaulting to local storage.
		storage = localstorage.New(sitePath, sessionPath)
	}

	return storage
}

// RunningInAWS returns true if running in AWS services. When running in
// App Runner, it will be set: AWS_EXECUTION_ENV=AWS_ECS_FARGATE.
func RunningInAWS() bool {
	_, exists := os.LookupEnv("AWS_EXECUTION_ENV")
	return exists
}

// RunningInAWSLambda returns true if running in AWS Lambda.
func RunningInAWSLambda() bool {
	_, exists := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME")
	return exists
}

// RunningInGoogle returns true if running in Google. When running in
// Google Cloud Run, will be set: K_SERVICE=NAME.
func RunningInGoogle() bool {
	_, exists := os.LookupEnv("K_SERVICE")
	return exists
}

// RunningInAzureFunction returns true if running in Azure Functions.
func RunningInAzureFunction() bool {
	_, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	return exists
}
