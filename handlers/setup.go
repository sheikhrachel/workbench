package handlers

import (
	aws "github.com/sheikhrachel/workbench/api_common/aws_utils"
	"github.com/sheikhrachel/workbench/api_common/call"
	"github.com/sheikhrachel/workbench/workbench_service"
)

// Handler is a struct that holds the common dependencies for all endpoint handlers
type Handler struct {
	// appEnv is the environment the app is running in
	appEnv string
	// aws is the aws client
	aws *aws.AWSClient
	// workbenchService is the service interface responsible for business logic
	workbenchService *workbench_service.WorkbenchService
}

// New returns a new Handler pointer
func New(cc call.Call) *Handler {
	awsClient := aws.Init(cc)
	workbenchService := workbench_service.New(cc, awsClient)
	return &Handler{
		appEnv:           cc.Env,
		aws:              awsClient,
		workbenchService: workbenchService,
	}
}
