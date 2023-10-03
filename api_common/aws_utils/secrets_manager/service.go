package secrets_manager

import "github.com/sheikhrachel/workbench/api_common/call"

type SecretsManagerService interface {
	GetSecretValue(cc call.Call, key string) (string, error)
}
