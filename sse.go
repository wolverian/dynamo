package dynamo

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// SSEType is used to specify the type of server side encryption
// to use on a table
type SSEType string

// Possible SSE types for tables
const (
	SSETypeAES256 SSEType = "AES256"
	SSETypeKMS    SSEType = "KMS"
)

type SSEDescription struct {
	InaccessibleEncryptionDateTime time.Time
	KMSMasterKeyArn                string
	SSEType                        types.SSEType
	Status                         types.SSEStatus
}
